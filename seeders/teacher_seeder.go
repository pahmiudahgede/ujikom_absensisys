package seeders

import (
	"absensibe/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedTeachers(db *gorm.DB) error {
	log.Println("👨‍🏫 Seeding teachers...")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("guru123"), bcrypt.DefaultCost)

	teachers := []models.Teacher{
		{
			ID:           "550e8400-e29b-41d4-a716-446655440010",
			TeacherCode:  nil,
			PhotoProfile: stringPtr("/assets/teacherpict/kepalasekolah.jpg"),
			NUPTK:        "197205082001121001",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "H. RAHMAT DZAKIR, S.P.,M.Pd",
			Email:        "rahmat.dzakir@smkn1bandung.sch.id",
			Phone:        "081234567890",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1972, 5, 8, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Bandung",
			PositionType: "kepala sekolah",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440011",
			TeacherCode:  stringPtr("4"),
			PhotoProfile: stringPtr("/assets/teacherpict/junaidi_4.jpeg"),
			NUPTK:        "4063748650200013",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "H.M.JUNAIDI HNT,S.Pd",
			Email:        "junaidi.hnt@smkn1bandung.sch.id",
			Phone:        "081234567891",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1974, 8, 6, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Bandung",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440012",
			TeacherCode:  stringPtr("5"),
			PhotoProfile: stringPtr("/assets/teacherpict/mawardi_5.jpeg"),
			NUPTK:        "9563745648200453",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "MAWARDI,S.Pd",
			Email:        "mawardi@smkn1bandung.sch.id",
			Phone:        "081234567892",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1974, 5, 6, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Cimahi",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440013",
			TeacherCode:  stringPtr("15"),
			PhotoProfile: stringPtr("/assets/teacherpict/robi-asmara_15.png"),
			NUPTK:        "9533764666200033",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "ROBI ASMARA, S.Pd",
			Email:        "robi.asmara@smkn1bandung.sch.id",
			Phone:        "081234567893",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1976, 4, 6, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Sumedang",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440014",
			TeacherCode:  stringPtr("16"),
			PhotoProfile: stringPtr("/assets/teacherpict/suhartono_16.png"),
			NUPTK:        "8563762663131343",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "SUHARTONO, S.PdI",
			Email:        "suhartono@smkn1bandung.sch.id",
			Phone:        "081234567894",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1976, 2, 6, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Garut",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440015",
			TeacherCode:  stringPtr("20"),
			PhotoProfile: stringPtr("/assets/teacherpict/Bustanul-Arifin_20.png"),
			NUPTK:        "2550763664130163",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "BUSTANUL ARIFIN, S.Pd",
			Email:        "bustanul.arifin@smkn1bandung.sch.id",
			Phone:        "081234567895",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1976, 3, 6, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Bandung",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440016",
			TeacherCode:  stringPtr("24"),
			PhotoProfile: stringPtr("/assets/teacherpict/misna_24.jpeg"),
			NUPTK:        "8745735637200032",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "MISNA HIDAYATI, S.Pd",
			Email:        "misna.hidayati@smkn1bandung.sch.id",
			Phone:        "081234567896",
			Password:     string(hashedPassword),
			Gender:       "P",
			BirthDate:    time.Date(1973, 5, 7, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Bandung",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440017",
			TeacherCode:  stringPtr("29"),
			PhotoProfile: stringPtr("/assets/teacherpict/faizaturrohmi_29.jpeg"),
			NUPTK:        "9245771672130043",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "FAIZATURROHMI, S. Pd.",
			Email:        "faizaturrohmi@smkn1bandung.sch.id",
			Phone:        "081234567897",
			Password:     string(hashedPassword),
			Gender:       "P",
			BirthDate:    time.Date(1977, 1, 6, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Cimahi",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440018",
			TeacherCode:  stringPtr("34"),
			PhotoProfile: stringPtr("/assets/teacherpict/AMRULLAH_34.jpeg"),
			NUPTK:        "2856768669130172",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "AMRULLAH, S. Pd",
			Email:        "amrullah@smkn1bandung.sch.id",
			Phone:        "081234567898",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1976, 8, 6, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Sumedang",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
		{
			ID:           "550e8400-e29b-41d4-a716-446655440019",
			TeacherCode:  stringPtr("40"),
			PhotoProfile: stringPtr("/assets/teacherpict/rodi_40.png"),
			NUPTK:        "1853771672130142",
			SchoolID:     "550e8400-e29b-41d4-a716-446655440001",
			Fullname:     "RODI MAHENDRA HUSAEN,S.Pd",
			Email:        "rodi.mahendra@smkn1bandung.sch.id",
			Phone:        "081234567899",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1977, 1, 7, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Garut",
			PositionType: "Guru Mapel",
			Status:       "aktif",
		},
	}

	for _, teacher := range teachers {
		if err := db.FirstOrCreate(&teacher, models.Teacher{NUPTK: teacher.NUPTK}).Error; err != nil {
			return err
		}
	}

	log.Printf("✅ Successfully seeded %d teachers", len(teachers))
	return nil
}
