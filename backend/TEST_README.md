# Unit Test Documentation

## Overview
Unit test suite untuk backend project menggunakan Go testing package dan testify untuk assertions.

## Struktur Test

### Repository Tests
- `internal/repositories/user_repository_test.go`
- `internal/repositories/jabatan_repository_test.go`
- `internal/repositories/aspek_repository_test.go`
- `internal/repositories/kriteria_repository_test.go`
- `internal/repositories/target_profile_repository_test.go`
- `internal/repositories/tenaga_kerja_repository_test.go`
- `internal/repositories/nilai_tenaga_kerja_repository_test.go`

### Service Tests
- `internal/services/user_service_test.go`
- `internal/services/jabatan_service_test.go`
- `internal/services/aspek_service_test.go`
- `internal/services/kriteria_service_test.go`
- `internal/services/target_profile_service_test.go`
- `internal/services/tenaga_kerja_service_test.go`
- `internal/services/nilai_tenaga_kerja_service_test.go`
- `internal/services/profile_matching_service_test.go`

### Controller Tests
- `internal/controllers/auth_controller_test.go`
- `internal/controllers/user_controller_test.go`
- `internal/controllers/jabatan_controller_test.go`
- `internal/controllers/aspek_controller_test.go`
- `internal/controllers/kriteria_controller_test.go`
- `internal/controllers/target_profile_controller_test.go`
- `internal/controllers/tenaga_kerja_controller_test.go`
- `internal/controllers/nilai_tenaga_kerja_controller_test.go`
- `internal/controllers/profile_matching_controller_test.go`

### DTO Tests
- `internal/dto/mapper_test.go`

## Menjalankan Test

### Menjalankan semua test
```bash
go test ./...
```

### Menjalankan test dengan coverage
```bash
go test -cover ./...
```

### Menjalankan test dengan verbose output
```bash
go test -v ./...
```

### Menjalankan test untuk package spesifik
```bash
go test ./internal/repositories
go test ./internal/services
go test ./internal/controllers
```

### Menjalankan test dengan coverage report
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Setup

### 1. Environment Configuration
Copy `.env.example` ke `.env` dan sesuaikan konfigurasi:

```bash
cp .env.example .env
```

Pastikan `TEST_DB_NAME` di-set:
```env
TEST_DB_NAME=spk_profile_matching_test
```

### 2. Database Setup
Buat database test terpisah dari database aplikasi:

```bash
# Menggunakan docker-compose
docker-compose up -d

# Atau manual
mysql -u root -p < db-init/create-dbs.sql
```

Database yang dibuat:
- `spk_profile_matching` - untuk aplikasi
- `spk_profile_matching_test` - untuk testing

**PENTING**: Test database akan di-clear sebelum setiap test untuk memastikan test isolation.

Untuk detail lebih lanjut, lihat [README_ENV.md](./README_ENV.md).

## Requirements

### Database
Test menggunakan database MySQL terpisah: `spk_profile_matching_test`

### Environment Variables
- `DB_USER` (default: root)
- `DB_PASSWORD` (default: kosong, set jika MySQL menggunakan password)
- `DB_HOST` (default: 127.0.0.1)
- `DB_PORT` (default: 3306)
- `DB_NAME` (default: spk_profile_matching) - untuk aplikasi
- `TEST_DB_NAME` (default: spk_profile_matching_test) - untuk testing
- `SECRET_KEY` (untuk JWT testing)

## Test Coverage

### Repository Layer
- ✅ Create operations
- ✅ Read operations (GetByID, GetAll)
- ✅ Update operations
- ✅ Delete operations
- ✅ Special queries (FindByEmail, GetByJabatanID, etc.)
- ✅ Error handling (Not Found, etc.)

### Service Layer
- ✅ Business logic validation
- ✅ Data transformation
- ✅ Error handling
- ✅ Password hashing
- ✅ Duplicate checking
- ✅ Complex calculations (Profile Matching)

### Controller Layer
- ✅ HTTP request handling
- ✅ Request validation
- ✅ Response formatting
- ✅ Error responses
- ✅ Status codes
- ✅ DTO mapping

## Best Practices

1. **Isolation**: Setiap test menggunakan database yang bersih
2. **Setup/Teardown**: Database di-clear sebelum setiap test
3. **Assertions**: Menggunakan testify untuk assertions yang jelas
4. **Test Data**: Menggunakan test data yang realistis
5. **Error Cases**: Test untuk error cases dan edge cases
6. **Integration**: Test untuk integrasi antar layer

## Notes

- Test database akan di-migrate otomatis sebelum test
- Semua tabel di-clear sebelum setiap test untuk isolasi
- Test menggunakan real database (tidak menggunakan mock)
- Pastikan MySQL server berjalan sebelum menjalankan test

