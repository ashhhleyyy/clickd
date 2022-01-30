package routes

import (
	"net/http"

	"github.com/ashhhleyyy/clickd/tracker"
	"github.com/gin-gonic/gin"
)

func GetScript(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", tracker.Script())
}
