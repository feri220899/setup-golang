# 🚀 Golang REST API Starter

Project REST API menggunakan Golang dengan struktur mirip Laravel. Menggunakan:

- Gin (web framework)
- GORM (ORM)
- Viper (.env config)
- Mendukung MySQL dan PostgreSQL

---

## 📦 Instalasi Cepat

```bash
# Clone repository
git clone https://github.com/feri220899/setup-golang.git
cd project-api

# Buat file .env
echo "DB_DRIVER=mysql
DB_USER=root
DB_PASS=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=golang_api" > .env

# Install dependency
go mod tidy

# Jalankan server
go run main.go

# Atau build jadi binary Windows
go build -o app.exe
./app.exe
