package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("HELLO!")
	e := echo.New()
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	e.Use(serverHeader)
	//QueryParam along with PathVariable Get API
	e.GET("/cats", GetCats)
	e.GET("/cats2/:data", GetCats2)
	//Post Request
	e.POST("/cats", AddCat)
	e.Logger.Fatal(e.Start(":8000"))
}

//GET API which return the name of the cats specified in QueryParam
//http://localhost:8000/cats?name=arnold&type=fluffy

func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	return c.String(http.StatusOK, fmt.Sprintf("your cat name is : %s\nand cat type is : %s\n", catName, catType))
}

//http://localhost:8000/cats/json?name=arnold&type=fluffy
//data path variable accepts value as json/string

func GetCats2(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is : %s\nand cat type is : %s\n", catName, catType))
	} else if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as Sting or JSON"})
	}
}

func AddCat(c echo.Context) error {
	type Cat struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	cat := Cat{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&cat)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is yout cat %#v", cat)
	return c.String(http.StatusOK, "We got your Cat!!!")
}

//Custom Middleware
// ServerHeader middleware adds a custom header to the response.
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Custom-Header", "blah!!!")
		return next(c)
	}
}