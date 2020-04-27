package datatypes

import "fmt"

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK  = true
)

func DispConst() {
	fmt.Println("The Constant Values Are")
	fmt.Println("Product : ",PRODUCT)
	fmt.Println("Quantity :",QUANTITY)
	fmt.Println("Price :",PRICE)
	fmt.Println("Stock :",STOCK)
}