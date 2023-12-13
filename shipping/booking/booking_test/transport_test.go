package booking

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-kit/examples/shipping/booking"
	mock_booking "github.com/go-kit/examples/shipping/booking/mock"
	"github.com/go-kit/examples/shipping/cargo"
	"github.com/go-kit/examples/shipping/location"
	kitlog "github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_bookCargoHandler(t *testing.T) {

	timeNow := time.Now()

	type bookNewCargo struct {
		output interface{}
		err    error
	}
	type mockScenario struct {
		bookNewCargo *bookNewCargo
	}
	type args struct {
		requestBody string
	}

	tests := []struct {
		name               string
		args               args
		mockScenario       mockScenario
		expectedStatusCode int
	}{
		{
			name: "expected request body returns OK",
			args: args{
				requestBody: fmt.Sprintf(`{"origin":"origin","destination":"destination","arrival_deadline":"%s"}`, timeNow.Format(time.RFC3339)),
			},
			mockScenario: mockScenario{
				bookNewCargo: &bookNewCargo{
					output: cargo.NextTrackingID(),
					err:    nil,
				},
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "invalid request body returns bad request",
			args: args{
				requestBody: fmt.Sprintf(`{"origin":"","destination":"","arrival_deadline":"%s"}`, timeNow.Format(time.RFC3339)),
			},
			mockScenario: mockScenario{
				bookNewCargo: &bookNewCargo{
					output: cargo.NextTrackingID(),
					err:    booking.ErrInvalidArgument,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "unexpected error returns internal server error",
			args: args{
				requestBody: fmt.Sprintf(`{"origin":"origin","destination":"destination","arrival_deadline":"%s"}`, timeNow.Format(time.RFC3339)),
			},
			mockScenario: mockScenario{
				bookNewCargo: &bookNewCargo{
					output: cargo.NextTrackingID(),
					err:    errors.New("unexpected error"),
				},
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockCtr := gomock.NewController(t)
			mockService := mock_booking.NewMockService(mockCtr)
			handler := booking.MakeHandler(mockService, kitlog.NewNopLogger())
			responseRecorder := httptest.NewRecorder()

			if test.mockScenario.bookNewCargo != nil {
				mockService.EXPECT().
					BookNewCargo(
						gomock.AssignableToTypeOf(location.UNLocode("")),
						gomock.AssignableToTypeOf(location.UNLocode("")),
						gomock.AssignableToTypeOf(timeNow),
					).
					Return(test.mockScenario.bookNewCargo.output, test.mockScenario.bookNewCargo.err)
			}

			request := httptest.NewRequest(
				http.MethodPost,
				"http://localhost/booking/v1/cargos",
				strings.NewReader(test.args.requestBody),
			)
			handler.ServeHTTP(responseRecorder, request)

			response := responseRecorder.Result()
			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			assert.Equal(t, "application/json; charset=utf-8", response.Header.Get("Content-Type"))

		})
	}
}
