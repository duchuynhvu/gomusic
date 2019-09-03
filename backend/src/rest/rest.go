package rest

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//RunAPI - app entry point
func RunAPI(address string) error {
	h, err := NewHandler("mysql", "root:secret@/gomusic")
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

//RunAPIWithHandler run API with Handler
func RunAPIWithHandler(address string, h HandlerInterface) error {
	//Get gin's default engine
	r := gin.Default()
	//r.Use(MyCustomLogger())
	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", h.GetCustomerByID)
		userGroup.GET("/:id/orders", h.GetOrders)
		userGroup.POST("/:id/signout", h.SignOut)
	}
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
		usersGroup.GET("", h.GetCustomers)
	}
	r.Use(static.ServeRoot("/", "../public/build"))
	return r.Run(address)
	//return r.RunTLS(address, "cert.pem", "key.pem")
}

//MyCustomLogger - custom middleware
func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("************************************")
		c.Next()
		fmt.Println("************************************")
	}
}
