package routes

import (
	"net/http"
	"net/url"
	"time"

	"github.com/ashhhleyyy/clickd/config"
	"github.com/ashhhleyyy/clickd/model"
	"github.com/gin-gonic/gin"
)

func TrackEvent(c *gin.Context) {
	// Check Global Privacy Control early and bail :)
	if config.Configuration.FollowGPC && c.Request.Header.Get("Sec-GPC") == "1" {
		c.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "GPC indicates tracking not wanted",
		})
		return
	}

	var body struct {
		Type    string `json:"type" binding:"required"`
		Path    string `json:"path" binding:"required"`
		Referer string `json:"referer"`
		Width   int32  `json:"width" binding:"required"`
		Height  int32  `json:"height" binding:"required"`
	}

	domain := c.Request.Host

	if c.Bind(&body) == nil {
		refererUrl, err := url.Parse(body.Referer)
		if err != nil {
			panic(err)
		}
		tx := model.Database.Create(&model.Event{
			Domain:    domain,
			Timestamp: time.Now(),
			Path:      body.Path,
			Referer:   refererUrl.Host,
			Width:     body.Width,
			Height:    body.Height,
		})
		tx.Commit()
		c.JSON(http.StatusAccepted, gin.H{
			"ok": true,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "BadRequest",
		})
	}
}
