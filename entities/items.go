package entities

type Items struct {
	ItemId int
	Description   string
	Rate float64	
}
 

func (i *Items) GetRate() float64 {
	return i.Rate
}