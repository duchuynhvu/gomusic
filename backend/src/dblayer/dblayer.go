package dblayer

import (
	"errors"
	"models"
)

//ErrINVALIDPASSWORD - invalid password error
var ErrINVALIDPASSWORD = errors.New("Invalid password")

//DBLayer - DBlayer interface
type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetAllCustomers() ([]models.Customer, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(int) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
	AddOrder(models.Order) error
	GetCreditCardCID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
	DeleteUser(int) error
}
