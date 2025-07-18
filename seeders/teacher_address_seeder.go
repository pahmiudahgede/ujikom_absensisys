package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type TeacherAddressSeeder struct{}

func (s *TeacherAddressSeeder) GetName() string {
	return "Teacher Addresses"
}

func (s *TeacherAddressSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.TeacherAddress{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var teachers []models.Teacher
	if err := db.Limit(3).Find(&teachers).Error; err != nil {
		return err
	}

	if len(teachers) == 0 {
		return nil
	}

	addresses := []models.TeacherAddress{
		{
			TeacherID: teachers[0].ID,
			AddressBase: models.AddressBase{
				Province:      "DKI Jakarta",
				District:      "Jakarta Timur",
				Subdistrict:   "Cakung",
				Hamlet:        "Cakung Timur",
				Village:       "Cakung Timur",
				Neighbourhood: "Cakung Timur",
				PostalCode:    "13910",
				Detail:        stringPtr("Jl. Raya Cakung No. 15"),
				Religion:      "islam",
			},
		},
	}

	if len(teachers) > 1 {
		addresses = append(addresses, models.TeacherAddress{
			TeacherID: teachers[1].ID,
			AddressBase: models.AddressBase{
				Province:      "Jawa Barat",
				District:      "Bandung",
				Subdistrict:   "Cicendo",
				Hamlet:        "Cicendo",
				Village:       "Cicendo",
				Neighbourhood: "Cicendo",
				PostalCode:    "40172",
				Detail:        stringPtr("Jl. Cicendo Raya No. 8"),
				Religion:      "islam",
			},
		})
	}

	return db.Create(&addresses).Error
}
