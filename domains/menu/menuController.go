package menu

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"warung_makan_gerin/utils/message"
	"warung_makan_gerin/utils/tools"
)

type Controller struct {
	MenuService MenuServiceInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{MenuService: NewMenuService(db)}
}

func (s *Controller) HandleGETAllMenus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var page string = tools.GetPathVar("page", r)
		var limit string = tools.GetPathVar("limit", r)
		var status string = tools.GetPathVar("status", r)
		var orderBy string = tools.GetPathVar("orderBy", r)
		var sort string = tools.GetPathVar("sort", r)
		var keyword string = tools.GetPathVar("keyword", r)

		Menus, err := s.MenuService.GetMenus(keyword, page, limit, status, orderBy, sort)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Search All Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Search All Success", http.StatusOK, Menus))
	}
}

func (s *Controller) HandleGETMenu() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		Menu, err := s.MenuService.GetMenuByID(tools.GetPathVar("id", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Search by ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Search by ID Success", http.StatusOK, Menu))
	}
}

func (s *Controller) HandlePOSTMenus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data Menu
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.MenuService.HandlePOSTMenu(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Posting Success", http.StatusOK, result))
	}
}

func (s *Controller) HandleUPDATEMenus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data Menu
		tools.Parser(r, &data)

		result, err := s.MenuService.HandleUPDATEMenu(tools.GetPathVar("id", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Update Success", http.StatusOK, result))
	}
}

func (s *Controller) HandleDELETEMenus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.MenuService.HandleDELETEMenu(tools.GetPathVar("id", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Delete By ID Success", http.StatusOK, result))
	}
}
