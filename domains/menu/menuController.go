package menu

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"warung_makan_gerin/utils/message"
	"warung_makan_gerin/utils/tools"

	"github.com/gorilla/mux"
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

		var page string = mux.Vars(r)["page"]
		var limit string = mux.Vars(r)["limit"]
		var status string = mux.Vars(r)["status"]
		var orderBy string = mux.Vars(r)["orderBy"]
		var sort string = mux.Vars(r)["sort"]
		var keyword string = mux.Vars(r)["keyword"]

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

		// ini syntax untuk dapetin file nya, uploadedFile itu file nya, handler itu untuk dapetin .jpg file nya
		// di postman pilih POST -> BODY -> Form-data -> key nya file dan value nya adalah file itu sendiri
		//////////////////////////////////////////////////////////////////////////////////////////
		r.ParseMultipartForm(1024) // ini untuk batesin file size nya biar maks 1 MB
		uploadedFile, handler, err := r.FormFile("file")
		if err != nil {
			log.Println(`Error while parsing file`, err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(message.Respone("Upload Photos Failed", http.StatusInternalServerError, err.Error()))
			return
		}
		defer uploadedFile.Close()
		//////////////////////////////////////////////////////////////////////////////////////////

		var data Menu
		data.MenuName = r.FormValue("menuname")
		data.Category = r.FormValue("category")
		harga, _ := strconv.Atoi(r.FormValue("harga"))
		stock, _ := strconv.Atoi(r.FormValue("stock"))

		data.Harga = harga
		data.Stock = stock
		// identifyUser, _ := r.Cookie("user")
		identifyUser := "Gerin"

		fmt.Println(data)
		fmt.Println(handler.Filename)

		// LANJUT ke service untuk di proses
		result, err := s.MenuService.HandlePOSTMenu(data, identifyUser, uploadedFile, handler)
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

// Ini function untuk serve yang ada di harddisk, function ini di panggil di file menuRoute.go
func (s *Controller) HandleGetImages() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}

		fileId := tools.GetPathVar("id", r)
		fileLocation := filepath.Join(dir, "files", fileId)

		w.Header().Set("Content-Type", "image/jpeg")
		http.ServeFile(w, r, fileLocation)
	}
}
