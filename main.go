package main

import (
	"net/http"

	docs "dream/backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Hell Blue
// @Router /hello [get]
func Hello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name": "Blue",
	})
}

// HelloEcho godoc
// @Summary Hello with name
// @Description Return greeting with name
// @Tags example
// @Accept json
// @Produce json
// @Param name path string true "Name to greet"
// @Success 200 {string} string "Hello {name}"
// @Router /hello/{name} [get]
func HelloEcho(c *gin.Context) {
	name := c.Param("name")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name": name,
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("view/index.html")
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			eg.GET("/hello", Hello)
			eg.GET("/hello/:name", HelloEcho)
		}

	}
	// Swagger API documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/hello", Hello)
	router.GET("/hello/:name", HelloEcho)

	router.Run(":8000")
}
