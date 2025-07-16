# Sistem Absensi SMK - Backend

Backend API untuk sistem absensi sekolah menggunakan Go, Fiber, GORM, MySQL, dan Redis.

## Tech Stack

- **Go** - Programming language
- **Fiber** - Web framework  
- **GORM** - ORM library
- **MySQL** - Database
- **Redis** - Cache (optional)

## Setup

### Prerequisites

- Go 1.21+
- MySQL 8.0+
- Redis (optional)

### Installation

1. Clone repository
```bash
git clone <repository-url>
cd absensi_be
```

2. Install dependencies
```bash
go mod download
```

3. Setup environment
```bash
cp .env.example .env
# Edit .env file dengan konfigurasi database Anda
```

4. Create database
```sql
CREATE DATABASE absensi_db;
```

5. Run migrations dan seeders
```bash
# Fresh setup dengan sample data
go run main.go -fresh

# Atau hanya migration
go run main.go -migrate

# Atau hanya seeder
go run main.go -seed
```

6. Start server
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## Database Structure

Project ini memiliki model untuk:
- Schools (Sekolah)
- Students (Siswa) 
- Teachers (Guru)
- Classes (Kelas)
- Subjects (Mata Pelajaran)
- Attendance (Absensi)
- dan lainnya

## Commands

```bash
# Reset database
go run main.go -reset

# Migration only
go run main.go -migrate

# Seeder only  
go run main.go -seed

# Fresh migration + seed
go run main.go -fresh
```