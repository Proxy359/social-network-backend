package storage

import (
	"errors"
	"time"
)

type Massage struct {
	AuthorID           int        `json:"author_id"`
	MassageDescriprion string     `json:"massage_descriprion"`
	SendingTime        time.Time  `json:"sending_time"`
}

func (u *TaskStorage) WriteMessage(userInfo FullUserInfo) error {
	var queryResult FullUserInfo
	err := u.TaskStore.QueryRow("SELECT user_chat_list FROM users WHERE user_id = ?", userInfo.SomeLogInfo.userID).Scan(queryResult.ChatList)
	if err != nil {
		return errors.New("ошибка при получении списка чатов")
	}

	newMassage := userInfo.ChatList[userInfo.FriendsList[0]] //id собеседника берём из получаемой переменной в первом элементе списка друзей
	pastMassages := queryResult.ChatList[userInfo.FriendsList[0]]

	newVersChat := append(newMassage, pastMassages...)
	queryResult.ChatList[userInfo.FriendsList[0]] = newVersChat
	queryResult.ChatList[userInfo.SomeLogInfo.userID] = newVersChat

	insert, err := u.TaskStore.Query("UPDATE user_chat_list SET user_chat_list = ? WHERE user_id = ?",
		queryResult, userInfo.SomeLogInfo.userID)
	if err != nil {
		return errors.New("ошибка при отправке сообщения")
	}
	defer insert.Close()
	return nil
}

func (u *TaskStorage) UpdateMessage(userInfo FullUserInfo) error {
	var queryResult FullUserInfo
	err := u.TaskStore.QueryRow("SELECT user_chat_list FROM users WHERE user_id = ?", userInfo.SomeLogInfo.userID).Scan(queryResult.ChatList)
	if err != nil {
		return errors.New("ошибка при получении списка чатов")
	}

	for i, cycleMassage := range queryResult.ChatList[queryResult.SomeLogInfo.userID] {
		if cycleMassage.AuthorID == userInfo.ChatList[userInfo.FriendsList[0]][0].AuthorID &&
			cycleMassage.SendingTime == userInfo.ChatList[userInfo.FriendsList[0]][0].SendingTime {
			queryResult.ChatList[queryResult.SomeLogInfo.userID][i].MassageDescriprion = userInfo.ChatList[userInfo.FriendsList[0]][0].MassageDescriprion

			insert, err := u.TaskStore.Query("UPDATE user_chat_list SET user_chat_list = ? WHERE user_id = ?",
				queryResult, userInfo.SomeLogInfo.userID)
			if err != nil {
				return errors.New("ошибка при отправке сообщения")
			}
			defer insert.Close()
			return nil
		}
	}
	return errors.New("ошибка при редактировании сообщения")
}

func (u *TaskStorage) DeleteMessage(userInfo FullUserInfo) error {
	var queryResult FullUserInfo
	err := u.TaskStore.QueryRow("SELECT user_chat_list FROM users WHERE user_id = ?", userInfo.SomeLogInfo.userID).Scan(queryResult.ChatList)
	if err != nil {
		return errors.New("ошибка при получении списка чатов")
	}

	for i, cycleMassage := range queryResult.ChatList[queryResult.SomeLogInfo.userID] {
		if cycleMassage.AuthorID == userInfo.ChatList[userInfo.FriendsList[0]][0].AuthorID &&
			cycleMassage.SendingTime == userInfo.ChatList[userInfo.FriendsList[0]][0].SendingTime {
			queryResult.ChatList[queryResult.SomeLogInfo.userID] =
				append(queryResult.ChatList[queryResult.SomeLogInfo.userID][:i], queryResult.ChatList[queryResult.SomeLogInfo.userID][i+1:]...)

			insert, err := u.TaskStore.Query("UPDATE user_chat_list SET user_chat_list = ? WHERE user_id = ?",
				queryResult, userInfo.SomeLogInfo.userID)
			if err != nil {
				return errors.New("ошибка при удалении сообщения")
			}
			defer insert.Close()
			return nil
		}
	}
	return errors.New("ошибка при удалении сообщения сообщения")
}
