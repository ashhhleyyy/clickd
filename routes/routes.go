package routes

import (
	"fmt"
	"time"

	"github.com/ashhhleyyy/clickd/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginDriver "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func SetupRoutes() *gin.Engine {
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  2,
	}

	rateLimitStore := memory.NewStore()
	rateLimiter := limiter.New(rateLimitStore, rate)

	rateLimitMiddleware := ginDriver.NewMiddleware(rateLimiter)

	fmt.Println(config.Configuration)

	cors := cors.New(cors.Config{
		AllowOrigins: config.Configuration.AllowedHosts,
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type"},
	})

	r := gin.Default()
	r.GET("/tracker.js", GetScript)
	r.GET("/views-by-day", GetViewsByDay)
	r.GET("/page-rankings", GetPageRankings)
	r.GET("/referer-rankings", GetRefererRankings)
	r.GET("/dashboard/*path", GetDashboard)
	trackGroup := r.Group("/", rateLimitMiddleware, cors)
	trackGroup.POST("track", TrackEvent)
	trackGroup.OPTIONS("track")

	return r
}
