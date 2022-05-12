package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Initialize Struct
type (
	AreaRepository struct {
		DB *gorm.DB
	}
	Area struct {
		ID        int64   `gorm:"column:id;primaryKey;"`
		AreaValue float64 `gorm:"column:area_value"`
		AreaType  string  `gorm:"column:area_type"`
	}
)

func main() {
	// Connect to database with GORM
	gorm, err := Connect()
	if err != nil {
		panic(err)
	}
	// migrate the schema
	gorm.AutoMigrate(&Area{})
	ar := AreaRepository{
		DB: gorm,
	}
	service := service{
		repository: ar,
	}
	service.Service()

	fmt.Println("success")
}

// Setup the connection with the database
func Connect() (*gorm.DB, error) {
	serverHost := "localhost"
	serverPort := ":3306"
	serverUser := "root"
	serverPassword := "bangadam"
	serverDatabase := "backend_test_majoo_areas"
	url := fmt.Sprintf("%s:%s@(%s%s)/%s?charset=utf8&parseTime=True&loc=Local", serverUser, serverPassword, serverHost, serverPort, serverDatabase)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Create Insert Area Repository
func (_r *AreaRepository) InsertArea(param1 int64, param2 int64, typeArea string) (err error) {
	area := new(Area)
	switch typeArea {
		case "segitiga":
			formula := float64(0.5) * float64((param1 * param2))
			area.AreaValue = formula
			area.AreaType = typeArea
		case "persegi panjang", "persegi":
			formula := param1 * param2
			area.AreaValue = float64(formula)
			area.AreaType = typeArea
		default:
			area.AreaValue = 0
			area.AreaType = "undefined data"
	}

	err = _r.DB.Create(&area).Error
	if err != nil {
		return err
	}
	return nil
}

// Set service struct
type service struct {
	repository AreaRepository
}


func (_u service) Service() error {
	err := _u.repository.InsertArea(1, 2, "persegi panjang")
	if err != nil {
		return err
	}

	err = _u.repository.InsertArea(1, 2, "segitiga")
	if err != nil {
		return err
	}

	err = _u.repository.InsertArea(1, 2, "persegi")
	if err != nil {
		return err
	}
	return nil
}