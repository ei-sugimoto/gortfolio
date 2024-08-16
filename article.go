package main

import (
	"context"
	"log"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ei-sugimoto/gortfolio/ent"
	"github.com/ei-sugimoto/gortfolio/ent/articlehistory"
	"github.com/ei-sugimoto/gortfolio/services"
)

func GetArticle(client *ent.Client, ctx context.Context) ([]*ent.Article, error) {

	articleHis, err := client.ArticleHistory.Query().Order(articlehistory.ByCreatedAt(sql.OrderDesc())).WithArticle().First(ctx)
	if err != nil {
		log.Println(err)
	}
	now := time.Now()
	if articleHis == nil || now.Sub(articleHis.CreatedAt) > 30*24*time.Hour {
		qiita := services.NewQiita("")
		items, err := qiita.GetItems()
		if err != nil {
			log.Println(err)
		}
		newArticleHis, err := client.Debug().ArticleHistory.Create().SetCreatedAt(now).Save(ctx)
		if err != nil {
			log.Println(err)
		}
		articles, err := client.Article.MapCreateBulk(items, func(ac *ent.ArticleCreate, i int) {
			ac.SetTitle(items[i].Title)
			ac.SetURL(items[i].URL)
			ac.SetOwner(newArticleHis)
		}).Save(ctx)
		if err != nil {
			log.Println(err)
		}

		return articles, nil
	}
	return articleHis.QueryArticle().All(ctx)
}
