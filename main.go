package main

import (
	"net/http"
	"os"

	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	// swagger
	router.StaticFS("/public/", http.Dir("public"))

	// monitoring
	router.GET("/readyz", routes.Status)
	router.GET("/livez", routes.Status)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// api
	api := router.Group("api")
	// api v1
	v1 := api.Group("v1")

	// customer
	customer := v1.Group("customer")
	customer.POST("/", routes.CreateCustomer)
	customer.GET("/", routes.GetCustomers)
	customer.GET("/:id/", routes.GetCustomerById)
	customer.PUT("/:id", routes.UpdateCustomer)
	customer.DELETE("/:id", routes.DeleteCustomer)

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}
