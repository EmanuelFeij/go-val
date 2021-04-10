package main

import (
	"log"
	"regexp"

	"github.com/go-playground/validator"
)

type Car struct {
	Manufacturer string `json:"manufacturer" validate:"required"`
	Plate        string `json:"plate" validate:"required,plate"`
	Owner        string `json:"owner" validate:"required"`
	OwnerEmail   string `json:"ownerEmail" validate:"required,email"`
}

func NewCar(manufacturer string, plate string, owner string, ownerEmail string) (*Car, error) {
	c := &Car{
		Manufacturer: manufacturer,
		Plate:        plate,
		Owner:        owner,
		OwnerEmail:   ownerEmail,
	}
	err := c.Validate()
	if err != nil {

		return nil, err
	}
	return c, nil
}

func (c *Car) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("plate", ValidateMyPlate)
	return validate.Struct(c)
}

func ValidateMyPlate(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("[A-Z]{2}-[0-9]{2}-[A-Z]{2}")
	s := fl.Field().String()
	matches := re.FindStringSubmatch(s)
	return len(matches) == 1

}

func main() {
	car, err := NewCar("smart", "AA-35-PP", "Emanuel", "mm@mm.com")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(car)
	}

}
