# first commit

# Project Structure

entry file: main.go

## Directory and packages
main package:
  routes.go - contains all the routes
  util.go - for all the utility functions (like logging)

handlers/ - contains all the hadnlers
  ../users.go

model/ - contains all the models struct
  ../user.go

services/
  ../user.go - contains all the business logic 


localhost:3002/api/users/
--
{
"username": "bapak",
"password": "anak"
}



localhost:3002/api/parent/create
--
{
"name" : "ibu 2",
"user_id": 1
}

localhost:3002/api/children/create
--
{
"name": "jerry",
"age" : 20,
"height": "178 cm",
"parent_id": 2
}

localhost:3002/api/login
--
{
"username" : "ibu",
"password" : "anak"
}

