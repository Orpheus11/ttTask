package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"ttTask/model"
)

type relationStruct struct {
	Id    int64  `json:id`
	State string `json:state`
	Type  string `json:type`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome,it is Connor`s TanTan task!")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := model.GetUsers()
	if len(users) > 0 {
		json.NewEncoder(w).Encode(users)
		return
	}
	json.NewEncoder(w).Encode(make([]model.User, 0))
}

func AddUsers(w http.ResponseWriter, r *http.Request) {
	var user model.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	user.Type = "user"
	model.CreateUser(&user)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func GetRelationshipsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id_string := vars["user_id"]
	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		panic(err)
	}
	relationships, _ := model.GetRelationshipsByUserId(int64(user_id))
	if len(relationships) > 0 {
		rps := make([]relationStruct, 0)
		for _, relationship := range relationships {
			rp := relationStruct{}
			rp.Id = relationship.Other_user_id
			rp.State = relationship.State
			rp.Type = relationship.Type
			rps = append(rps, rp)
		}
		json.NewEncoder(w).Encode(rps)
		return
	}
	json.NewEncoder(w).Encode(make([]model.Relationship, 0))
}
func AddRelationships(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id_string := vars["user_id"]
	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		panic(err)
	}
	other_user_id_string := vars["other_user_id"]
	other_user_id, err := strconv.Atoi(other_user_id_string)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var rp relationStruct
	if err := json.Unmarshal(body, &rp); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	rp.Id = int64(other_user_id)
	rp.Type = "relationship"
	relationship := model.Relationship{User_id: int64(user_id), Other_user_id: int64(other_user_id), State: rp.State, Type: "relationship"}
	model.CreateUpdateRelationship(&relationship)
	otherRelationship, _ := model.GetRelationship(int64(other_user_id), int64(user_id))
	if relationship.State == "liked" && otherRelationship.State == "liked" {
		rp.State = "matched"
	}
	json.NewEncoder(w).Encode(rp)
}
