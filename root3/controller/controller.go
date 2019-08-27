package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"root3/domain"
	"root3/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

var CC = CustomerController{
	store: mapstore.NewMapStore(),
}

func Create(w http.ResponseWriter, r *http.Request) {
	c := domain.Customer{
		ID:    "C101",
		Name:  "AMAN GUPTA",
		Email: "abc123@gmail.com",
	}
	err := CC.store.Create(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Customer has been created")
	rs, _ := json.Marshal("new user created\n")
	rs1, _ := json.Marshal(c)
	w.Write(rs)
	w.Write(rs1)

}

func Update(w http.ResponseWriter, r *http.Request) {
	c := domain.Customer{
		ID:    "C101",
		Name:  "AMAN GUPTA",
		Email: "abc321@gmail.com",
	}
	err := CC.store.Update(c.Email, c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Customer has been updated")
	rs, _ := json.Marshal("new user created\n")
	rs1, _ := json.Marshal(c)
	w.Write(rs)
	w.Write(rs1)

}

func GetById(w http.ResponseWriter, r *http.Request) {
	c := domain.Customer{
		ID: "C101",
	}
	c, err := CC.store.GetById(c.ID)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(c)
	rs, _ := json.Marshal("new user created\n")
	rs1, _ := json.Marshal(c)
	w.Write(rs)
	w.Write(rs1)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	c, err := CC.store.GetAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, val := range c {

		fmt.Println(val)
	}
	rs, _ := json.Marshal("new user created\n")
	rs1, _ := json.Marshal(c)
	w.Write(rs)
	w.Write(rs1)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	c := domain.Customer{
		ID: "C101",
	}
	err := CC.store.Delete(c.ID)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(c)
	rs, _ := json.Marshal("new user created\n")
	rs1, _ := json.Marshal(c)
	w.Write(rs)
	w.Write(rs1)
}
