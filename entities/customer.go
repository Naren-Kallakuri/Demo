package entities

type Customer struct {
	Name   string
	CustId int
	OrderList []Orders
}

func (c *Customer) PlaceOrder(o Orders){
	c.OrderList = append(c.OrderList, o)
}

func GetTotalAmountForCustomer(c *Customer) float64{
	totalAmount := 0.0

	for _,element := range c.OrderList {
		totalAmount += GetAmountForOrders(&element)
		//fmt.Printf(element.Item.Description)
	}
	return totalAmount	
}


 
