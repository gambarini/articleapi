package repo

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

type (
	ArticleRepository struct {
		session *mgo.Session
	}
)

func NewArticleRepository() (*ArticleRepository, error) {

	session, err := mgo.Dial("localhost")

	if err != nil {
		return nil, fmt.Errorf("fail to dial to mongodb, %s", err)
	}

	return &ArticleRepository{
		session: session,
	}, nil

}

func (repo *ArticleRepository) articleCollection() (*mgo.Session, *mgo.Collection) {

	session := repo.session.Copy()

	session.SetMode(mgo.Monotonic, true)

	return session, session.DB("articleapi").C("article")
}

func (repo *ArticleRepository) CleanUp() {
	repo.session.Close()
}
