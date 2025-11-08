#!/bin/bash

# Script untuk setup environment file (.env) dari template

ENV_FILE=".env"
ENV_EXAMPLE=".env.example"

# Check if .env already exists
if [ -f "$ENV_FILE" ]; then
    echo "File .env sudah ada."
    echo "Apakah Anda ingin membuat backup dan membuat ulang? (y/n)"
    read -r response
    if [[ "$response" =~ ^[Yy]$ ]]; then
        cp "$ENV_FILE" "${ENV_FILE}.backup.$(date +%Y%m%d_%H%M%S)"
        echo "Backup dibuat: ${ENV_FILE}.backup.$(date +%Y%m%d_%H%M%S)"
    else
        echo "Membatalkan operasi."
        exit 0
    fi
fi

# Check if .env.example exists
if [ ! -f "$ENV_EXAMPLE" ]; then
    echo "Error: File .env.example tidak ditemukan!"
    echo "Membuat .env.example dari template..."
    cat > "$ENV_EXAMPLE" << 'EOF'
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
EOF
    echo "File .env.example telah dibuat."
fi

# Copy from .env.example to .env
cp "$ENV_EXAMPLE" "$ENV_FILE"
echo "File .env telah dibuat dari .env.example"
echo ""
echo "⚠️  PENTING: Edit file .env dan sesuaikan dengan konfigurasi Anda:"
echo "   - SECRET_KEY: Ganti dengan secret key yang kuat"
echo "   - DB_PASSWORD: Set password MySQL jika diperlukan"
echo "   - DB_USER, DB_HOST, DB_PORT: Sesuaikan jika berbeda"
echo ""
echo "File: $ENV_FILE"

