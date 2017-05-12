package main

import (
	"html/template"
	"os"
)

const tax = 6.75 / 100

type Product struct {
	Name string
	Price float32
}


func (p Product) PriceWithTax() float32 {
	return p.Price * (1 + tax)
}

// Minus indicates do not print the new lines
const templateString = `
 {{- "Item Informatioan"}}
 Name: {{ .Name }}
 Price: {{ .Price | printf "$%.2f" }}
 PriceWithTax: {{ printf "$%.2f" .PriceWithTax }} 
 `

func main() {
	p := Product {
		Name: "Lemonade",
		Price: 2.16,
	}

	t := template.Must(template.New("").Parse(templateString))
	t.Execute(os.Stdout, p)
}