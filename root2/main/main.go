package main

import (
	"fmt"
	"root2/domain"
	"root2/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (cc CustomerController) Update(id string, c domain.Customer) {
	err := cc.store.Update(id, c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer has been updated")
}

func (cc CustomerController) GetById(id string) {
	c, err := cc.store.GetById(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(c)
}
func (cc CustomerController) GetAll() {
	c, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, val := range c {
		fmt.Println(val)
	}
}
func (cc CustomerController) Delete(id string) {
	err := cc.store.Delete(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("")
}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}

	customer := domain.Customer{
		ID:    "C101",
		Name:  "AMAN GUPTA",
		Email: "abc123@gmail.com",
	}

	customer1 := domain.Customer{
		ID:    "C102",
		Name:  "AMAN KUMAR",
		Email: "DEF123@gmail.com",
	}
	controller.Add(customer)
	controller.Add(customer1)
	customer.Email = "1234@gmial.com"
	controller.Update("C101", customer)
	controller.GetById("C101")
	controller.GetAll()
	controller.Delete("C102")
	fmt.Println("---------------")
	controller.GetAll()

}
