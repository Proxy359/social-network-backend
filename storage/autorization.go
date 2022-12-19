package storage

import (
	"database/sql"
	"errors"
)

type FullUserInfo struct {
	SomeLogInfo  LogInfo             `json:"some_log_info"`
	SomeMainInfo MainInfo            `json:"some_main_info"`
	FriendsList  []int               `json:"friends_list"`
	ChatList     map[int][]Massage   `json:"chat_list"`
}

type UserStorage struct {
	UserStore *sql.DB
}

type LogInfo struct {
	email    string  `json:"email"`
	password string  `json:"password"`
	userID   int     `json:"user_id"`
}

func (u *UserStorage) CreateUser(userInfo FullUserInfo) error {
	var queryResult LogInfo

	err := u.UserStore.QueryRow("SELECT email FROM users WHERE email = ?", userInfo.SomeLogInfo.email).Scan(&queryResult.email)
	if err != nil {
		return errors.New("ошибка при проверке имейла")
	}

	passwordLen := []rune(userInfo.SomeLogInfo.password)
	if len(passwordLen) < 5 {
		return errors.New("пароль должен состоять минимум из 5 символов")
	}

	insert, err := u.UserStore.Query("INSERT INTO users (email,password,nickname) VALUES (?,?)", userInfo.SomeLogInfo.email, userInfo.SomeLogInfo.password)
	if err != nil {
		return errors.New("ошибка при создании нового пользователя")
	}

	defer insert.Close()
	return nil
}

func (u *UserStorage) UserAuthorization(userInfo FullUserInfo) error {
	var queryResult LogInfo

	err := u.UserStore.QueryRow("SELECT email,password FROM users WHERE email = ?", userInfo.SomeLogInfo.email).Scan(&queryResult.email, &queryResult.password)
	if err != nil {
		return errors.New("ошибка при проверке имейла")
	}

	if userInfo.SomeLogInfo.password != queryResult.password {
		return errors.New("введен неверный пароль")
	}
	return nil
}
