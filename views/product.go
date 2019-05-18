package views

import (
	"github.com/rmarasigan/golang-batangas/mora/models"
	"github.com/uadmin/uadmin"
	"html/template"
	"net/http"
	"strings"
)

func Product(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/product.html"))
	pID := strings.TrimPrefix(r.URL.Path, "/p/")

	p := models.Product{}
	uadmin.Get(&p, "id = ?", pID)
	uadmin.Preload(&p)

	context := map[string]interface{}{}

	context["Image"] = p.Image
	context["Name"] = p.Name
	context["Description"] = template.HTML(p.Description)
	context["Vendor"] = p.Vendor

	tmpl.Execute(w, context)
}
