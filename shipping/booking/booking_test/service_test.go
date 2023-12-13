package booking_test

import (
	"errors"
	"testing"

	"github.com/go-kit/examples/shipping/booking"
	"github.com/go-kit/examples/shipping/cargo"
	mock_cargo "github.com/go-kit/examples/shipping/cargo/mock"
	mock_location "github.com/go-kit/examples/shipping/location/mock"
	mock_routing "github.com/go-kit/examples/shipping/routing/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_service_AssignCargoToRoute(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping possibly time-consuming test in short mode.")
	}

	type cargosFind struct {
		argId  cargo.TrackingID
		output *cargo.Cargo
		err    error
	}
	type cargosStore struct {
		argCargo *cargo.Cargo
		err      error
	}
	type mockScenario struct {
		cargosFind  *cargosFind
		cargosStore *cargosStore
	}
	type args struct {
		id        cargo.TrackingID
		itinerary cargo.Itinerary
	}
	tests := []struct {
		name         string
		args         args
		mockScenario mockScenario
		expectedErr  error
	}{
		{
			name: "empty id returns expected error",
			args: args{
				id: "",
				itinerary: cargo.Itinerary{
					Legs: []cargo.Leg{
						{},
					},
				},
			},
			mockScenario: mockScenario{},
			expectedErr:  booking.ErrInvalidArgument,
		},
		{
			name: "empty itinerary legs returns expected error",
			args: args{
				id:        "id",
				itinerary: cargo.Itinerary{},
			},
			mockScenario: mockScenario{},
			expectedErr:  booking.ErrInvalidArgument,
		},
		{
			name: "failed to find cargo returns expected error",
			args: args{
				id: "id",
				itinerary: cargo.Itinerary{
					Legs: []cargo.Leg{
						{},
					},
				},
			},
			mockScenario: mockScenario{
				cargosFind: &cargosFind{
					argId:  "id",
					output: nil,
					err:    errors.New("unexpected error"),
				},
			},
			expectedErr: errors.New("unexpected error"),
		},
		{
			name: "failed to store cargo returns expected error",
			args: args{
				id: "id",
				itinerary: cargo.Itinerary{
					Legs: []cargo.Leg{
						{},
					},
				},
			},
			mockScenario: mockScenario{
				cargosFind: &cargosFind{
					argId:  "id",
					output: &cargo.Cargo{},
					err:    nil,
				},
				cargosStore: &cargosStore{
					argCargo: &cargo.Cargo{
						TrackingID:         "",
						Origin:             "",
						RouteSpecification: cargo.RouteSpecification{},
						Itinerary: cargo.Itinerary{
							Legs: []cargo.Leg{
								{},
							},
						},
						Delivery: cargo.Delivery{
							Itinerary: cargo.Itinerary{
								Legs: []cargo.Leg{
									{},
								},
							},
							RoutingStatus:   cargo.Routed,
							TransportStatus: cargo.NotReceived,
							NextExpectedActivity: cargo.HandlingActivity{
								Type: cargo.Receive,
							},
						},
					},
					err: errors.New("unexpected error"),
				},
			},
			expectedErr: errors.New("unexpected error"),
		},
		{
			name: "successfully assign cargo to route",
			args: args{
				id: "id",
				itinerary: cargo.Itinerary{
					Legs: []cargo.Leg{
						{},
					},
				},
			},
			mockScenario: mockScenario{
				cargosFind: &cargosFind{
					argId:  "id",
					output: &cargo.Cargo{},
					err:    nil,
				},
				cargosStore: &cargosStore{
					argCargo: &cargo.Cargo{
						TrackingID:         "",
						Origin:             "",
						RouteSpecification: cargo.RouteSpecification{},
						Itinerary: cargo.Itinerary{
							Legs: []cargo.Leg{
								{},
							},
						},
						Delivery: cargo.Delivery{
							Itinerary: cargo.Itinerary{
								Legs: []cargo.Leg{
									{},
								},
							},
							RoutingStatus:   cargo.Routed,
							TransportStatus: cargo.NotReceived,
							NextExpectedActivity: cargo.HandlingActivity{
								Type: cargo.Receive,
							},
						},
					},
					err: nil,
				},
			},
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtr := gomock.NewController(t)
			mockCargos := mock_cargo.NewMockRepository(mockCtr)
			mockLocations := mock_location.NewMockRepository(mockCtr)
			mockHandlingEvents := mock_cargo.NewMockHandlingEventRepository(mockCtr)
			mockRoutingService := mock_routing.NewMockService(mockCtr)

			s := booking.NewService(
				mockCargos,
				mockLocations,
				mockHandlingEvents,
				mockRoutingService,
			)

			if test.mockScenario.cargosFind != nil {
				mockCargos.EXPECT().
					Find(test.mockScenario.cargosFind.argId).
					Return(test.mockScenario.cargosFind.output, test.mockScenario.cargosFind.err)
			}
			if test.mockScenario.cargosStore != nil {
				mockCargos.EXPECT().
					Store(test.mockScenario.cargosStore.argCargo).
					Return(test.mockScenario.cargosStore.err)
			}

			err := s.AssignCargoToRoute(test.args.id, test.args.itinerary)
			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

// TODO: Test_service_BookNewCargo
// TODO: Test_service_Cargos
// TODO: Test_service_ChangeDestination
// TODO: Test_service_LoadCargo
// TODO: Test_service_Locations
// TODO: Test_service_RequestPossibleRoutesForCargo
