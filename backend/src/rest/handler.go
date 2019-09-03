package rest

import (
	"dblayer"
	"fmt"
	"log"
	"models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

//HandlerInterface interface
type HandlerInterface interface {
	GetCustomerByID(c *gin.Context)
	GetProducts(c *gin.Context)
	GetCustomers(c *gin.Context)
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

//NewHandler - constructor with params
func NewHandler(dbtype, constring string) (HandlerInterface, error) {
	db, err := dblayer.NewORM(dbtype, constring+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

//NewHandlerWithDB - contructor with db
func NewHandlerWithDB(db dblayer.DBLayer) HandlerInterface {
	return &Handler{db: db}
}

//GetCustomerByID - get customer by id
func (h *Handler) GetCustomerByID(c *gin.Context) {
	if h.db == nil {
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	customer, err := h.db.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	customer.Pass = ""
	c.JSON(http.StatusOK, customer)
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
	fmt.Printf("Found %d products\n", len(products))
	c.JSON(http.StatusOK, products)
}

//GetPromos returns a list of all promotions to the client
func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {
		return
	}
	promos, err := h.db.GetPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promos)
}

//GetCustomers return a list of customers
func (h *Handler) GetCustomers(c *gin.Context) {
	if h == nil {
		return
	}
	customers, err := h.db.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
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
		//if the error is invalid password, return forbidden http error
		if err == dblayer.ErrINVALIDPASSWORD {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
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
	err = h.db.SignOutUserByID(id)
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
	orders, err := h.db.GetCustomerOrdersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

//Charge posts purchase charge
func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	//declare the type that can accept the data receiving from the frontend
	request := struct {
		models.Order
		Remember    bool   `json:"rememberCard"`
		UseExisting bool   `json:"useExisting"`
		Token       string `json:"token"`
	}{}
	//parse the incoming JSON payload
	err := c.ShouldBindJSON(&request)
	log.Printf("request: %+v \n", request)
	if err != nil {
		c.JSON(http.StatusBadRequest, request)
		return
	}
	// Set your secret key: remember to change this to your live secret key in production
	// Keys can be obtained from: https://dashboard.stripe.com/account/apikeys
	// They key below is just for testing
	stripe.Key = "sk_test_NA4BTh8j2IJxqVn3lImdIdEZ00rLr93apK"
	//test cards available at:	https://stripe.com/docs/testing#cards
	//setting charge parameters
	chargeP := &stripe.ChargeParams{
		//the price we obtained from the incoming request
		Amount: stripe.Int64(int64(request.Price)),
		//the currency
		Currency: stripe.String("usd"),
		//the description
		Description: stripe.String("GoMusic charge..."),
	}
	//initialize the stripe customer ID string
	stripeCustomerID := ""
	//Either remembercard or use exeisting should be enabled but not both
	if request.UseExisting {
		log.Println("Getting credit card id...")
		//this is a new method which retrieve the stripe customer id from the database
		stripeCustomerID, err = h.db.GetCreditCardCID(request.CustomerID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		cp := &stripe.CustomerParams{}
		cp.SetSource(request.Token)
		customer, err := customer.New(cp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		stripeCustomerID = customer.ID

		log.Printf("Remember? %t", request.Remember)
		// if the request asks to remember the card
		if request.Remember {
			log.Printf("Getting CCID: %s", stripeCustomerID)
			//save the stripe customer id, and link it to the actual customer id in our database
			err = h.db.SaveCreditCardForCustomer(request.CustomerID, stripeCustomerID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	/*
		for simplicity, let's assume it's a new order
		otherwise, we should check if the customer already ordered the same item or not
	*/

	chargeP.Customer = stripe.String(stripeCustomerID)
	//charge the credit card
	_, err = charge.New(chargeP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//store the order to database
	err = h.db.AddOrder(request.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
