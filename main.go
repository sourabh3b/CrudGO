package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/unrolled/render"
	"github.com/CrudGO/user"
	"github.com/gorilla/mux"
	"encoding/json"
)


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


	exist, err  := user.InsertNewUser(inputUser)

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
	render := render.New()

	//getting username from the API url
	vars := mux.Vars(r)
	username := vars["username"]



	err := user.DeleteUser(username)

	deleteAPIResponse := user.Response{}

	if(err != nil){
		deleteAPIResponse.Message = "Unable to delete user into database"
		deleteAPIResponse.Status = http.StatusNotFound
		render.JSON(w,http.StatusNotFound, deleteAPIResponse)
	}else{
		deleteAPIResponse.Message = "Deleted user from database"
		deleteAPIResponse.Status = http.StatusOK
		render.JSON(w,http.StatusOK, deleteAPIResponse)
	}

	return
}


func main() {

	log.Info("Started CRUD following are the APIs available....")


	//router for all APIs
	router := mux.NewRouter()


	router.HandleFunc("/users", GetAllUsers).Methods("GET") //Get all users present in the database
	router.HandleFunc("/users", InsertUser).Methods("POST") //inserts a user into database
	router.HandleFunc("/users/{username}", GetUserByName).Methods("GET")
	router.HandleFunc("/users/{username}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8888", router))
}
