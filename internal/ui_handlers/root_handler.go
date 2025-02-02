package ui_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taalhach/Golang-Demo/internal/common"
	"github.com/taalhach/Golang-Demo/internal/models"
	"github.com/taalhach/Golang-Demo/pkg/forms"
)

type RootHandlerResponse struct {
	forms.BasicResponse
	Expenses []*models.Expense
}

func RootHandler(c echo.Context) error {
	cc := c.(*common.CustomContext)
	// create session
	session := cc.DbSession()

	//fetch expenses list from db
	var expenses []*models.Expense
	session.Find(&expenses)

	ret := RootHandlerResponse{}
	ret.Success = true
	ret.Expenses = expenses

	return c.JSON(http.StatusOK, ret)
}
