package entities

import "fmt"

type Company struct {

	Name string
	Customers []Customer
	
}

func (c *Company) Init(){
	c.Name = "Pratian"
}

func DataGreet(c Company){
	fmt.Println("Welcome to Company service " + c.Name)
	
}

func (c *Company) SetName(company_name string){
	c.Name = company_name
}

func (c *Company) GetName() string {
	return c.Name
}

func GetTotalWorthOfOrderPlaced(c *Company) float64{
	value := 0.0

	for _,cust := range c.Customers {
		value += GetTotalAmountForCustomer(&cust)
	}
	return value	
}
 

