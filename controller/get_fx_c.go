package controller

import (
	"fmt"
	"net/http"

	"github.com/hunzo/go-gin-swagger/models"

	"github.com/gin-gonic/gin"
)

// GetAll godoc
// @Summary  search Username
// @Description Search username
// @name surapong.n@nida.ac.th
// @Accept  json
// @Produce  json
// @Param username path string false "test path param"
// @Router /users/{username} [get]
// @Success 200 {object} string 'ok'
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ret := c.Param("username")

		fmt.Println(ret)

		c.JSON(200, gin.H{
			"function": "GetAll",
			"username": ret,
		})
	}
}

// SomePost godoc
// @Summary  Post User/pass
// @Description post user/password
// @name surapong.n@nida.ac.th
// @Accept  json
// @Produce  json
// @Param b body models.BodyUser true "post body" desc
// @Router /somepost [post]
// @Success 200 {object} models.BodyUser
func SomePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.BodyUser{}

		if e := c.ShouldBindJSON(&b); e != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		r := gin.H{
			"info": "test add",
			"data": b,
		}

		c.JSON(http.StatusOK, r)

	}

}
