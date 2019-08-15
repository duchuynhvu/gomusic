package rest

import (
	"github.com/gin-gonic/gin"
)

//RunAPI - app entry point
func RunAPI(address string) error {
	//get gin's default engine
	r := gin.Default()
	//define a handler
	h, _ := NewHandler("sql", "root:/gomusic")
	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	r.GET("/promos", h.GetPromos)
	//post user sign in
	r.POST("/users/signin", h.SignIn)
	//add user
	r.POST("/users", h.AddUser)
	//post user sign out
	r.POST("/user/:id/signout", h.SignOut)
	//get user orders
	r.GET("/user/:id/orders", h.GetOrders)
	//post purchase charge
	r.POST("/users/charge", h.Charge)
	//run the server
	return r.Run(address)
}
