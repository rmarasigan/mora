package main

import (
	"github.com/rmarasigan/golang-batangas/mora/models"
	"github.com/rmarasigan/golang-batangas/mora/views"
	"github.com/uadmin/uadmin"
	"net/http"
)

func main() {
	uadmin.Register(
		models.Vendor{},
		models.Product{},
	)

	http.HandleFunc("/p/", views.Product)

	uadmin.StartServer()
}
