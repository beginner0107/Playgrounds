package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// POST, GET, PUT, DELETE

type API_REQUEST uint8

const (
	GET API_REQUEST = iota
	POST
	PUT
	DELETE
)

type header struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

type response struct {
	*header
	Result interface{} `json:"result"`
}

// Router에서 사용가능한 범용성있는 유틸 함수

func res(c *gin.Context, code int, res interface{}, data ...string) {
	c.JSON(code, &response{
		header: &header{Status: code, Data: strings.Join(data, " ,")},
		Result: res,
	})
}

func (n *Network) register(path string, t API_REQUEST, h ...gin.HandlerFunc) gin.IRoutes {
	switch t {
	case GET:
		return n.engin.GET(path, h...)
	case POST:
		return n.engin.POST(path, h...)
	case PUT:
		return n.engin.PUT(path, h...)
	case DELETE:
		return n.engin.DELETE(path, h...)

	default:
		return nil
	}
}

func (n *Network) verifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := getSecretByAuthorization(c)

		if secret == "" {
			res(c, http.StatusUnauthorized, nil, "auth secret need")
		} else {
			if valid, err := n.authenticator.VerifySecret(secret); err != nil {
				res(c, http.StatusUnauthorized, nil, err.Error())
				c.Abort()
			} else if !valid {
				res(c, http.StatusUnauthorized, nil, "not valid")
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}

func getSecretByAuthorization(c *gin.Context) string {
	auth := c.Request.Header.Get("Authorization")
	authSlied := strings.Split(auth, " ")
	if len(authSlied) > 1 {
		return authSlied[1]
	} else {
		return ""
	}
}
