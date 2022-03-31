package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zyyw/hello-expo/pkg"
	e "github.com/zyyw/hello-expo/pkg/error"
)

const HelloExpo = "Hello Expo"

func Hello(c *gin.Context) {
	g := pkg.Gin{C: c}

	g.Response(http.StatusOK, e.SUCCESS, HelloExpo)
}
