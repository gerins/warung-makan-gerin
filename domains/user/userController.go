package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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
		userLogin.Username = tools.GetPathVar("user", r)
		userLogin.Password = tools.GetPathVar("password", r)

		User, err := s.UserService.HandleUserLogin(userLogin)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Login Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Login Success", http.StatusOK, User))
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

func (s *Controller) LoginPage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, `
		<html lang="en">
		  <head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<title>User Login</title>
		  </head>
		  <body>
			<h1>User Login</h1>
			<form action="">
			  <label for="inputUsername">Username</label>
			  <input
				required
				maxlength="16"
				minlength="4"
				name="user"
				type="text"
				id="inputUsername"
				placeholder="Username"
			  />
			  <label for="inputPassword">Password</label>
			  <input
				required
				maxlength="16"
				minlength="4"
				name="password"
				type="password"
				id="inputPassword"
				placeholder="Password"
			  />
			  <button type="submit">Login</button>
			</form>
		  </body>
		</html>
		
		`)
	}
}
