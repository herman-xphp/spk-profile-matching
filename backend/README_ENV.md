# Environment Configuration & Database Setup

## Overview
Dokumentasi untuk setup environment variables dan konfigurasi database untuk aplikasi dan unit test.

## File Environment

### .env.example
Template file untuk environment variables. Copy file ini ke `.env` dan sesuaikan dengan konfigurasi Anda.

### .env (Development)
File environment untuk development. File ini **tidak** di-commit ke git (ada di .gitignore).

### .env.test (Testing)
File environment khusus untuk testing (opsional). Bisa digunakan dengan tools seperti `godotenv` untuk load environment saat testing.

## Environment Variables

### Server Configuration
```env
PORT=8000                    # Port untuk server
GIN_MODE=debug              # Mode Gin: debug, release, test
```

### JWT Configuration
```env
SECRET_KEY=your-secret-key-here-change-in-production
```
**PENTING**: Pastikan untuk mengubah SECRET_KEY di production dengan value yang kuat dan aman!

### Database Configuration (Aplikasi)
```env
DB_USER=root
DB_PASSWORD=                 # Kosongkan jika tidak pakai password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=spk_profile_matching
```

### Test Database Configuration
```env
TEST_DB_NAME=spk_profile_matching_test
```
Jika `TEST_DB_NAME` tidak diset, akan menggunakan pattern `{DB_NAME}_test`.

## Database Separation

### Aplikasi Database
- **Nama**: `spk_profile_matching` (dari `DB_NAME`)
- **Digunakan untuk**: Production/Development aplikasi
- **Connection**: Melalui `database.ConnectDB()`

### Test Database
- **Nama**: `spk_profile_matching_test` (dari `TEST_DB_NAME` atau `DB_NAME_test`)
- **Digunakan untuk**: Unit testing
- **Connection**: Melalui `database.ConnectTestDB()`
- **Auto-cleanup**: Database di-clear sebelum setiap test

## Setup Database

### 1. Membuat Database

Jalankan script SQL untuk membuat database:

```bash
# Menggunakan docker-compose (otomatis)
docker-compose up -d

# Atau manual
mysql -u root -p < db-init/create-dbs.sql
```

Script akan membuat 2 database:
- `spk_profile_matching` - untuk aplikasi
- `spk_profile_matching_test` - untuk testing

### 2. Setup Environment File

```bash
# Copy template
cp .env.example .env

# Edit sesuai kebutuhan
nano .env
```

### 3. Verify Configuration

```bash
# Check environment variables
cat .env

# Test database connection
go run cmd/api/main.go
```

## Menjalankan Aplikasi

### Development Mode
```bash
# Load .env dan jalankan aplikasi
go run cmd/api/main.go
```

Aplikasi akan:
1. Load environment dari `.env`
2. Connect ke database `spk_profile_matching`
3. Auto-migrate schema jika diperlukan

### Test Mode
```bash
# Jalankan semua test
go test ./...

# Test akan otomatis:
# 1. Connect ke database `spk_profile_matching_test`
# 2. Auto-migrate schema
# 3. Clear semua data sebelum setiap test
# 4. Run tests
```

## Perbedaan Database Aplikasi vs Test

| Aspek | Aplikasi Database | Test Database |
|-------|------------------|---------------|
| Nama | `spk_profile_matching` | `spk_profile_matching_test` |
| Connection Function | `ConnectDB()` | `ConnectTestDB()` |
| Auto-migrate | Ya | Ya |
| Auto-cleanup | Tidak | Ya (sebelum setiap test) |
| Logger | Info level | Silent level |
| Environment | `.env` | `.env` atau `TEST_DB_NAME` |

## Best Practices

1. **Jangan pernah commit `.env`** - File ini ada di `.gitignore`
2. **Gunakan `.env.example`** sebagai template untuk team
3. **Pisahkan database** - Selalu gunakan database terpisah untuk test
4. **SECRET_KEY yang kuat** - Gunakan random string yang panjang untuk production
5. **Environment per environment** - Buat `.env` berbeda untuk dev, staging, production

## Troubleshooting

### Database connection error
```bash
# Check MySQL service
sudo systemctl status mysql

# Check database exists
mysql -u root -p -e "SHOW DATABASES;"

# Create database manually
mysql -u root -p
CREATE DATABASE spk_profile_matching;
CREATE DATABASE spk_profile_matching_test;
```

### Test database not found
```bash
# Pastikan TEST_DB_NAME di set di .env atau environment
export TEST_DB_NAME=spk_profile_matching_test

# Atau create manual
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS spk_profile_matching_test;"
```

### Password authentication error
```bash
# Jika MySQL menggunakan password, set di .env
DB_PASSWORD=your_password
```

## Environment Variables Priority

1. Environment variables dari sistem
2. `.env` file (jika menggunakan godotenv)
3. Default values di code

## Security Notes

- Jangan hardcode credentials di code
- Gunakan environment variables untuk semua sensitive data
- Rotate SECRET_KEY secara berkala
- Jangan commit `.env` file ke git
- Gunakan different SECRET_KEY untuk setiap environment

