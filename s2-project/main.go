package main

import (
	"fmt"
)

const shopName = "John's Auto Shop"

// Declare struct
type Store struct {
	Product Car
}

type Car struct {
	vehicle  string
	name     string
	model    string
	quantity int
	price    float64
	currency string
	id       string
}

type BuyerData struct {
	firstName string
	lastName  string
	address   string
	carId     int
	quantity  int
}

var orderList = make([]BuyerData, 0)

// Initilize product struct
var product1 = Store{Product: Car{
	vehicle:  "Car",
	name:     "BMW",
	model:    "X6",
	quantity: 4,
	price:    32000,
	currency: "USD",
	id:       "1",
},
}
var product2 = Store{
	Product: Car{
		vehicle:  "Car",
		name:     "Honda",
		model:    "Camaro",
		quantity: 18,
		price:    6000,
		currency: "USD",
		id:       "2",
	},
}

var product3 = Store{
	Product: Car{
		vehicle:  "Car",
		name:     "Toyota",
		model:    "2020 Corolla Sports",
		quantity: 12,
		price:    21000,
		currency: "USD",
		id:       "3",
	},
}

var storeItems = map[int]Store{
	1: product1,
	2: product2,
	3: product3,
}

func main() {
	welcomePage()
	car()
	productStartLevel()
	for {
		firstName, lastName, address, carId, quantity := getUserInputFromBuyer()
		isValidName, isValidAddress, isCarIdValid, isQuantityRequestValid := validateInputFromBuyer(firstName, lastName, address, int(carId), int(quantity))
		if isValidName && isValidAddress && isCarIdValid && isQuantityRequestValid {
			orderCreation(firstName, lastName, address, int(carId), int(quantity))
			productCurrentLevel(storeItems, carId, quantity)
			// generateReceipt(firstName, lastName, address, int(carId), int(quantity))
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name")
			}
			if !isValidAddress {
				fmt.Println("Please enter a valid address")
			}
			if !isCarIdValid {
				fmt.Println("Please enter a valid ID of the Car you want to buy")
			}
			if !isQuantityRequestValid {
				fmt.Println("Please enter a valid number for quantity")
			}
		}
	}

}

func welcomePage() {
	fmt.Printf("Welcome to %v store", shopName)
	fmt.Println("")

}

func car() {
	var cars []string
	for _, value := range storeItems {
		cars = append(cars, value.Product.name)
	}
	fmt.Println("Here are the cars available in this store: ", cars)
}

/* func (s Store) car() Store {
	for _, value := range storeItems {
		fmt.Printf("Here are the products available in this store - %v \n", value.Product)
	}
	return s
} */

func productStartLevel() {
	for _, value := range storeItems {
		fmt.Printf("Here are the products available in this store - %v \n", value.Product)
	}

}

func getUserInputFromBuyer() (string, string, string, int, int) {
	var firstName string
	var lastName string
	var address string
	var carId int
	var quantity int

	fmt.Println("Please enter your Firstname")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your Lastname")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your Address")
	fmt.Scan(&address)
	fmt.Println("Please enter the ID of the Car you want to buy")
	fmt.Scan(&carId)
	fmt.Println("Please enter the quantity required")
	fmt.Scan(&quantity)

	return firstName, lastName, address, carId, quantity
}

func validateInputFromBuyer(firstName string, lastName string, address string, carId int, quantity int) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidAddress := len(address) >= 6 && len(address) <= 40
	isCarIdValid := carId > 0 && carId <= 4

	selectedCar := storeItems[carId]
	remaining := selectedCar.Product.quantity
	isQuantityRequestValid := remaining >= quantity

	return isValidName, isValidAddress, isCarIdValid, isQuantityRequestValid
}

func inventory(storeItems map[int]Store, carId int, quantity int) map[int]Store {
	innerValueOfMap, ok := storeItems[carId]
	if !ok {
		fmt.Println("PLEASE INPUT A CORRECT CAR TYPE ID")
		return nil
	}

	remaining := innerValueOfMap.Product.quantity
	remaining = remaining - quantity
	innerValueOfMap.Product.quantity = remaining
	return storeItems

}

func orderCreation(firstName string, lastName string, address string, carId int, quantity int) {
	inv := inventory(storeItems, carId, quantity)
	var userInfo = BuyerData{
		firstName: firstName,
		lastName:  lastName,
		address:   address,
		carId:     carId,
		quantity:  quantity,
	}

	orderList = append(orderList, userInfo)

	fmt.Println("This is the list of cars to be sold", orderList)
	fmt.Println("Here is the list of the shop's remaining items: \n", inv)
}

func productCurrentLevel(storeItems map[int]Store, carId int, quantity int) map[int]Store {
	innerValues := storeItems[carId]

	remaining := innerValues.Product.quantity
	remaining = remaining - quantity
	if remaining < 0 {
		fmt.Println("Product Requested for is out of stock")
	}
	return storeItems
}
