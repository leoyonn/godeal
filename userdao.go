package godeal

import (
	"errors"
)

const (
	sqlReg        = "INSERT INTO `user` (`account`, `phone`, `token`) VALUES (?, ?, ?)"
	sqlSetPass    = "UPDATE `user` SET `password` = ? WHERE `account` = ?"
	sqlDelId      = "DELETE FROM `user` WHERE `id` = ?"
	sqlDelAccount = "DELETE FROM `user` WHERE `account` = ?"
	sqlQuery      = "SELECT * FROM `user` WHERE `account` = ?"
)

func reg(phone string, token string) (int64, error) {
	stmt, e := db.Prepare(sqlReg)
	if e != nil {
		Error("Prepare create user", user, e)
		return 0, e
	}
	defer stmt.Close()
	result, e := stmt.Exec(phone, phone, token)
	if e != nil {
		Error("Executing create user", phone, e)
		return 0, e
	}
	id, e := result.LastInsertId()
	if e != nil {
		Warn("Get last inserted id", phone, e);
		return 0, e
	}
	Debug("Update id in user", user)
	return id, nil
}

func setPass(account string, token string, newPass string, oldPass string) error {
	user, e := query(account)
	if e != nil {
		return e
	}
	if token != user.Token {
		return errors.New("Token mismatch")
	}
	if (len(user.Pass) != 0 || len(oldPass) != 0) && oldPass != user.Pass {
		return errors.New("Password doesn't match")
	}
	stmt, e := db.Prepare(sqlSetPass)
	if e != nil {
		Error("Prepare set password", user, pass, e)
		return e
	}
	defer stmt.Close()
	result, e := stmt.Exec(newPass, account)
	if e != nil {
		Error("Executing set password", user, pass, e)
		return e
	}
	rows, e := result.RowsAffected()
	if e != nil {
		Warn("Get rows affected by setPass", user, pass, e);
		return e
	}
	Debug("Set pass", user, pass, rows)
	return nil
}

func update(user*User) error {
	return nil
}

func query(account string) (*User, error) {
	stmt, e := db.Prepare(sqlQuery)
	if e != nil {
		Error("Prepare query user", user, e)
		return nil, e
	}
	defer stmt.Close()
	row := stmt.QueryRow(account)
	if e != nil {
		Error("Executing query user", user, e)
		return nil, e
	}
	var user User
	if e := row.Scan(&user.Id, &user.Account, &user.Name, &user.Desc, &user.Gender, &user.Email,
		&user.Phone, &user.Avatar, &user.Pass, &user.Token, &user.CreateAt); e != nil {
		return nil, e
	}
	Debug("Update id in user", user)
	return &user, nil
}

func login(account string, pass string) (*User, error) {
	user, e := query(account)
	if e != nil {
		return nil, e
	}
	if user.Pass != pass {
		return nil, errors.New("Password wrong")
	}
	return user, nil
}

func DelById(id int64) error {
	stmt, e := db.Prepare(sqlDelId)
	if e != nil {
		Error("Prepare delete user by id", id, e)
		return e
	}
	defer stmt.Close()
	result, e := stmt.Exec(id)
	if e != nil {
		Error("Execute delete user by id", id, e)
		return e
	}
	rows, e := result.RowsAffected()
	if e != nil {
		Warn("Get rows affected by delete", id, e);
		return e
	}
	Debug("Delete by id", id, rows)
	return nil
}

func DelByAccount(account string) error {
	stmt, e := db.Prepare(sqlDelAccount)
	if e != nil {
		Error("Prepare delete user by account", account, e)
		return e
	}
	defer stmt.Close()
	result, e := stmt.Exec(account)
	if e != nil {
		Error("Execute delete user by account", account, e)
		return e
	}
	rows, e := result.RowsAffected()
	if e != nil {
		Warn("Get rows affected by delete", account, e);
		return e
	}
	Debug("Delete by account", account, rows)
	return nil
}

