package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"WebProject/app/model"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//возвращаем список клиенту в формате JSON
	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
