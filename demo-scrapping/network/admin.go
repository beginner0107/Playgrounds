package network

import (
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

	network.register(basePath+"/view", GET, a.view)
	network.register(basePath+"/view-all", GET, a.viewAll)
	network.register(basePath+"/delete", DELETE, a.delete)
}

func (a *admin) add(c *gin.Context) {
	res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
}

func (a *admin) update(c *gin.Context) {
	res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
}

func (a *admin) view(c *gin.Context) {
	res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
}

func (a *admin) viewAll(c *gin.Context) {
	res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
}

func (a *admin) delete(c *gin.Context) {
	res(c, http.StatusOK, "test입니다.", "afdkjklfajkdsf")
}
