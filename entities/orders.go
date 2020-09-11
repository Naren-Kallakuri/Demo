package entities
//import "fmt"

type Orders struct {
	OrderId int
	OrderItems []OrderItem

}

func GetAmountForOrders(o *Orders) float64{
	total := 0.0

	for _,element := range o.OrderItems {
		total += GetAmountForItem(&element)
		//fmt.Printf(element.Item.Description)
	}
	return total
	
}
 
