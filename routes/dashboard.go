package routes

import (
	"mime"
	"net/http"

	paths "path"

	"github.com/ashhhleyyy/clickd/dashboard"
	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	path := c.Params.ByName("path")
	data, err := dashboard.DashboardAsset(path)
	if err != nil {
		panic(err)
	}
	ext := paths.Ext(path)
	var mimetype string
	if ext == "" {
		mimetype = "text/plain"
	} else {
		mimetype = mime.TypeByExtension(ext)
	}
	c.Data(http.StatusOK, mimetype, data)
}
