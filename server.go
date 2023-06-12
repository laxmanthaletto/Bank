package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ACC struct
{
	AC int
	Name string
	Balance int
	Pin int
}

type Request struct 
{
	Ac  int
	Pin int
	Amt int
}

type API int

var database []ACC

type CreateAccountRequest struct {
	Name string
	Pin  int
}

func (a *API) CreateAccount(req *CreateAccountRequest, reply *int) error {
	ac := len(database) + 9999
	account := ACC{
		AC:      ac,
		Name:    req.Name,
		Balance: 0,
		Pin:     req.Pin,
	}
	database = append(database, account)
	*reply = ac
	return nil
}

func (a *API) GetAccount(req *Request, reply *ACC) error {
	var temp ACC
	for idx, val := range database {
		if val.AC == req.Ac {
			if val.Pin == req.Pin {
				temp = database[idx]
			}
		}
	}
	*reply = temp
	return nil
}

func (a *API) Withdraw(req *Request, reply *ACC) error {
	var temp ACC
	for idx, val := range database {
		if val.AC == req.Ac {
			if val.Pin == req.Pin {
				if val.Balance > req.Amt {
					database[idx].Balance = val.Balance - req.Amt
					temp = database[idx]
				}
			}
		}
	}
	*reply = temp
	return nil
}

func (a *API) Deposit(req *Request, reply *ACC) error {
	var temp ACC
	for idx, val := range database {
		if val.AC == req.Ac {
			if val.Pin == req.Pin {
				if req.Amt <= 20000 {
					database[idx].Balance = val.Balance + req.Amt
					temp = database[idx]
				}
			}
		}
	}
	*reply = temp
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Error binding to port 8080", err)
	}

	log.Printf("Bank Server started on port %d", 8080)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error Serving: ", err)
	}
}