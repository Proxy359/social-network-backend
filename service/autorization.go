package service

import (
	"encoding/json"
	"gomod/storage"
	"io/ioutil"
	"net/http"
)

type UserService struct {
	UserServ *storage.UserStorage
}

func (u *UserService) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var userInfo storage.FullUserInfo
		if err := json.Unmarshal(content, &userInfo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = u.UserServ.CreateUser(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("пользователь успешно зарегистрирован")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *UserService) UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var userInfo storage.FullUserInfo
		if err := json.Unmarshal(content, &userInfo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = u.UserServ.UserAuthorization(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("пользователь успешно анторизован")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
