package repo

import (
	"github.com/gambarini/articleapi/internal/model"
	"gopkg.in/mgo.v2/bson"
)

func (repo *ArticleRepository) FindTag(tag, date string) (resultTag model.Tag, err error) {

	session, collection := repo.articleCollection()

	defer session.Close()

	query := bson.M{"$and":
	[]bson.M{
		{"tags": bson.M{"$in": []string{tag}}},
		{"date": date},
	},
	}

	var articles []model.Article
	count := make(chan int)

	go func() {
		c, _ := collection.Find(query).Count()
		count <- c
		close(count)
	}()

	err = collection.Find(query).Sort("-datetime").Limit(10).All(&articles)

	if err != nil {
		return resultTag, err
	}

	resultTag = model.Tag{
		Tag:         tag,
		Articles:    []string{},
		RelatedTags: []string{},
	}

	rTgsMap := make(map[string]string)

	for _, art := range articles {

		resultTag.Articles = append(resultTag.Articles, art.ID)

		for _, t := range art.Tags {
			if t != tag {
				rTgsMap[t] = t
			}
		}
	}

	for key := range rTgsMap {
		resultTag.RelatedTags = append(resultTag.RelatedTags, key)
	}

	resultTag.Count = <-count

	return resultTag, nil
}
