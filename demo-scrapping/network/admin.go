package network

import (
	"demo-scrapping/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type admin struct {
	network *Network
}

func newAdmin(network *Network) {
	a := &admin{network: network}

	basePath := "/admin"

	network.register(basePath+"/add", POST, a.add)
	network.register(basePath+"/update", PUT, a.update)

	network.register(basePath+"/delete", DELETE, a.delete)
	network.register(basePath+"/view", GET, a.view)
	network.register(basePath+"/view-all", GET, a.viewAll)
}

func (a *admin) add(c *gin.Context) {
	var req types.AddReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else {
		fmt.Println(req.CardSelector)
		res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
	}

}

func (a *admin) update(c *gin.Context) {
	var req types.UpdateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else {
		res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
	}
}

func (a *admin) view(c *gin.Context) {
	var req types.ViewReq

	if err := c.ShouldBindQuery(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else {
		res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
	}
}

func (a *admin) viewAll(c *gin.Context) {
	res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
}

func (a *admin) delete(c *gin.Context) {
	var req types.DeleteReq

	if err := c.ShouldBindQuery(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else {
		res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
	}
}
