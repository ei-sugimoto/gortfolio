package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Qiita struct {
	Token string
}

func NewQiita(token string) *Qiita {
	return &Qiita{Token: token}
}

type Item struct {
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Qiita) GetItems() ([]Item, error) {
	res, err := http.Get("https://qiita.com/api/v2/users/ei-sugimoto/items?page=1&per_page=10")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	data := make([]Item, 0)

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}
