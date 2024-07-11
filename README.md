# Booklyifywizzbangthingaroo 
Project demonstrating a basic rest api solution to better understand go. This project will leverage the gin framework, sqllite
and will provide a user auth mechanism to protect certain endpoints. This will be accomplished via JWT.

## Dependencies
* Go
* Gin
  * `go get -u github.com/gin-gonic/gin`
* SQLite
  * `go get github.com/glebarez/go-sqlite`

## Plan
Create an "events booking" API with the following:

| Action | Endpoint              | Description                    | Auth Required               |
|:-------|:----------------------|:-------------------------------|:----------------------------|
| GET    | /events               | Get a list of available events | No                          |
| GET    | /events/<id>          | Get a specific event by id     | No                          |
| POST   | /events               | Create a new bookable event    | Yes                         |
| PUT    | /events/<id>          | Update an event                | Yes - only by creator       |
| DELETE | /events/<id>          | Delete an event                | Yes - only by creator       |
| POST   | /signup               | Create a new user              | No                          |
| POST   | /login                | Authenticate a user            | No - auth will leverage JWT |
| POST   | /events/<id>/register | Register user for an event     | Yes                         |
| DELETE | /events/<id>/register | Cancel an event registration   | Yes                         |

