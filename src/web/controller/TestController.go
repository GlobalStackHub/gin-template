package controller

import (
	"app-test/src/web/entity"
	"app-test/src/web/entity/response"
	"app-test/src/web/repo"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (TestController) Insert(c *gin.Context) {
	email := entity.Email{ID: 10, UserID: 1111, Email: "931305033@qq.com", Subscribed: true}
	err := repo.InsertEmail(email)
	if err != nil {
		print(err)
	} else {
		response.ResSuccessMsg(c)
	}
}

func (TestController) Other(c *gin.Context) {
	println("other .")
}
