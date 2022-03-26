package main

import (
	"fmt"
	"github.com/jack410/factory_patterns/desp/factory"
)

func main() {
	//user := factory.CreateUser(factory.Aminuser)(123, "Kaka").(*models.Admin)
	//fmt.Println(user)
	newBook := new(factory.TechFactory).CreateProduct(factory.ProductTechBook).GetInfo()
	fmt.Println(newBook)

	newBriefs := new(factory.DailyFactory).CreateProduct(factory.ProductDailyBriefs).GetInfo()
	fmt.Println(newBriefs)
}
