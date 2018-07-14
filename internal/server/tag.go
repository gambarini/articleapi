package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
)

func HandleTags(srv *Server) {

	srv.Router().HandleFunc("/tags/{tagName}/{date}", srv.GetTag).Methods("GET")

}

func (srv *Server) GetTag(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	tagName := vars["tagName"]
	date := vars["date"]

	tag, err := srv.Repo.FindTag(tagName, date)

	if err != nil {
		log.Print(err)
		writer.WriteHeader(500)
		return
	}

	if len(tag.Articles) == 0 {
		writer.WriteHeader(404)
		return
	}

	payloadJSON, err := json.Marshal(&tag)

	if err != nil {
		log.Print(err)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(payloadJSON)

}
