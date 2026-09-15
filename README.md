# HR App Server

Backend REST API untuk aplikasi Human Resources, dibuat menggunakan Go dan MySQL.

## Fitur Saat Ini

- Health check server.
- Mengambil daftar karyawan dari database MySQL.
- CORS untuk frontend lokal di `http://localhost:5173`.
- Logging method, path, dan durasi setiap request.

## Prasyarat

- Go `1.27.1` atau versi yang kompatibel dengan `go.mod`
- MySQL 8 atau versi yang kompatibel

## Setup Database

Buat database terlebih dahulu:

```sql
CREATE DATABASE hr_app;
```

Jalankan migrasi dari root repository:

```bash
mysql -u root -p hr_app < Server/migrations/001_create_employees.sql
```

Migrasi membuat tabel `employees` dan memasukkan tiga data awal.

## Konfigurasi

Server membaca environment variable berikut:

| Variable | Default | Keterangan |
| --- | --- | --- |
| `SERVER_ADDRESS` | `:8080` | Alamat dan port HTTP server |
| `MYSQL_DSN` | `root:@tcp(127.0.0.1:3306)/hr_app?parseTime=true` | Connection string MySQL |

Contoh konfigurasi PowerShell:

```powershell
$env:MYSQL_DSN = "root:password@tcp(127.0.0.1:3306)/hr_app?parseTime=true"
$env:SERVER_ADDRESS = ":8080"
```

## Menjalankan Server

```bash
cd Server
go mod download
go run ./cmd/api
```

Server berjalan di `http://localhost:8080`.

## Endpoint

### Health check

```http
GET /health
```

Contoh respons:

```json
{"status":"ok"}
```

### Daftar karyawan

```http
GET /api/v1/employees
```

Contoh respons:

```json
[
	{
		"id": "EMP-001",
		"name": "Aulia Sari",
		"role": "People Operations",
		"department": "People",
		"status": "Aktif",
		"joinedAt": "2024-03-11"
	}
]
```

Contoh pengecekan dengan curl:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/employees
```

## Struktur Utama

```text
cmd/api/                 # Entry point aplikasi
internal/config/         # Pembacaan konfigurasi
internal/handler/        # HTTP handler dan serialisasi JSON
internal/model/          # Model response API
internal/repository/     # Query MySQL
internal/router/         # Route dan middleware
internal/service/        # Logika bisnis
migrations/              # SQL migration
```