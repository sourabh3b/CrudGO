package user

import (
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2"
	log "github.com/sirupsen/logrus"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var (

	//correct test fields
	correctUserName = "sbhagat"
	correctDisplayName = "Sourabh Bhagat"
	correctDepartment= "Development"

	//user object with all valid fields
	testUser = User {
		UserName: correctUserName,
		DisplayName:correctDisplayName,
		Department:correctDepartment,
	}

	//user object with empty username
	emptyUserNameInput = User {
		UserName: "",
		DisplayName:correctDisplayName,
		Department:correctDepartment,
	}


)


//SetupData - sets up data in testDB
func SetupData() {

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Println("Mongo error", err.Error())
	}
	defer session.Close()

	//insert test testUser into testDB
	err = session.DB("testDB").C("User").Insert(testUser)
	if err != nil {
		fmt.Println("unable to insert test user to DB", err.Error())
	}



}

//ClearData - clearup data from testDB after running testcases
func ClearData() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Println("Mongo error", err.Error())
	}
	defer session.Close()

	//mongo query to get delete test user
	err = session.DB("testDB").C("User").Remove(bson.M{"username": correctUserName})
	if err != nil {
		fmt.Println("unable to delete test user from DB", err.Error())
	}

}

//TestInsertNewUser - test case to test InsertNewUser function
func TestInsertNewUser(t *testing.T) {
	/*
	Test cases
	1. Invalid input
		1.1 empty username
	2. Valid input
		2.1 valid username

	*/
	SetupData()
	Convey("Testing InsertNewUser function", t, func() {

		Convey("Testing for incorrect input", func() {

			Convey("Testing for empty username input", func() {
				_, err := InsertNewUser(emptyUserNameInput)
				Convey("Error returned by InsertNewUser should not be nil", func() {
					So(err, ShouldNotBeNil)
				})

			})

		})
		Convey("Testing for correct input", func() {
			Convey("Testing for valid input (inserting test user)", func() {
				_, err := InsertNewUser(testUser)
				Convey("Error returned by InsertNewUser should be nil", func() {
					So(err, ShouldBeNil)
				})

			})

		})
	})
	ClearData()
}

//TestGetAllUsers - test case to test GetAllUsers function
func TestGetAllUsers(t *testing.T) {
	/*
		Testing function to get all users present in the database
		Not checking invalid username (because it should not allowed to be present in the database
	*/
	SetupData()
	Convey("Testing GetAllUsers function", t, func(){
		Convey("Testing for valid input", func(){
			allUsers, err := GetAllUsers()
			Convey("Error returned by GetAllUsers should be nil", func(){
				So(err, ShouldBeNil)
			})

			Convey("Length of users list returned by GetAllUsers should be not be nil", func(){
				So(len(allUsers), ShouldNotBeNil)
			})
		})

	})
	ClearData()
}

//TestGetUserByName - test case to test GetUserByName function
func TestGetUserByName(t *testing.T) {
	/*
		Test cases :
		1. Valid input
		2. Invalid input
	*/

	SetupData()
	Convey("Testing GetUserByName function", t, func(){
		Convey("Testing for invalid input",func(){
			Convey("Testing for empty username",func(){
				_,err := GetUserByName("")
				Convey("Error returned by GetUserByName should be nil", func(){
					So(err, ShouldBeNil)
				})

			})
		})
		Convey("Testing for valid input",func(){
			_, err := GetUserByName(correctUserName)
			Convey("Error returned by GetUserByName should be nil", func(){
				So(err, ShouldBeNil)
			})

		})
	})
	ClearData()
}

//TestDeleteUser - function to test DeleteUser function
func TestDeleteUser(t *testing.T) {
	SetupData()
	/*
		Test cases
		1. Valid input:  Since this function deletes document from database, therefore, ClearData function is not called
	*/
	Convey("Testing DeleteUser function", t,func(){
		DeleteUser(correctUserName)
		Convey("Document should be deleted after calling DeleteUser function", func(){

			session, err := mgo.Dial("127.0.0.1")
			if err != nil {
				log.Println("Mongo error", err.Error())
			}
			defer session.Close()

			userObject := User{}

			err = session.DB("testDB").C("User").Find(bson.M{"username": correctUserName}).One(&userObject)
			if err != nil {
				fmt.Println("unable to insert test user to DB", err.Error())
			}

			//Since user object is deleted, mongo query should NOT return error, for empty data
			So(err, ShouldBeNil)
		})
	})
}