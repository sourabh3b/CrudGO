package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/unrolled/render"
	"github.com/CrudGO/user"
	"github.com/gorilla/mux"
	"encoding/json"
)

//TestRoute - test route
func TestRoute(w http.ResponseWriter, r *http.Request) {
	//render := render.New()
	fmt.Fprint(w, "Hello World !")
	//render.JSON(w, http.StatusOK, nil)
	return
}

/*
GetAllUsers - §  Return a list of all users as a JSON string
e.g.:
{ "users":
    [{ "username": "jsmith", "displayName": "John Smith", "department": "Sales" },
    { "username": "jdoe", "displayName": "John Doe", "department": "Development" }]
}
*/

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	render := render.New()
	users,err := user.GetAllUsers()

	//check for error, if there is error then return 500 status code with empty user, else return status 200 with all users
	if(err != nil){
		render.JSON(w,http.StatusInternalServerError, users)
	}else{
		render.JSON(w,http.StatusOK, users)
	}

}

/*
GetUserByName - Return the data of a particular user as a JSON string or status code 404 if not found
e.g.:
 { "displayName": "John Smith", "department": "Sales" }
*/
func GetUserByName(w http.ResponseWriter, r *http.Request) {
	render := render.New()

	//getting username from the API url
	vars := mux.Vars(r)
	username := vars["username"]

	fmt.Fprintf(w, "You've requested the username: %s\n", username)


	log.Info("username >>>>>.",username)
	user,err := user.GetUserByName(username);

	//check for error, if there is error then return 500 status code with empty user, else return status 200 with all users
	if(err != nil){
		render.JSON(w,http.StatusForbidden, user)
	}else{
		render.JSON(w,http.StatusOK, user)
	}


	return
}

//InsertUser - test route
func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	render := render.New()

	inputUser := user.User{}

	insertAPIResponse := user.Response{}


	//decoding the request into team, so that it can be used to save the team details
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		insertAPIResponse.Message = "Invalid input : Unable to parse input body"
		insertAPIResponse.Status = http.StatusBadRequest
		render.JSON(w,http.StatusBadRequest, insertAPIResponse)
		return
	}


	err, exist  := user.InsertNewUser(inputUser)

	//check if user already exist in the database
	if exist {
		insertAPIResponse.Message = "User already found in the database"
		insertAPIResponse.Status = http.StatusConflict
		render.JSON(w,http.StatusConflict, insertAPIResponse)
		return
	}

	//check for error returned from InsertNewUser function
	if(err != nil){
		insertAPIResponse.Message = "Unable to insert user into database"
		insertAPIResponse.Status = http.StatusForbidden
		render.JSON(w,http.StatusForbidden, insertAPIResponse)
	}else{
		insertAPIResponse.Message = "Inserted user into database"
		insertAPIResponse.Status = http.StatusOK
		render.JSON(w,http.StatusOK, insertAPIResponse)
	}


	return
}

//DeleteUser - test route
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//render := render.New()
	fmt.Fprint(w, "Hello World !")
	//render.JSON(w, http.StatusOK, nil)
	return
}


func main() {

	log.Info("Started CRUD following are the APIs available....")

	//router for all APIs
	http.HandleFunc("/test", TestRoute)
	http.HandleFunc("/users", GetAllUsers) //Get all users present in the database
	http.HandleFunc("/insert", InsertUser) //inserts a user into database
	http.HandleFunc("/users/{username}", GetUserByName) //get user by user name
	//http.HandleFunc("/users/{username}", DeleteUser) //Deletes user from database

	//listening on 8889 port number
	http.ListenAndServe(":8889", nil)
}
