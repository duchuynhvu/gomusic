package dblayer

import (
	"models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	return products, db.Find(&products).Error
}

//GetPromos returns all of the promotions
func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

//GetCustomerByName returns the customer's information
func (db *DBORM) GetCustomerByName(firstname, lastname string) (customer models.Customer, err error){
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error)
}

//GetCustomerByID retrieve a customer by customer ID
func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

//GetProduct retrieve a product by product ID
func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

//AddUser - add user
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error){
	//we will cover the hashpassword function later
	hashpassword(&customer.Pass)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

//SignInUser - user sign in
func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	//verify the password, we will cover this function later
	if !checkPassword(pass) {
		return customer, error.New("Invalid password")
	}
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	//update the loggedin field
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	return customer, result.Find(&cutomer).Error
}

//SignOutUserById - sign out user by Id
func (db DBORM) SignOutUserById(id int) error {
	//create a customer Go struct with the provided id
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	//update the customer row to reflect the fact that the customer is not logged in
	return db.Table("Customer").Where(&customer).Update("loggedin", 0).Error
}

//GetCustomerOrdersByID - get customer orders by ID
func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error){
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id=?", id).Scan(&orders).Error
}