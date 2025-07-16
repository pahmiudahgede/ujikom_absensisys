package seeders

import (
	"absensibe/models"
	"fmt"

	"gorm.io/gorm"
)

type AddressSeeder struct {
	BaseSeeder
}

func (s *AddressSeeder) Run(db *gorm.DB) error {
	if DataExists(db, &models.Address{}, "reference_type = ?", "school") {
		return nil
	}

	var schools []models.School
	db.Find(&schools)

	for i, school := range schools {
		address := models.Address{
			ReferenceID:   school.ID,
			ReferenceType: "school",
			Province:      "DKI Jakarta",
			District:      "Jakarta Pusat",
			Subdistrict:   "Menteng",
			Hamlet:        "Menteng",
			Village:       "Menteng",
			Neighbourhood: fmt.Sprintf("RT.001/RW.00%d", i+1),
			PostalCode:    "10310",
			Detail:        stringPtr(fmt.Sprintf("Jl. Pendidikan No. %d", i+1)),
			Religion:      "islam",
			Latitude:      float64Ptr(-6.200000),
			Longitude:     float64Ptr(106.816666),
		}
		db.Create(&address)
	}

	return nil
}
