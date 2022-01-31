package routes

import (
	"net/http"
	"time"

	"github.com/ashhhleyyy/clickd/model"
	"github.com/gin-gonic/gin"
)

type ViewByDay struct {
	Date  int64 `json:"date"`
	Count int   `json:"count"`
}

func GetViewsByDay(c *gin.Context) {
	var result []ViewByDay = []ViewByDay{}
	rows, err := model.Database.Raw("SELECT DATE(timestamp) as date, COUNT(*) as count FROM events GROUP BY date ORDER BY date").Rows()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var row struct {
			Date  time.Time
			Count int
		}
		model.Database.ScanRows(rows, &row)
		result = append(result, ViewByDay{
			Date:  row.Date.Unix(),
			Count: row.Count,
		})
	}

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{
		"values": result,
	})
}

type Ranking struct {
	Key   string `json:"key"`
	Rank  int    `json:"rank"`
	Value int    `json:"value"`
}

func GetPageRankings(c *gin.Context) {
	var result []Ranking = []Ranking{}
	rows, err := model.Database.Raw("SELECT domain, path, COUNT(*) as count FROM events GROUP BY domain, path ORDER BY count DESC LIMIT 10").Rows()
	if err != nil {
		panic(err)
	}

	ranking := 1
	for rows.Next() {
		var row struct {
			Domain string
			Path   string
			Count  int
		}
		model.Database.ScanRows(rows, &row)
		result = append(result, Ranking{
			Key:   row.Domain + row.Path,
			Rank:  ranking,
			Value: row.Count,
		})
		ranking += 1
	}

	c.JSON(http.StatusOK, gin.H{
		"rankings": result,
	})
}

func GetRefererRankings(c *gin.Context) {
	var result []Ranking = []Ranking{}
	rows, err := model.Database.Raw("SELECT referer, COUNT(*) as count FROM events WHERE referer != '' GROUP BY referer ORDER BY count DESC LIMIT 10").Rows()
	if err != nil {
		panic(err)
	}

	ranking := 1
	for rows.Next() {
		var row struct {
			Referer string
			Count   int
		}
		model.Database.ScanRows(rows, &row)
		result = append(result, Ranking{
			Key:   row.Referer,
			Rank:  ranking,
			Value: row.Count,
		})
		ranking += 1
	}

	c.JSON(http.StatusOK, gin.H{
		"rankings": result,
	})
}
