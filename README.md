# CrudGO
Basic CRUD API in Golang
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
    