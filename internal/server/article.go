package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/gambarini/articleapi/internal/model"
	"time"
	"github.com/satori/go.uuid"
	"log"
)

func HandleArticle(srv *Server) {

	srv.Router().HandleFunc("/articles/{id}", srv.getArticle).Methods("GET")
	srv.Router().HandleFunc("/articles", srv.postArticle).Methods("POST")

}

func (srv *Server) getArticle(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	id := vars["id"]

	article, err := srv.Repo.Find(id)

	if err != nil {
		log.Print(err)
		writer.WriteHeader(404)
		return
	}

	payloadJSON, err := json.Marshal(&article)

	if err != nil {
		log.Print(err)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(payloadJSON)
}

func (srv *Server) postArticle(writer http.ResponseWriter, request *http.Request) {

	dec := json.NewDecoder(request.Body)

	var article model.Article
	err := dec.Decode(&article)

	if err != nil {
		log.Print(err)
		writer.WriteHeader(500)
		return
	}

	id, _ := uuid.NewV4()

	article.ID = id.String()
	article.DateTime = time.Now().UTC()
	article.Date = article.DateTime.Format("2006-01-02")

	err = srv.Repo.Store(article)

	if err != nil {
		log.Print(err)
		writer.WriteHeader(500)
		return
	}

	payloadJSON, err := json.Marshal(&article)

	if err != nil {
		log.Print(err)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(201)
	writer.Write(payloadJSON)

}
