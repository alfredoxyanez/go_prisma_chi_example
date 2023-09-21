package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alfredoxyanez/go_prisma_chi_example/database"
	"github.com/alfredoxyanez/go_prisma_chi_example/prisma/db"
	"github.com/go-chi/chi"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient
	allUsers, err := pClient.Client.User.FindMany().Exec(pClient.Context)
	if err != nil {
		fmt.Println("Cannot fetch users")
		return

	}
	usersMap := make(map[string]interface{})
	usersMap["users"] = allUsers
	out, err := json.MarshalIndent(usersMap, "", "\t")
	if err != nil {
		fmt.Println("Error Creating JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		fmt.Println("Error Creating Writing")
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pClient := database.PClient
	user, err := pClient.Client.User.FindUnique(
		db.User.ID.Equals(id)).Exec(pClient.Context)
	if err != nil {
		fmt.Println("Cannot find user")
		return
	}
	out, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		fmt.Println("Error Creating JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		fmt.Println("Error Creating Writing")
		return
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userResp db.UserModel
	err := json.NewDecoder(r.Body).Decode(&userResp)
	if err != nil {
		fmt.Println("Cannot decode user")
		return
	}
	pClient := database.PClient
	user, err := pClient.Client.User.CreateOne(
		db.User.Email.Set(userResp.Email),
		db.User.Password.Set(userResp.Password), // You should hash password !!!
		db.User.Firstname.Set(userResp.Firstname),
		db.User.Lastname.Set(userResp.Lastname),
	).Exec(pClient.Context)
	if err != nil {
		fmt.Println("Cannot create user")
		return
	}
	out, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		fmt.Println("Error Creating JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		fmt.Println("Error Creating Writing")
	}

}
