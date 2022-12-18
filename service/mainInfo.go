package service

import (
	"encoding/json"
	"gomod/storage"
	"io/ioutil"
	"net/http"
)

type TaskService struct {
	TaskServ *storage.TaskStorage
}

func (u *TaskService) EditNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
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

		err = u.TaskServ.EditName(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("имя успешно обновлено")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *TaskService) EditLastNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
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

		err = u.TaskServ.EditLastName(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("фамилия успешно обновлена")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *TaskService) EditSexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
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

		err = u.TaskServ.EditSex(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("пол успешно обновлен")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *TaskService) EditFamilyStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
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

		err = u.TaskServ.EditFamilyStatus(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("семейный статус успешно обновлен")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *TaskService) EditCityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
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

		err = u.TaskServ.EditFamilyStatus(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("город успешно обновлен")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *TaskService) AddUnclesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
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

		err = u.TaskServ.AddUncles(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("список 'бабушки,дедушки' успешно обновлен")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *TaskService) DeleteUnclesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
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

		err = u.TaskServ.DeleteUncles(userInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("список 'бабушки,дедушки' успешно обновлен")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
