package main

import (
	"log"
	"github.com/gambarini/articleapi/internal/server"
	"github.com/gambarini/articleapi/internal/repo"
)

func main() {

	repo, err := repo.NewArticleRepository()

	if err != nil {
		log.Fatalf("Failed to connect db, %s", err)
	}

	defer repo.CleanUp()

	srv := server.NewServer(":8000", repo)

	log.Println("Listening on port 8000. Ctrl+C to stop")
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start server, %s", err)
	}

}