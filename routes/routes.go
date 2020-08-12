package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-repository-gin-gorm-wire/wire"
)


func SetRoutesConfig(db *gorm.DB) *gin.Engine {

	// initialize API
	roleAPI := wire.InitRoleApi(db)
	//



	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/roles", roleAPI.GetAll)
		v1.GET("/roles/:id", roleAPI.GetById)
		v1.POST("/roles", roleAPI.PostRole)
	}

	return r
}
