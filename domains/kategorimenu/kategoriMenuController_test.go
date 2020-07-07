package kategorimenu

import (
	"database/sql"
	"net/http"
	"reflect"
	"testing"
)

func TestNewController(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *Controller
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_HandleGETAllKategoriMenus(t *testing.T) {
	tests := []struct {
		name string
		s    *Controller
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HandleGETAllKategoriMenus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.HandleGETAllKategoriMenus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_HandleGETKategoriMenu(t *testing.T) {
	tests := []struct {
		name string
		s    *Controller
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HandleGETKategoriMenu(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.HandleGETKategoriMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_HandlePOSTKategoriMenus(t *testing.T) {
	tests := []struct {
		name string
		s    *Controller
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HandlePOSTKategoriMenus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.HandlePOSTKategoriMenus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_HandleUPDATEKategoriMenus(t *testing.T) {
	tests := []struct {
		name string
		s    *Controller
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HandleUPDATEKategoriMenus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.HandleUPDATEKategoriMenus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_HandleDELETEKategoriMenus(t *testing.T) {
	tests := []struct {
		name string
		s    *Controller
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HandleDELETEKategoriMenus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.HandleDELETEKategoriMenus() = %v, want %v", got, tt.want)
			}
		})
	}
}
