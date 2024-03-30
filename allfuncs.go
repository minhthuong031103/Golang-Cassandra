package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Restful API using Go and Cassandra!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var NewUser User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "wrong data")
	}
	json.Unmarshal(reqBody, &NewUser)
	if err := Session.Query("INSERT INTO user(UserId, Username, Age) VALUES(?, ?, ?)",
		NewUser.UserId, NewUser.Username, NewUser.Age).Exec(); err != nil {
		fmt.Println("Error while inserting")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(NewUser, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM user").Iter()
	for iter.MapScan(m) {
		users = append(users, User{
			UserId:   m["userid"].(int),
			Username: m["username"].(string),
			Age:      m["age"].(int),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(users, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["UserId"]
	var users []User
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM user WHERE UserId=?", userId).Iter()
	for iter.MapScan(m) {
		users = append(users, User{
			UserId:   m["userid"].(int),
			Username: m["username"].(string),
			Age:      m["age"].(int),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(users, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}

func CountAllUsers(w http.ResponseWriter, r *http.Request) {

	var Count string
	err := Session.Query("SELECT count(*) FROM user").Scan(&Count)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s ", Count)

}

// func DeleteOneStudent(w http.ResponseWriter, r *http.Request) {
// 	userId := mux.Vars(r)["UserId"]
// 	if err := Session.Query("DELETE FROM students WHERE id = ?", StudentID).Exec(); err != nil {
// 		fmt.Println("Error while deleting")
// 		fmt.Println(err)
// 	}
// 	fmt.Fprintf(w, "deleted successfully the student num %s ", StudentID)
// }

// func DeleteAllStudents(w http.ResponseWriter, r *http.Request) {

// 	if err := Session.Query("TRUNCATE students").Exec(); err != nil {
// 		fmt.Println("Error while deleting all students")
// 		fmt.Println(err)
// 	}
// 	fmt.Fprintf(w, "deleted all successfully")

// }

// func UpdateStudent(w http.ResponseWriter, r *http.Request) {
// 	StudentID := mux.Vars(r)["id"]
// 	var UpdateStudent Student
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data properly")
// 	}
// 	json.Unmarshal(reqBody, &UpdateStudent)
// 	if err := Session.Query("UPDATE students SET firstname = ?, lastname = ?, age = ? WHERE id = ?",
// 		UpdateStudent.Firstname, UpdateStudent.Lastname, UpdateStudent.Age, StudentID).Exec(); err != nil {
// 		fmt.Println("Error while updating")
// 		fmt.Println(err)
// 	}
// 	fmt.Fprintf(w, "updated successfully")

// }
