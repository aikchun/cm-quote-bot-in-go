package quotes

import (
	"github.com/aikchun/cm-quote-bot-in-go/db"
	"time"
)

type Quote struct {
	tableName struct{}  `pg:"quote_quote"`
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	UserID    int64     `json:"userId"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUserQuotes(q *[]Quote, user_id int64) error {
	d := db.NewDB()
	defer d.Close()

	return d.Model(q).Where("user_id = ?", user_id).Select()
}

func GetUserLatestQuote(q *Quote, user_id int64) error {
	d := db.NewDB()
	defer d.Close()

	return d.Model(q).Where("user_id = ?", user_id).Order("created_at DESC").Limit(1).Select()
}

func SaveQuote(q *Quote) error {
	d := db.NewDB()
	defer d.Close()

	_, err := d.Model(q).Insert()
	return err
}