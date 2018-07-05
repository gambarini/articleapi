package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
)

func HandleTags(r *mux.Router, srv *Server) {

	r.HandleFunc("/tags/{tagName}/{date}", srv.getTag).Methods("GET")

}


func (srv *Server) getTag(writer http.ResponseWriter, request *http.Request) {

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
