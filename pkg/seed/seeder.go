package seed

import "gorm.io/gorm"

type Seeder struct {
	Func SeederFunc
	Name string
}

var seeders []Seeder

type SeederFunc func(db *gorm.DB)

func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Func: fn,
		Name: name,
	})
}
