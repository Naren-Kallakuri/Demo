package entities

type OrderItem struct {
	Qty int
	Item Items
	
}

func GetAmountForItem(oi *OrderItem) float64{
	return ((oi.Item.GetRate()) * float64(oi.Qty))	
}
 
