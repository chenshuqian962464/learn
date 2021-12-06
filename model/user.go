package model

type User struct {
	ID        int64  `db:"id"`
	Firstname string `db:"firstname"`
	Age       int    `db:"age"`
	High      int    `db:"high"`
}
