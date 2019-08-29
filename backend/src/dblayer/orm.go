package dblayer

import (
	"errors"
	"models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

//DBORM struct
type DBORM struct {
	*gorm.DB
}

//NewORM - Constructor
func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

//GetAllProducts returns all of the products
func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Table("products").Scan(&products).Error
}

//GetPromos returns all of the promotions
func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Table("products").Where("promotion <> ?", 0).Scan(&products).Error
}

//GetCustomerByName returns the customer's information
func (db *DBORM) GetCustomerByName(firstname, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Scan(&customer).Error
}

//GetCustomerByID retrieve a customer by customer ID
func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

//GetProduct retrieve a product by product ID
func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.Table("products").First(&product, id).Error
}

//AddUser - add user
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	//pass received password by reference so that we can change it to its hashed version
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	customer.Pass = ""
	return customer, err
}

//SignInUser - user sign in
func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	//obtain a *gorm.DB object representing our customer's row
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	//retrieve the data for the customer with the passed email
	err = result.Scan(&customer).Error
	if err != nil {
		return customer, err
	}
	//compare the saved hashed password with the password provided by the user trying to sign in
	if !checkPassword(customer.Pass, pass) {
		//if failed, returns an error
		return customer, ErrINVALIDPASSWORD
	}
	//set cusomter pass to empty because we don't need to share this information again
	customer.Pass = ""
	//update the loggedin field
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	//return the new customer row
	return customer, result.Scan(&customer).Error
}

//SignOutUserByID - sign out user by Id
func (db DBORM) SignOutUserByID(id int) error {
	//create a customer Go struct with the provided id
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	//update the customer row to reflect the fact that the customer is not logged in
	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

//GetCustomerOrdersByID - get customer orders by ID
func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").
		Joins("join customers on customers.id = customer_id").
		Joins("join products on products.id = product_id").
		Where("customer_id=?", id).Scan(&orders).Error
}

//AddOrder - add the order to the orders table
func (db *DBORM) AddOrder(order models.Order) error {
	return db.Create(&order).Error
}

//GetCreditCardCID - Get the id representing the credit card from the database
func (db *DBORM) GetCreditCardCID(id int) (string, error) {
	customerWithCCID := struct {
		models.Customer
		CCID string `gorm:"column:cc_customerid"`
	}{}
	return customerWithCCID.CCID, db.Table("customers").Where(&customerWithCCID, id).Scan(&customerWithCCID).Error
}

//SaveCreditCardForCustomer - save the credit card information for the customer
func (db *DBORM) SaveCreditCardForCustomer(id int, ccid string) error {
	result := db.Table("customers").Where("id=?", id)
	return result.Update("cc_customerid", ccid).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	//convert password string to byte slice so that we can use it with the bcrypt package
	sBytes := []byte(*s)
	//Obtain hashed password via the GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	//return an error if the hash does not match the provided pass string
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}
