package app

import "github.com/jmoiron/sqlx"

func buildDB(dbType, address string) (*sqlx.DB, error) {
	client, err := sqlx.Connect(dbType, address) //client  客户端
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	return client, nil
}
