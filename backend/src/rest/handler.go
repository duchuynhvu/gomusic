package rest

import (
	"dblayer"
	"models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HandlerInterface interface
type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

//Handler define a DBLayer object
type Handler struct {
	db dblayer.DBLayer
}

//NewHandler will return a new pointer to the Handler object
func NewHandler() (*Handler, error) {
	//this creates a new pointer to the Handler object
	return new(Handler), nil
}

//GetProducts will retrieve a Handler pointer and
//return a list of all products to the client
func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return
	}
	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

//GetPromos will retrieve a Handler pointer and
//return a list of all promotions to the client
func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {
		return
	}
	promos, err := h.db.GetAllPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promos)
}

//SignIn post user sign in
func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	//extract JSON document from the HTTP request body and parse it to the customer argument
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.SignInUser(customer.Email, customer.Pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

//AddUser posts user added
func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

//SignOut posts user sign out
func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {
		return
	}
	//get id parameter
	p := c.Param("id")
	//convert string to an integer type
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.db.SignOutUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

//GetOrders returns a list of all orders of a user to the client
func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		return
	}
	//get id param
	p := c.Param("id")
	//convert string to int type
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//call the database layer method to get orders from id
	orders, err := h.db.GetCustomerOrderByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

//Charge posts purchase charge
func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		return
	}
}
