package user

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//User - data model for a user
type User struct {
	UserName    string `bson:"username" json:"username"`       // username of user example :jsmith
	DisplayName string `bson:"displayName" json:"displayName"` // displayName of user example :John Smith
	Department  string `bson:"department" json:"department"`   // department for the user example :Sales
}


//GetUserResponse - Get User Response
type GetUserResponse struct {
	DisplayName string `bson:"displayName" json:"displayName"` // displayName of user example :John Smith
	Department  string `bson:"department" json:"department"`   // department for the user example :Sales
}

//Response - data model for API response
type Response struct {
	Status  int     `bson:"status" json:"status"`
	Message string  `bson:"message" json:"message"`
}



//GetAllUsers -  return all users present in the database
func GetAllUsers() ([]User, error) {

	//initialising user array object to be returned when calling this function
	users := []User{}

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Println("Mongo error", err.Error())
		return users, errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()

	//mongo query to get all users in myDB database and User table
	err = session.DB("myDB").C("User").Find(nil).All(&users) //todo: move constant string to a constant file //todo: rename db name

	return users, err
}


//GetUserByName -  return all users present in the database
func GetUserByName(username string) (GetUserResponse, error) {

	//initialising user array object to be returned when calling this function
	user := User{}
	getUserResponse := GetUserResponse{}

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Println("Mongo error", err.Error())
		return getUserResponse, errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()

	//mongo query to get all users in myDB database and User table
	err = session.DB("myDB").C("User").Find(bson.M{"username": username}).One(&user)

	getUserResponse.Department = user.Department
	getUserResponse.DisplayName = user.DisplayName

	return getUserResponse, err
}


//InsertNewUser -  return all users present in the database
func InsertNewUser(user User) ( bool,error) {

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Println("Mongo error", err.Error())
		return  false, errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()

	//check if user already exist
	err = session.DB("myDB").C("User").Find(bson.M{"username": user.UserName}).One(&user)
	if err == nil {
		return  true,nil
	}

	//mongo query to insert new user
	err = session.DB("myDB").C("User").Insert(user)

	return  false, err
}


//DeleteUser -  return all users present in the database
func DeleteUser(username string) (error) {

	//initialising user array object to be returned when calling this function
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Println("Mongo error", err.Error())
		return errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()

	//mongo query to get all users in myDB database and User table
	err = session.DB("myDB").C("User").Remove(bson.M{"username": username})

	return err
}


