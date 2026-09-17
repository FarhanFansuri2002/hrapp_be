# HR App Server

Backend REST API untuk aplikasi Human Resources, dibuat menggunakan Node.js dan MySQL.

## Fitur Saat Ini

- Health check server.
- Mengambil daftar karyawan dari database MySQL.
- CORS untuk frontend lokal di `http://localhost:5173`.
- Logging method, path, dan durasi setiap request.

## Prasyarat

- Node.js `22` atau versi yang kompatibel
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
| `PORT` | `8080` | Port HTTP server |
| `SERVER_ADDRESS` | `:8080` | Fallback port untuk kompatibilitas konfigurasi lama |
| `MYSQL_DSN` | - | Format Go lama, tetap didukung |
| `MYSQL_HOST` | `127.0.0.1` | Host MySQL |
| `MYSQL_PORT` | `3306` | Port MySQL |
| `MYSQL_USER` | `root` | User MySQL |
| `MYSQL_PASSWORD` | kosong | Password MySQL |
| `MYSQL_DATABASE` | `hr_app` | Nama database |
| `MYSQL_URL` | - | URL koneksi MySQL, jika ingin memakai satu variable |

Contoh konfigurasi PowerShell:

```powershell
$env:MYSQL_USER = "root"
$env:MYSQL_PASSWORD = "password"
$env:MYSQL_DATABASE = "hr_app"
$env:PORT = "8080"
$env:SERVER_ADDRESS = ":8080"
```

## Menjalankan Server

```bash
cd Server
npm install
npm start
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
src/server.js            # Entry point, route, middleware, dan query MySQL
package.json             # Script dan dependency Node.js
migrations/              # SQL migration
```