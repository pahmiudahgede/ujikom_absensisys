package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type SchoolAddressSeeder struct{}

func (s *SchoolAddressSeeder) GetName() string {
	return "School Addresses"
}

func (s *SchoolAddressSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.SchoolAddress{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var schools []models.School
	if err := db.Find(&schools).Error; err != nil {
		return err
	}

	if len(schools) == 0 {
		return nil
	}

	addresses := []models.SchoolAddress{
		{
			SchoolID: schools[0].ID,
			AddressBase: models.AddressBase{
				Province:      "DKI Jakarta",
				District:      "Jakarta Pusat",
				Subdistrict:   "Gambir",
				Hamlet:        "Gambir",
				Village:       "Gambir",
				Neighbourhood: "Gambir",
				PostalCode:    "10110",
				Detail:        stringPtr("Jl. Medan Merdeka Selatan No. 11"),
				Religion:      "islam",
				Latitude:      float64Ptr(-6.1751),
				Longitude:     float64Ptr(106.8650),
			},
		},
	}

	if len(schools) > 1 {
		addresses = append(addresses, models.SchoolAddress{
			SchoolID: schools[1].ID,
			AddressBase: models.AddressBase{
				Province:      "DKI Jakarta",
				District:      "Jakarta Selatan",
				Subdistrict:   "Kebayoran Baru",
				Hamlet:        "Kebayoran Baru",
				Village:       "Kebayoran Baru",
				Neighbourhood: "Kebayoran Baru",
				PostalCode:    "12120",
				Detail:        stringPtr("Jl. Radio Dalam Raya No. 5"),
				Religion:      "islam",
				Latitude:      float64Ptr(-6.2297),
				Longitude:     float64Ptr(106.7831),
			},
		})
	}

	return db.Create(&addresses).Error
}
