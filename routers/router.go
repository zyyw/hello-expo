package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zyyw/hello-expo/routers/api/v1"
	"github.com/zyyw/hello-expo/routers/static"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	// static files
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", static.HomePage)
	r.GET("/contact", static.ContactUs)
	r.GET("/about", static.About)

	// api routes
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/hello", v1.Hello)
	}

	return r
}
