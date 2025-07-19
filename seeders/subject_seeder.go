package seeders

import (
	"absensibe/models"
	"log"

	"gorm.io/gorm"
)

func SeedSubjects(db *gorm.DB) error {
	log.Println("📚 Seeding subjects...")

	subjects := []models.Subject{

		{
			ID:          "550e8400-e29b-41d4-a716-446655440020",
			Code:        "BIND",
			Name:        "Bahasa Indonesia",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 2,
			Description: stringPtr("Mata pelajaran Bahasa Indonesia"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440021",
			Code:        "BING",
			Name:        "Bahasa Inggris",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 3,
			Description: stringPtr("Mata pelajaran Bahasa Inggris"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440022",
			Code:        "MTK",
			Name:        "Matematika",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 4,
			Description: stringPtr("Mata pelajaran Matematika"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440023",
			Code:        "PKN",
			Name:        "Pendidikan Kewarganegaraan",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 2,
			Description: stringPtr("Mata pelajaran Pendidikan Kewarganegaraan"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440024",
			Code:        "PAI",
			Name:        "Pendidikan Agama Islam",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 3,
			Description: stringPtr("Mata pelajaran Pendidikan Agama Islam"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440025",
			Code:        "PJOK",
			Name:        "Pendidikan Jasmani dan Olahraga",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 2,
			Description: stringPtr("Mata pelajaran Pendidikan Jasmani dan Olahraga"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440026",
			Code:        "PKK",
			Name:        "Produk Kreatif dan Kewirausahaan",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 5,
			Description: stringPtr("Mata pelajaran Produk Kreatif dan Kewirausahaan"),
			IsActive:    true,
		},

		{
			ID:          "550e8400-e29b-41d4-a716-446655440027",
			Code:        "PSASISSM",
			Name:        "Pemeliharaan Sasis Sepeda Motor",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 8,
			Description: stringPtr("Mata pelajaran Pemeliharaan Sasis Sepeda Motor"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440028",
			Code:        "PMESINSM",
			Name:        "Pemeliharaan Mesin Sepeda Motor",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 8,
			Description: stringPtr("Mata pelajaran Pemeliharaan Mesin Sepeda Motor"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440029",
			Code:        "PKLISTRIKAN",
			Name:        "Pemeliharaan Kelistrikan Sepeda Motor",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 6,
			Description: stringPtr("Mata pelajaran Pemeliharaan Kelistrikan Sepeda Motor"),
			IsActive:    true,
		},

		{
			ID:          "550e8400-e29b-41d4-a716-446655440030",
			Code:        "UPACARA",
			Name:        "Upacara Bendera",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 1,
			Description: stringPtr("Kegiatan upacara bendera setiap hari Senin"),
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440031",
			Code:        "IMTAQ",
			Name:        "Iman dan Takwa",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			CreditHours: 1,
			Description: stringPtr("Kegiatan pembinaan iman dan takwa setiap hari Jumat"),
			IsActive:    true,
		},
	}

	for _, subject := range subjects {
		if err := db.FirstOrCreate(&subject, models.Subject{Code: subject.Code}).Error; err != nil {
			return err
		}
	}

	log.Printf("✅ Successfully seeded %d subjects including all curriculum subjects", len(subjects))
	return nil
}
