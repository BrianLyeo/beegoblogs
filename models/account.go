package modules

import (
	"log"
	"database/sql"
)

type Account struct {
	Id				int
	EMail 			string
	Password		string
	DisplayName 	string
	CreateTime 		int64
	LastLoginTime 	int64
	ShowEmail		bool
	Level			int
}

func CreateNewAccount(account Account) error {
	var lastInsertId int
	err = db.QueryRow(
		`INSERT INTO users(
			email,
			name,
			password,
			create,
			lastlogin,
			showemail,
			level
		) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;`,
	 		account.EMail,
			account.DisplayName,
			account.Password,
			account.CreateTime,
			account.CreateTime,
			account.ShowEmail,
			account.Level
		).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	
	log.Println("Create New Account with Id %d\n", lastInsertId)
}