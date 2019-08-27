package rest

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//RunAPI - app entry point
func RunAPI(address string) error {
	h, err := NewHandler("mysql", "root:secret@tcp(localhost:3306)/gomusic")
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
	/*
		//post user sign in
		r.POST("/user/signin", h.SignIn)
		//post user sign out
		r.POST("/user/:id/signout", h.SignOut)
		//get user orders
		r.GET("/user/:id/orders", h.GetOrders)
		//post purchase charge
		r.POST("/user/charge", h.Charge)
	*/
	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
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
