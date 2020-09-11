package main


import (
	"fmt"
	companyservice "github.com/abhisheksarkar444/entities"
	
)

func main() {

	//Create Company
	var a companyservice.Company
	a.Init()
	//fmt.Println(a.Name)
	companyservice.DataGreet(a)
	ptr := &a
	ptr.SetName("Amazon") 
	//fmt.Println(a)
	//companyservice.DataGreet(companyservice.Company{Name: "Pratian"})
	//companyservice.DataGreet(*ptr)

	//Create instance of Item Struct
	var i1 = companyservice.Items{ItemId: 1, Description: "Mobile" , Rate: 3000}
	//iPtr := &i1
	//fmt.Println(*iPtr)

	//Create instance of OrderItem struct
	var oi1 = companyservice.OrderItem{Qty: 2, Item: i1}
	//o1Ptr := &oi1
	//fmt.Println(oi1)
	//fmt.Println(oi1.Item.Description)
	//rate := companyservice.GetAmountForItem(o1Ptr)
	//fmt.Println("Rate of ",oi1.Item.Description," : ", rate)

	//Create an instance of Order Struct for Customer
	var o1 = companyservice.Orders{OrderId: 1}
	o1.OrderItems = append(o1.OrderItems,oi1)
	//fmt.Println(o1)

	//Customer Places Order
	//Create instance of Customer Struct and add Order for that customer
	var c1 = companyservice.Customer{Name: "Rahul", CustId: 11}
	var c2 = companyservice.Customer{Name: "Abhishek", CustId: 17}
	c1.PlaceOrder(o1)
	//fmt.Println("Name: ", c1.Name, " CustId: ", c1.CustId)
	//fmt.Println("TOTAL AMOUNT: ",companyservice.GetTotalAmountForCustomer(&c1))

	//Add Customer to the Company
	ptr.Customers = append(ptr.Customers,c1)
	ptr.Customers = append(ptr.Customers,c2)
	//fmt.Println(*ptr)

	fmt.Println("======================================================================")
	fmt.Println("\t\t\t ORDER DETAILS")
	fmt.Println("Company: ",ptr.GetName() )

	for _,cust := range ptr.Customers {
		fmt.Println("\nCustomer ID: ", cust.CustId,"\nCustomer Name: ",cust.Name)
		
		for _,order := range cust.OrderList{
			fmt.Println("\nOrder List:")

			for _,odrItem := range order.OrderItems{
				odPtr := &odrItem
				fmt.Println("\t\tItem: " , odrItem.Item.Description, "\t Rate: ", odrItem.Item.Rate,
			      "\tQuantity: ",odrItem.Qty, "\tTotal Amount: ", companyservice.GetAmountForItem(odPtr))
			}
		}
		
		fmt.Println("-----------------------")
	}

	fmt.Println("=========================================================================================")
		fmt.Println("\t\t\t\t\t       OrderValue Total: Rs.",companyservice.GetTotalWorthOfOrderPlaced(ptr))
}