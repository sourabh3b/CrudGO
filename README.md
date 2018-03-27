# CrudGO
Basic CRUD API in Golang
## Installation

```
1. cd CrudGo 
2. go run main.go
3. Open Postman & test APIs
4. run testcases : go test -v OR goconvey

```
```
Note :
Markdown and HTML are turned off in code blocks:
<i>This is not italic</i>, and [this is not a link](http://example.com)

1. [Install Go ](https://golang.org/doc/install)
2. [Install Mongo](https://docs.mongodb.com/manual/installation/)
3. Install GoConvey (go get  "github.com/smartystreets/goconvey/convey")
4. Install Logrus (go get "github.com/sirupsen/logrus")

```
## Schema 
#### User
```
{ "users": 
    [{ "username": "jsmith", "displayName": "John Smith", "department": "Sales" }, 
    { "username": "jdoe", "displayName": "John Doe", "department": "Development" }] 
} 
```


### API Lists

#### API 1 : Get All Users 


* **URL**

  http://localhost:8888/users
  
* **Method:**
 
  `GET` 
*  **URL Params**

   None
   
* **Data Params**

  None
  
 * **Success Response:**
   
   Returns all users present in the database
 
   * **Code:** 200 <br />
     **Content:** 

            [
     
                       {
                           "username": "jsmith",
                           "displayName": "John Smith",
                           "department": "Sales"
                       },
                       {
                           "username": "sbhagat",
                           "displayName": "Sourabh Bhagat",
                           "department": "Development"
                       },
                       {
                           "username": "jdoe",
                           "displayName": "John Doe",
                           "department": "Sales"
                       }
             ]

* **Error Response:**

  This case occurs when mongo db is down. 

  * **Code:** 400 StatusBadRequest <br />
    **Content:** `{ user : [] }`
    
* **Notes:**
    
      


#### API 2 : Insert a user
* **URL**

  http://localhost:8888/users
  
* **Method:**
 
  `POST` 
*  **URL Params**

   None
   
* **Data Params**

     
     displayName=[string]
     
     department=[string]
     
   **Required:**
   
    username=[string] 
  
 * **Success Response:**
   
    
   * **Code:** 200 <br />
     **Content:** 

           {
               "status": 200,
               "message": "Inserted user into database"
           }

* **Error Response:**

  This case occurs when mongo db is down. 

  * **Code:** 400 StatusBadRequest <br />
    **Content:**
    
            Empty username
            {
                    "status": 400,
                    "message": "Invalid input : Please enter username"
            }

        OR User already found in the database
            {
                    "status": 409,
                    "message": "User already found in the database"
            }

         OR MongoDB is down
            {
                    "status": 403,
                    "message": Unable to insert user into database
            }
* **Notes:**
    make sure that database is created with name myDB
#### API 3 : Get user by user name
* **URL**

  http://localhost:8888/users/{username}
  
* **Method:**
 
  `GET` 
*  **URL Params**

   username
   
* **Data Params**
      
      NONE
     
   **Required:**
   
    username=[string] 
  
 * **Success Response:**
   
    
    URL : http://localhost:8888/users/jdoe
    
   * **Code:** 200 <br />
     **Content:** 

           {
               "displayName": "John Doe",
               "department": "Sales"
           }

* **Error Response:**

  This case occurs when mongo db is down. 

  * **Code:** 404 StatusNotFound <br />
    **Content:**
    
            If username doesnot exist in the database OR Mongo is down
           {
               "status": 404,
               "message": "User not found"
           }

* **Notes:**
    make sure that database is created with name myDB
    
#### API 4 : Delete user by username 

* **URL**

  http://localhost:8888/users/{username}
  
* **Method:**
 
  `DELETE` 
*  **URL Params**

        username
   
* **Data Params**
      
      NONE
     
   **Required:**
   
    username=[string] 
  
 * **Success Response:**
   
    
    URL : http://localhost:8888/users/jdoe
    
   * **Code:** 200 <br />
     **Content:** 

          {
              "status": 200,
              "message": "Deleted user from database"
          }

* **Error Response:**

  This case occurs when mongo db is down. 

  * **Code:** 404 StatusNotFound <br />
    **Content:**
    
            If username doesnot exist in the database OR Mongo is down
          {
              "status": 404,
              "message": "Unable to delete user from database"
          }

* **Notes:**
    make sure that database is created with name myDB
    
    


#### Running Test cases

```
sourabh:user sourabh$ go test -v
=== RUN   TestInsertNewUser

  Testing InsertNewUser function 
    Testing for incorrect input 
      Testing for empty username input 
        Error returned by InsertNewUser should not be nil ✔
    Testing for correct input 
      Testing for valid input (inserting test user) 
        Error returned by InsertNewUser should be nil ✔


2 total assertions

--- PASS: TestInsertNewUser (0.01s)
=== RUN   TestGetAllUsers

  Testing GetAllUsers function 
    Testing for valid input 
      Error returned by GetAllUsers should be nil ✔
      Length of users list returned by GetAllUsers should be not be nil ✔


4 total assertions

--- PASS: TestGetAllUsers (0.00s)
=== RUN   TestGetUserByName

  Testing GetUserByName function 
    Testing for invalid input 
      Testing for empty username 
        Error returned by GetUserByName should be nil ✔
    Testing for valid input 
      Error returned by GetUserByName should be nil ✔


6 total assertions

--- PASS: TestGetUserByName (0.00s)
=== RUN   TestDeleteUser

  Testing DeleteUser function 
    Document should be deleted after calling DeleteUser function ✔


7 total assertions

--- PASS: TestDeleteUser (0.02s)
PASS
ok      github.com/CrudGO/user  0.053s

```
