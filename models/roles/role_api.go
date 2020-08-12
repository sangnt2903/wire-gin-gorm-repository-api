package roles

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleApi struct {
	RoleService RoleService
}

func ProvideRoleApi(service RoleService) RoleApi {
	return RoleApi{RoleService: service}
}

func (p *RoleApi) GetAll(ctx *gin.Context) {
	roles := p.RoleService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"roles" : roles,
	})
}

func (p *RoleApi) GetById(ctx *gin.Context){
	id,_ := strconv.Atoi(ctx.Param("id"))
	role := p.RoleService.GetById(uint(id))

	ctx.JSON(http.StatusOK, gin.H{
		"role" : role,
	})
}

func (p *RoleApi) PostRole(ctx *gin.Context){
	var roleInput Role
	if err := ctx.ShouldBindJSON(&roleInput); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return
	}

	role := p.RoleService.Save(roleInput)
	ctx.JSON(http.StatusCreated, gin.H{
		"role" : role,
	})

	return
}