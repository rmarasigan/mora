package models

import (
	"github.com/uadmin/uadmin"
)

type Product struct {
	uadmin.Model
	Vendor      Vendor `uadmin:"required"`
	VendorID    uint
	Name        string `uadmin:"required"`
	Description string `uadmin:"html"`
	Price       float64
	Image       string `uadmin:"image"`
}
