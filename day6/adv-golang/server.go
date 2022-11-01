package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Resp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Users []User

func saveUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	u := addUserToDatabase(name, email)

	return c.JSON(http.StatusCreated, Resp{
		Success: true,
		Message: "User created",
		Data:    u,
	})
}

func addUserToDatabase(name, email string) User {
	Users = append(Users, User{Name: name, Email: email})
	return User{Name: name, Email: email}
}

// func getUsers(c echo.Context) error {

// 	return c.JSON(http.StatusOK, Resp{
// 		Success: true,
// 		Data:    Users,
// 	})
// }

// path parameters
func getUser(c echo.Context) error {
	id := c.Param("id")

	u := getUserByName(id)

	return c.JSON(http.StatusOK, Resp{
		Success: true,
		Data:    u,
	})
}

func getUserByName(name string) User {
	var u User
	for _, user := range Users {
		if user.Name == name {
			u = user
			break
		}
	}
	return u
}

// query parameters
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// func updateUser(c echo.Context) error {
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	return c.JSON(http.StatusOK, Resp{
// 		Success: true,
// 		Data:    User{Name: name, Email: email},
// 	})
// }

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	u := deleteUserByName(id)

	return c.JSON(http.StatusOK, Resp{
		Success: true,
		Message: "User deleted",
		Data:    u,
	})
}

func deleteUserByName(name string) User {
	var u User
	var index int
	for i, user := range Users {
		if user.Name == name {
			index = i
			u = user
			break
		}
	}
	Users = removeUserIndex(Users, index)
	return u
}

func removeUserIndex(s []User, index int) []User {
	return append(s[:index], s[index+1:]...)
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/show", show)

	e.POST("/users", saveUser)

	// e.GET("/users", getUsers)

	e.GET("/users/:id", getUser)

	// e.PUT("/users/:id", updateUser)

	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323")) // e.Start should be at the end of code

}
