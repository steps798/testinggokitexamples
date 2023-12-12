package booking

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-kit/examples/shipping/cargo"
	"github.com/go-kit/examples/shipping/location"
	"github.com/go-kit/examples/shipping/routing"
)

func TestNewService(t *testing.T) {
	type args struct {
		cargos    cargo.Repository
		locations location.Repository
		events    cargo.HandlingEventRepository
		rs        routing.Service
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.cargos, tt.args.locations, tt.args.events, tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assemble(t *testing.T) {
	type args struct {
		c      *cargo.Cargo
		events cargo.HandlingEventRepository
	}
	tests := []struct {
		name string
		args args
		want Cargo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assemble(tt.args.c, tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assemble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_AssignCargoToRoute(t *testing.T) {
	type fields struct {
		cargos         cargo.Repository
		locations      location.Repository
		handlingEvents cargo.HandlingEventRepository
		routingService routing.Service
	}
	type args struct {
		id        cargo.TrackingID
		itinerary cargo.Itinerary
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cargos:         tt.fields.cargos,
				locations:      tt.fields.locations,
				handlingEvents: tt.fields.handlingEvents,
				routingService: tt.fields.routingService,
			}
			if err := s.AssignCargoToRoute(tt.args.id, tt.args.itinerary); (err != nil) != tt.wantErr {
				t.Errorf("AssignCargoToRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_BookNewCargo(t *testing.T) {
	type fields struct {
		cargos         cargo.Repository
		locations      location.Repository
		handlingEvents cargo.HandlingEventRepository
		routingService routing.Service
	}
	type args struct {
		origin      location.UNLocode
		destination location.UNLocode
		deadline    time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    cargo.TrackingID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cargos:         tt.fields.cargos,
				locations:      tt.fields.locations,
				handlingEvents: tt.fields.handlingEvents,
				routingService: tt.fields.routingService,
			}
			got, err := s.BookNewCargo(tt.args.origin, tt.args.destination, tt.args.deadline)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookNewCargo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BookNewCargo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Cargos(t *testing.T) {
	type fields struct {
		cargos         cargo.Repository
		locations      location.Repository
		handlingEvents cargo.HandlingEventRepository
		routingService routing.Service
	}
	tests := []struct {
		name   string
		fields fields
		want   []Cargo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cargos:         tt.fields.cargos,
				locations:      tt.fields.locations,
				handlingEvents: tt.fields.handlingEvents,
				routingService: tt.fields.routingService,
			}
			if got := s.Cargos(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cargos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_ChangeDestination(t *testing.T) {
	type fields struct {
		cargos         cargo.Repository
		locations      location.Repository
		handlingEvents cargo.HandlingEventRepository
		routingService routing.Service
	}
	type args struct {
		id          cargo.TrackingID
		destination location.UNLocode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cargos:         tt.fields.cargos,
				locations:      tt.fields.locations,
				handlingEvents: tt.fields.handlingEvents,
				routingService: tt.fields.routingService,
			}
			if err := s.ChangeDestination(tt.args.id, tt.args.destination); (err != nil) != tt.wantErr {
				t.Errorf("ChangeDestination() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_LoadCargo(t *testing.T) {
	type fields struct {
		cargos         cargo.Repository
		locations      location.Repository
		handlingEvents cargo.HandlingEventRepository
		routingService routing.Service
	}
	type args struct {
		id cargo.TrackingID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Cargo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cargos:         tt.fields.cargos,
				locations:      tt.fields.locations,
				handlingEvents: tt.fields.handlingEvents,
				routingService: tt.fields.routingService,
			}
			got, err := s.LoadCargo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCargo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadCargo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Locations(t *testing.T) {
	type fields struct {
		cargos         cargo.Repository
		locations      location.Repository
		handlingEvents cargo.HandlingEventRepository
		routingService routing.Service
	}
	tests := []struct {
		name   string
		fields fields
		want   []Location
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cargos:         tt.fields.cargos,
				locations:      tt.fields.locations,
				handlingEvents: tt.fields.handlingEvents,
				routingService: tt.fields.routingService,
			}
			if got := s.Locations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Locations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_RequestPossibleRoutesForCargo(t *testing.T) {
	type fields struct {
		cargos         cargo.Repository
		locations      location.Repository
		handlingEvents cargo.HandlingEventRepository
		routingService routing.Service
	}
	type args struct {
		id cargo.TrackingID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []cargo.Itinerary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cargos:         tt.fields.cargos,
				locations:      tt.fields.locations,
				handlingEvents: tt.fields.handlingEvents,
				routingService: tt.fields.routingService,
			}
			if got := s.RequestPossibleRoutesForCargo(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestPossibleRoutesForCargo() = %v, want %v", got, tt.want)
			}
		})
	}
}
