# Setup Environment File (.env)

## Format File .env yang Benar

File `.env` harus memiliki format berikut:

```env
# Server Configuration
PORT=8000
GIN_MODE=debug

# JWT Secret Key (PENTING: Ganti dengan secret key yang kuat di production!)
SECRET_KEY=your-secret-key-here-change-in-production

# Database Configuration (Aplikasi)
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=spk_profile_matching

# Test Database Configuration (optional, defaults to DB_NAME_test)
TEST_DB_NAME=spk_profile_matching_test
```

## Penjelasan Setiap Variable

### 1. PORT (Baris 1)
- **Default**: `8000`
- **Deskripsi**: Port yang digunakan oleh server HTTP
- **Contoh**: `PORT=8000`

### 2. GIN_MODE (Baris 2)
- **Default**: `debug`
- **Options**: `debug`, `release`, `test`
- **Deskripsi**: Mode untuk Gin framework
- **Contoh**: `GIN_MODE=debug`

### 3. SECRET_KEY (Baris 3-4)
- **Default**: Tidak ada (wajib diisi!)
- **Deskripsi**: Secret key untuk JWT token signing
- **PENTING**: 
  - Ganti dengan random string yang kuat di production
  - Minimal 32 karakter
  - Jangan commit ke git
- **Contoh**: `SECRET_KEY=my-super-secret-key-change-this-in-production-12345`

### 4. DB_USER (Baris 5)
- **Default**: `root`
- **Deskripsi**: Username untuk koneksi MySQL
- **Contoh**: `DB_USER=root`

### 5. DB_PASSWORD (Baris 6)
- **Default**: (kosong)
- **Deskripsi**: Password untuk koneksi MySQL
- **Note**: Jika MySQL tidak menggunakan password, biarkan kosong
- **Contoh**: `DB_PASSWORD=` atau `DB_PASSWORD=mypassword`

### 6. DB_HOST (Baris 7)
- **Default**: `127.0.0.1`
- **Deskripsi**: Host MySQL server
- **Contoh**: `DB_HOST=127.0.0.1` atau `DB_HOST=localhost`

### 7. DB_PORT (Baris 8)
- **Default**: `3306`
- **Deskripsi**: Port MySQL server
- **Contoh**: `DB_PORT=3306`

### 8. DB_NAME (Baris 9)
- **Default**: `spk_profile_matching`
- **Deskripsi**: Nama database untuk aplikasi
- **Contoh**: `DB_NAME=spk_profile_matching`

### 9. TEST_DB_NAME (Baris 10)
- **Default**: `spk_profile_matching_test` (otomatis dari DB_NAME + "_test")
- **Deskripsi**: Nama database untuk testing
- **Contoh**: `TEST_DB_NAME=spk_profile_matching_test`

## Cara Setup

### Metode 1: Menggunakan Script
```bash
chmod +x scripts/setup-env.sh
./scripts/setup-env.sh
```

### Metode 2: Manual
```bash
# Copy dari template
cp .env.example .env

# Edit file .env
nano .env
# atau
vim .env
```

### Metode 3: Membuat Langsung
Buat file `.env` di root folder `backend/` dengan isi seperti di atas.

## Checklist Setup

- [ ] File `.env` sudah dibuat
- [ ] `SECRET_KEY` sudah diubah (jangan gunakan default!)
- [ ] `DB_USER` sesuai dengan MySQL user
- [ ] `DB_PASSWORD` diisi jika MySQL menggunakan password
- [ ] `DB_NAME` sesuai dengan nama database aplikasi
- [ ] `TEST_DB_NAME` sesuai dengan nama database test
- [ ] File `.env` sudah di-ignore oleh git (cek `.gitignore`)

## Generate SECRET_KEY yang Kuat

### Menggunakan OpenSSL
```bash
openssl rand -base64 32
```

### Menggunakan Python
```python
import secrets
print(secrets.token_urlsafe(32))
```

### Menggunakan Go
```bash
go run -c "fmt.Println(randString(32))"
```

## Troubleshooting

### Error: "Error loading .env file"
- Pastikan file `.env` ada di root folder `backend/`
- Pastikan format file benar (tidak ada spasi sebelum nama variable)
- Pastikan tidak ada tanda kutip yang tidak perlu

### Error: "Invalid credentials" saat login
- Pastikan `SECRET_KEY` sudah di-set
- Pastikan `SECRET_KEY` sama antara saat login dan saat verifikasi token

### Error: "Could not connect to database"
- Pastikan MySQL server berjalan
- Pastikan `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT` benar
- Pastikan database `DB_NAME` sudah dibuat
- Test koneksi manual: `mysql -u $DB_USER -p -h $DB_HOST -P $DB_PORT`

### Test database tidak ditemukan
- Pastikan `TEST_DB_NAME` di-set di `.env`
- Atau pastikan database `{DB_NAME}_test` sudah dibuat
- Jalankan: `mysql -u root -p < db-init/create-dbs.sql`

## Contoh File .env untuk Development

```env
PORT=8000
GIN_MODE=debug
SECRET_KEY=dev-secret-key-change-in-production-12345678901234567890
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=spk_profile_matching
TEST_DB_NAME=spk_profile_matching_test
```

## Contoh File .env untuk Production

```env
PORT=8000
GIN_MODE=release
SECRET_KEY=<generate-strong-random-secret-key-here>
DB_USER=app_user
DB_PASSWORD=<secure-password>
DB_HOST=db.example.com
DB_PORT=3306
DB_NAME=spk_profile_matching_prod
TEST_DB_NAME=spk_profile_matching_test
```

**⚠️ PENTING untuk Production:**
- Gunakan `GIN_MODE=release`
- Gunakan `SECRET_KEY` yang sangat kuat (minimal 64 karakter random)
- Jangan hardcode password di file
- Gunakan environment variables atau secret management system
- Jangan commit file `.env` ke git

