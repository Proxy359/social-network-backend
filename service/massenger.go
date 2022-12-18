package service

import (
	"encoding/json"
	"gomod/storage"
	"io/ioutil"
	"net/http"
)

//данный метод должен вызываться дважды - вервый раз он обновляет переписку первого юзера, а затем второго
func (u *TaskService) WriteMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var firstUserInfo storage.FullUserInfo
		if err := json.Unmarshal(content, &firstUserInfo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = u.TaskServ.WriteMessage(firstUserInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("сообщение успешно отправлено")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

//данный метод должен вызываться дважды - вервый раз он обновляет сообщение первого юзера, а затем второго
func (u *TaskService) UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var firstUserInfo storage.FullUserInfo
		if err := json.Unmarshal(content, &firstUserInfo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = u.TaskServ.UpdateMessage(firstUserInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("сообщение успешно отредактировано")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

//данный метод должен вызываться дважды - вервый раз он обновляет сообщение первого юзера, а затем второго
func (u *TaskService) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var firstUserInfo storage.FullUserInfo
		if err := json.Unmarshal(content, &firstUserInfo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = u.TaskServ.DeleteMessage(firstUserInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, _ := json.Marshal("сообщение успешно отредактировано")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
