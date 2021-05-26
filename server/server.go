package server

import (
	"log"
	"strconv"

	"clog/database"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func addCLog(c *gin.Context) {
	var cLog database.CLog
	if err := c.Bind(&cLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": database.AddCLog(cLog)})
}

func getCLogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": database.GetCLogs()})
}

func getCLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": database.GetCLog(id)})
}

func deleteCLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DeleteCLog(id)
	c.JSON(http.StatusOK, gin.H{"result": "삭제했습니다."})
}

func updateCLog(c *gin.Context) {
	var cLog database.CLog
	if err := c.Bind(&cLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": database.UpdateCLog(cLog)})
}

func addCLogDays(c *gin.Context) {
	var cLogDays []database.CLogDay
	if err := c.Bind(&cLogDays); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": database.AddCLogDays(cLogDays)})
}

func getCLogDays(c *gin.Context) {
	cLogId, err := strconv.Atoi(c.Param("cLogId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": database.GetCLogDays(cLogId)})
}

func updateCLogDay(c *gin.Context) {
	var cLogDay database.CLogDay
	if err := c.Bind(&cLogDay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": database.UpdateCLogDay(cLogDay)})
}

func Start(port string) {
	database.Open()

	log.Println("서버를 시작합니다.")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowOrigins:     []string{"http://localhost:3000"},
	}))

	router.Static("/public", "./public")
	router.POST("/api/v1/cLog", addCLog)
	router.GET("/api/v1/cLogs", getCLogs)
	router.GET("/api/v1/cLog/:id", getCLog)
	router.DELETE("/api/v1/cLog/:id", deleteCLog)
	router.PUT("/api/v1/cLog", updateCLog)

	router.POST("/api/v1/cLogDays", addCLogDays)
	router.GET("/api/v1/cLogDays/:cLogId", getCLogDays)
	router.PUT("/api/v1/cLogDay", updateCLogDay)

	router.Run(port)
}
