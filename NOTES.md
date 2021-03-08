# Implementation Notes

##Disclaimer:
What's included:
- Basic CRUD API.
What's not included (and would have been included):
- Authentication,
- Authorization (RBAC)
- Unit tests.

##Details
Language selected: Go 1.6
Persistence:  MongoDB
Need to create a "db" database.

A docker-compose.yml file is provided so you can run against a dockerized MongoDB running in the container.

To fire it, in the project directory: 

$ docker-compose up -d

Then you can start the server, in development mode: 

$ go run .

A few seconds later, in the console you'll see the message: 
Starting Users API on port 8080...!!

This API provide the four basic CRUD operations: GET, POST, PUT and DELETE.

Use Postman or similar client to test the endpoints:

1) localhost:8080/users
Get all Users.
As of this implementation returns all Users.
An additional implementation can be provided to allow pagination on large dataset.

2) localhost:8080/users
Inserts a new User into the users collection.
The body payload, with the User's data, should be similar to this one: 

{
  "name": "Capitão America",
  "age": "38",
  "password": "cu3c4",
  "email": "capitao.america@gmail.com",
  "address": "rua da amparo, 2"
}

Password will be hashed and salted.

3) localhost:8080/users/6039696b8de5e083850e4781

Removes the user with the specified ID from the Collection.

It should return:

{
  "message": "Uaer removed"
}

4) localhost:8080/users/6039696b8de5e083850e4781

Updates the document in the Users collection that matches the ObjectId provided.
Follows an example of the payload with the information to be updated:

{
  "age": "58",
  "password": "x1x1x1",
  "address": "rua dois, s/n"
}

It should return:

{
  "message": "Uaer updated"
}
