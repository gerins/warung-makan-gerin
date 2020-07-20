package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"warung_makan_gerin/utils/message"
	"warung_makan_gerin/utils/tools"
)

type Controller struct {
	UserService UserServiceInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{UserService: NewUserService(db)}
}

func (s *Controller) HandleGETAllUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		Users, err := s.UserService.GetUsers()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Search All Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Search All Success", http.StatusOK, Users))
	}
}

func (s *Controller) HandleUserLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userLogin User
		err := json.NewDecoder(r.Body).Decode(&userLogin)
		log.Print(userLogin.Username)
		log.Print(userLogin.Password)

		userWithToken, err := s.UserService.HandleUserLogin(userLogin)
		if err != nil {
			log.Print(`gagal login`)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Login Failed", http.StatusBadRequest, err.Error()))
			return
		}

		// http.SetCookie(w, &http.Cookie{
		// 	Name:     "token",
		// 	Value:    "token",
		// 	Path:     "/",
		// 	Expires:  time.Now().Add(0 * time.Second),
		// 	HttpOnly: true,
		// })

		// http.SetCookie(w, &http.Cookie{
		// 	Name:    "username",
		// 	Value:   "username",
		// 	Path:    "/",
		// 	Expires: time.Now().Add(0 * time.Second),
		// })

		// http.SetCookie(w, &http.Cookie{
		// 	Name:     "token",
		// 	Value:    userWithToken.Token,
		// 	Path:     "/",
		// 	Expires:  time.Now().Add(120 * time.Second),
		// 	HttpOnly: true,
		// })

		// http.SetCookie(w, &http.Cookie{
		// 	Name:    "username",
		// 	Value:   userWithToken.User.Username,
		// 	Path:    "/",
		// 	Expires: time.Now().Add(120 * time.Second),
		// })

		log.Print(`Success Login`)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Login Success", http.StatusOK, userWithToken))
		return
	}
}

func (s *Controller) HandleRegisterNewUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data User
		tools.Parser(r, &data)

		result, err := s.UserService.HandleRegisterNewUser(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Register Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Register Success", http.StatusOK, "Selamat Datang "+result.Username))
	}
}

func (s *Controller) HandleUPDATEUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data User
		tools.Parser(r, &data)

		result, err := s.UserService.HandleUPDATEUser(tools.GetPathVar("id", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Update Success", http.StatusOK, result))
	}
}

func (s *Controller) HandleDELETEUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.UserService.HandleDELETEUser(tools.GetPathVar("id", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Delete By ID Success", http.StatusOK, result))
	}
}

func (s *Controller) UserLogOut() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "token",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Second),
		})

		http.SetCookie(w, &http.Cookie{
			Name:    "username",
			Value:   "username",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Second),
		})

		w.WriteHeader(http.StatusOK)
	}
}
