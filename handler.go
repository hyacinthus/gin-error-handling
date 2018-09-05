package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDemoHandler get handler demo
func GetDemoHandler(c *gin.Context) {
	// get path param id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// handle a predefined error and return
		Handle(c, ErrInvalidID)
		return
	}
	// fake getting id from auth middleware
	userID := 1
	if userID == 0 {
		Handle(c, ErrNoAuth)
		return
	}
	// now get data
	resp, err := GetUserDemo(userID, id)
	if err != nil {
		// handle golang error interface, don't forget to return
		Handle(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// PostDemoHandler post handler demo
func PostDemoHandler(c *gin.Context) {
	// fake getting id from auth middleware
	userID := 1
	if userID == 0 {
		Handle(c, ErrNoAuth)
		return
	}
	// use ShouldBind and handle the error ourselves
	var req = new(DemoData)
	err := c.ShouldBind(req)
	if err != nil {
		// handle the error and custom status code and error key
		Handle(c, NewError(http.StatusBadRequest, "BadRequest", err.Error()))
		return
	}
	// save data
	resp, err := CreateUserDemo(userID, req.ID, req.Data)
	if err != nil {
		Handle(c, err)
		return
	}
	// you'd better return the data finally saved to db
	c.JSON(http.StatusCreated, resp)
}
