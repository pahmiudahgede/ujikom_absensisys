package seeders

import (
	"absensibe/models"
	"fmt"

	"gorm.io/gorm"
)

type StudentAddressSeeder struct{}

func (s *StudentAddressSeeder) GetName() string {
	return "Student Addresses"
}

func (s *StudentAddressSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.StudentAddress{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var students []models.Student
	if err := db.Limit(5).Find(&students).Error; err != nil {
		return err
	}

	addresses := []models.StudentAddress{}
	for i, student := range students {
		address := models.StudentAddress{
			StudentID: student.ID,
			AddressBase: models.AddressBase{
				Province:      "DKI Jakarta",
				District:      "Jakarta Pusat",
				Subdistrict:   "Menteng",
				Hamlet:        "Menteng",
				Village:       "Menteng",
				Neighbourhood: "Menteng",
				PostalCode:    "10310",
				Detail:        stringPtr(fmt.Sprintf("Jl. Menteng Raya No. %d", i+1)),
				Religion:      "islam",
			},
		}
		addresses = append(addresses, address)
	}

	return db.Create(&addresses).Error
}
