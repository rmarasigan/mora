package models

import (
	"github.com/uadmin/uadmin"
)

type Vendor struct {
	uadmin.Model
	User    uadmin.User
	UserID  uint
	Name    string
	Address string
	Rating  float64
}
