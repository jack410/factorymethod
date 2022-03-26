package factory

import "github.com/jack410/factory_patterns/desp/models"

const (
	ProductTechBook = iota
	ProductDailyBriefs
)

type ProductType int

//总工厂
type IProductFactory interface {
	CreateProduct(t ProductType) IProduct
}

//总商品
type IProduct interface {
	GetInfo() string
}

type TechFactory struct{}

func (*TechFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductTechBook:
		return &models.Book{}
	}
	return nil
}

type DailyFactory struct{}

func (*DailyFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case ProductDailyBriefs:
		return &models.Briefs{}
	}
	return nil
}
