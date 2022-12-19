package storage

import (
	"database/sql"
	"errors"
)

type MainInfo struct {
	Name         string  `json:"name"`
	LastName     string  `json:"last_name"`
	Sex          string  `json:"sex"`
	FamilyStatus string  `json:"family_status"`
	Birthday     [3]int  `json:"birthday"`
	City         string  `json:"city"`
	Uncles       []int  `json:"uncles"`
}

type TaskStorage struct {
	TaskStore *sql.DB
}

func (u *TaskStorage) EditName(userInfo FullUserInfo) error {
	insert, err := u.TaskStore.Query("UPDATE name SET name = ? WHERE user_id = ?", userInfo.SomeMainInfo.Name, userInfo.SomeLogInfo.userID)
	if err != nil {
		return errors.New("ошибка при обновлении имени")
	}
	defer insert.Close()
	return nil
}

func (u *TaskStorage) EditLastName(userInfo FullUserInfo) error {
	insert, err := u.TaskStore.Query("UPDATE last_name SET last_name = ? WHERE user_id = ?", userInfo.SomeMainInfo.LastName, userInfo.SomeLogInfo.userID)
	if err != nil {
		return errors.New("ошибка при обновлении фамилии")
	}
	defer insert.Close()
	return nil
}

func (u *TaskStorage) EditSex(userInfo FullUserInfo) error {
	if userInfo.SomeMainInfo.Sex == "male" || userInfo.SomeMainInfo.Sex == "female" {
		insert, err := u.TaskStore.Query("UPDATE sex SET sex = ? WHERE user_id = ?", userInfo.SomeMainInfo.Sex, userInfo.SomeLogInfo.userID)
		if err != nil {
			return errors.New("ошибка при обновлении пола")
		}
		defer insert.Close()
		return nil
	}
	return errors.New("введите корректный пол")
}

func (u *TaskStorage) EditFamilyStatus(userInfo FullUserInfo) error {
	var queryResult MainInfo
	var maleFamilyStatus = [9]string{"Не женат", "Встречаюсь", "Помолвлен", "Женат", "В гражданском браке", "Помолвлен", "Влюблен", "В активном поиске", "Всё сложно"}
	var femaleFamilyStatus = [9]string{"Не замужем", "Встречаюсь", "Помолвлена", "Замужем", "В гражданском браке", "Помолвлена", "Влюблена", "В активном поиске", "Всё сложно"}

	err := u.TaskStore.QueryRow("SELECT sex FROM users WHERE user_id = ?", userInfo.SomeLogInfo.userID).Scan(&queryResult.Sex)
	if err != nil {
		return errors.New("ошибка при проверке пола юзера. перед выбором семейного положения, пожалуйста, укажите ваш пол")
	}

	if queryResult.Sex == "male" {
		for _, checkFamilyStatus := range maleFamilyStatus {
			if checkFamilyStatus == userInfo.SomeMainInfo.FamilyStatus {
				insert, err := u.TaskStore.Query("UPDATE family_status SET family_status = ? WHERE user_id = ?", userInfo.SomeMainInfo.FamilyStatus, userInfo.SomeLogInfo.userID)
				if err != nil {
					return errors.New("ошибка при обновлении семейного статуса")
				}
				defer insert.Close()
				return nil
			}
		}

		if queryResult.Sex == "female" {
			for _, checkFamilyStatus := range femaleFamilyStatus {
				if checkFamilyStatus == userInfo.SomeMainInfo.FamilyStatus {
					insert, err := u.TaskStore.Query("UPDATE family_status SET family_status = ? WHERE user_id = ?", userInfo.SomeMainInfo.FamilyStatus, userInfo.SomeLogInfo.userID)
					if err != nil {
						return errors.New("ошибка при обновлении семейного статуса")
					}
					defer insert.Close()
					return nil
				}
			}
		}
	}
	return errors.New("ошибка при проверке пола юзера. перед выбором семейного положения, пожалуйста, укажите ваш пол")
}

func (u *TaskStorage) EditCity(userInfo FullUserInfo) error {
	insert, err := u.TaskStore.Query("UPDATE сity SET сity = ? WHERE user_id = ?", userInfo.SomeMainInfo.City, userInfo.SomeLogInfo.userID)
	if err != nil {
		return errors.New("ошибка при обновлении города")
	}
	defer insert.Close()
	return nil
}

func (u *TaskStorage) AddUncles(userInfo FullUserInfo) error {
	var queryResult FullUserInfo
	err := u.TaskStore.QueryRow("SELECT friends_list, uncles FROM users WHERE user_id =?", userInfo.SomeLogInfo.userID).Scan(&queryResult.FriendsList, queryResult.SomeMainInfo.Uncles)
	if err != nil {
		return errors.New("ошибка при получении списка друзей/родственников")
	}

	for _, uncleID := range queryResult.SomeMainInfo.Uncles {
		if uncleID == userInfo.SomeMainInfo.Uncles[0] { //по задумке ID родственника "кладется" в пустой слайс единственным элементом
			return errors.New("данный пользователь уже находится в списке 'дедушки,бабушки'")
		}
	}

	queryResult.SomeMainInfo.Uncles = append(queryResult.SomeMainInfo.Uncles, userInfo.SomeMainInfo.Uncles[0])
	for _, friendID := range queryResult.FriendsList {
		if friendID == userInfo.SomeMainInfo.Uncles[0] {
			insert, err := u.TaskStore.Query("UPDATE uncles SET uncles = ? WHERE user_id = ?", queryResult.SomeMainInfo.Uncles, userInfo.SomeLogInfo.userID)
			if err != nil {
				return errors.New("ошибка при обновлении списка 'дедушки,бабушки'")
			}
			defer insert.Close()
			return nil
		}
	}
	return errors.New("нельзя указать родственника, если его нет в списке друзей")
}

func (u *TaskStorage) DeleteUncles(userInfo FullUserInfo) error {
	var queryResult FullUserInfo
	err := u.TaskStore.QueryRow("SELECT friends_list, uncles FROM users WHERE user_id =?", userInfo.SomeLogInfo.userID).Scan(&queryResult.FriendsList, queryResult.SomeMainInfo.Uncles)
	if err != nil {
		return errors.New("ошибка при получении списка друзей/родственников")
	}

	checkUnclePos := -1
	for pos, uncleID := range queryResult.SomeMainInfo.Uncles {
		if uncleID == userInfo.SomeMainInfo.Uncles[0] {
			checkUnclePos = pos
		}
	}
	if checkUnclePos == -1 {
		return errors.New("данного пользователя нет в списке родственников")
	}

	newUncleList := append(queryResult.SomeMainInfo.Uncles[:checkUnclePos], queryResult.SomeMainInfo.Uncles[checkUnclePos+1:]...)
	insert, err := u.TaskStore.Query("UPDATE uncles SET uncles = ? WHERE user_id = ?", newUncleList, userInfo.SomeLogInfo.userID)
	if err != nil {
		return errors.New("ошибка при обновлении списка 'дедушки,бабушки'")
	}
	defer insert.Close()
	return nil
}
