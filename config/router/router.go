package router

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// CreateRouter for creating new Route
func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

var serverHost, serverPort string

// StartServer routing
func StartServer(r *mux.Router) {
	configServer()
	log.Println("Server Start at http://" + serverHost + ":" + serverPort)
	http.ListenAndServe(serverHost+":"+serverPort, r)
}

func configServer() {
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var readResults []string
	for scanner.Scan() {
		configData := strings.Split(scanner.Text(), "=")[1]
		readResults = append(readResults, configData)
	}

	serverHost = readResults[6]
	serverPort = readResults[7]
}
