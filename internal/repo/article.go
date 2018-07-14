package repo

import (
	"github.com/gambarini/articleapi/internal/model"
	"gopkg.in/mgo.v2/bson"
)

func (repo *ArticleRepository) Store(article model.Article) error {

	session, collection := repo.articleCollection()

	defer session.Close()

	err := collection.Insert(article)

	if err != nil {
		return err
	}

	return nil
}

func (repo *ArticleRepository) Find(id string) (article model.Article, err error) {

	session, collection := repo.articleCollection()

	defer session.Close()

	err = collection.Find(bson.M{"_id": id}).One(&article)

	if err != nil {
		return article, err
	}

	return article, nil
}
