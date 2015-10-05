package models

import (
	"time"
	"log"
	"database/sql"
)

type Account struct {
	Id				int
	EMail 			string
	Password		string
	DisplayName 	string
	CreateTime 		time.Time
	LastLoginTime 	time.Time
	ShowEmail		bool
	Level			int
}

func CreateNewAccount2(account *Account) error {
	var lastInsertId int
	var db *sql.DB
	var err error
	db, err = BorrowDBConn()
	if err != nil {
		return err
	}
	
	defer ReleaseDBConn(db)
	err = db.QueryRow(
		`INSERT INTO user2(
			email,
			name,
			password,
			showemail,
			lastlogin
		) VALUES($1,$2,$3,$4,($5)) returning id;`,
	 		account.EMail,
			account.DisplayName,
			account.Password,
			account.ShowEmail,
			account.LastLoginTime).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	
	log.Printf("Create New Account with Id:%d, name:%s, email:%s\n",
		lastInsertId, account.DisplayName, account.EMail)
	return nil
}

func CreateNewAccount(account *Account) error {
	var lastInsertId int
	var db *sql.DB
	var err error
	db, err = BorrowDBConn()
	if err != nil {
		return err
	}
	
	defer ReleaseDBConn(db)
	err = db.QueryRow(
		`INSERT INTO users(
			email,
			dispalyname,
			userpassword,
			createtime,
			lastlogin,
			showemail,
			userlevel
		) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;`,
	 		account.EMail,
			account.DisplayName,
			account.Password,
			account.CreateTime,
			account.LastLoginTime,
			account.ShowEmail,
			account.Level).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	
	log.Printf("Create New Account with Id %d\n", lastInsertId)
	return nil
}

func QueryAccountByEmail(email string) (*Account, error) {
	
	var db *sql.DB
	var err error
	db, err = BorrowDBConn()
	if err != nil {
		return nil, err
	}
	
	defer ReleaseDBConn(db)
	rows, err := db.Query("SELECT * FROM users where email=$1", email)
	if err != nil {
		return nil, err
	}
	
	var account Account
	for rows.Next() {
        err = rows.Scan(&account.Id,
			&account.EMail, &account.DisplayName, &account.Password,
			&account.CreateTime, &account.LastLoginTime, &account.ShowEmail,
			&account.Level)
        if err != nil {
			return nil, err
		}
		
		return &account, nil
    }
	
	return nil, nil
}