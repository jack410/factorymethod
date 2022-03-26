package models

//技术类Tech
//图书
type Book struct {
	Id       int
	BookName string
}

func (this *Book) GetInfo() string {
	return "book"
}

//日用类
//内裤
type Briefs struct {
	Id   int
	Size string
}

func (this *Briefs) GetInfo() string {
	return "内裤"
}
