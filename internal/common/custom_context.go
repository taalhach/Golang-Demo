package common

import (
	"github.com/labstack/echo/v4"
	"github.com/taalhach/Golang-Demo/internal/configs"
	"gorm.io/gorm"
)

type CustomContext struct {
	echo.Context
	DB          *gorm.DB
	MainConfigs *configs.MainConfig
	dbSession   *gorm.DB
}

func (cc *CustomContext) DbSession() *gorm.DB {
	if cc.dbSession == nil {
		cc.dbSession = NewSession(cc.DB)
	}

	return cc.dbSession
}

func NewSession(db *gorm.DB) *gorm.DB {
	session := db.Session(&gorm.Session{})
	return session.Begin()
}
