package static

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", map[string]string{
		"title":   "Home",
		"content": "Home Page",
	})
}

func ContactUs(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", map[string]string{
		"title":   "Contact us",
		"content": "Contact us Page",
	})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", map[string]string{
		"title":   "About",
		"content": "About Page",
	})
}
