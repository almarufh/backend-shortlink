# [BACKEND] - ShortLink API v1.0.0

<p>
  <a href="https://golang.org/"><img style="border-radius:4px" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go"></a>
  <a href="https://github.com/gin-gonic/gin"><img style="border-radius:4px" src="https://img.shields.io/badge/Gin-008B8B?style=for-the-badge&logo=go&logoColor=white" alt="Gin"></a>
  <a href="https://github.com/joho/godotenv"><img style="border-radius:4px" src="https://img.shields.io/badge/Godotenv-ECD53F?style=for-the-badge&logo=dotenv&logoColor=black" alt="Godotenv"></a>
  <a href="https://www.postgresql.org/"><img style="border-radius:4px" src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL"></a>
  <a href="https://github.com/jackc/pgx"><img style="border-radius:4px" src="https://img.shields.io/badge/pgx-339933?style=for-the-badge&logo=postgresql&logoColor=white" alt="pgx"></a>
  <a href="https://redis.io/"><img style="border-radius:4px" src="https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white" alt="Redis"></a>
</p>

A robust backend RESTful API built with Go and the Gin Framework, developed as part of the Koda Batch 9 curriculum. This project serves as a comprehensive backend service for short link management and URL redirection.

## Tech Stack

- [![Go](https://img.shields.io/badge/Go-v1.27.1-blue?logo=go&logoColor=white)](https://go.dev/learn/) <br>Digunakan sebagai bahasa pemrograman utama untuk membangun seluruh logika backend aplikasi _ShortLink API_ ini.
- [![Gin-Gonic](https://img.shields.io/badge/Gin_Gonic-v1.12.0-green?logo=Gin&logoColor=white)](https://gin-gonic.com/en/)<br>Dipakai sebagai _HTTP router_ dan web framework untuk mengatur _endpoint_ RESTful API fitur autentikasi, manajemen URL, serta redirection.
- [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-v18.6-blue?logo=PostgreSQL&logoColor=white)](https://www.postgresql.org/) <br>Menjadi basis data relasional utama untuk menyimpan data aplikasi secara permanen seperti akun pengguna, profil, dan pemetaan tautan pendek (original URL & short slug).
- [![pgx](https://img.shields.io/badge/Pgx-v5.11.0-blue?)](https://github.com/jackc/pgx) <br>Digunakan sebagai _driver_ dan _toolkit_ penghubung antara kode Go dan PostgreSQL untuk mengeksekusi _query_ database secara cepat dan efisien.
- [![Redis](https://img.shields.io/badge/Redis-v8.10.2-blue?logo=go&logoColor=white)](https://redis.io/) <br>Diterapkan untuk mengelola sesi pengguna, seperti penyimpanan session login.
- [![GodotEnv](https://img.shields.io/badge/GodotEnv-v1.5.1-blue?)](https://github.com/joho/godotenv) <br>Berfungsi untuk membaca file konfigurasi `.env` guna mengelola variabel rahasia dan dinamis (seperti koneksi database PostgreSQL, port server, dan kunci rahasia token JWT).

## Features

- User Authentication (Register, Login)
- Short Link Generator (Create Short URL)
- Link Management (List User Links, Delete Link)
- Fast URL Redirection via Slug

## Usage Instruction

### Clone this repository

```bash
$ git clone https://github.com/almarufh/backend-shortlink
```

### Redis Config Setup

Setup your configuration for your use cache on the root directory named **redis.conf** example [redis.conf](https://github.com/almarufh/backend-shortlink/blob/main/deployment/redis.conf.example)

### Environment Docker Compose Setup

Create your environment on the root directory named **.env** example use [.env.example](https://github.com/almarufh/backend-shortlink/blob/main/deployment/.env.example)

### Environment Backend Setup

Create your environment on the root directory named **.env** example use [.env.example](https://github.com/almarufh/backend-shortlink/blob/main/.env.example)

### Docker Pull Image PostgreSQL and Redis from [**DokerHub**](https://hub.docker.com/)

1. PosgresSQL [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-v18.6-blue?logo=PostgreSQL&logoColor=white)](https://www.postgresql.org/)

```bash
$ docker pull postgres:18.6-alpine3.24
```

2. Redis [![Redis](https://img.shields.io/badge/Redis-v8.10.2-blue?logo=go&logoColor=white)](https://redis.io/)

```bash
$ docker pull redis:8.10.2-alpine3.23
```

### Running Database and Chache

```bash
$ make up
```

3. Create Schema Table

```zsh
$ make migrate-up
```

4. Install dependency

```bash
$ go mod download
```

5. Running Backend

```bash
$ go run cmd/main.go
```

## Routes

| Endpoint             | Method   | Auth  | Description        |
| :------------------- | :------- | :---- | :----------------- |
| `/api/auth/register` | `POST`   | `NO`  | Create New Account |
| `/api/auth/login`    | `POST`   | `NO`  | Log In             |
| `/api/auth/logout`   | `DELETE` | `YES` | Log Out            |
| `/api/links`         | `POST`   | `YES` | Forgot Password    |
| `/api/links`         | `PATCH`  | `YES` | Forgot Password    |
| `/api/links/:slug`   | `DELETE` | `YES` | Forgot Password    |
| `/:slug`             | `GET`    | `NO`  | Forgot Password    |

### Documentation

For complete documentation, visit [docs swagger endpooint **`/api/docs/index.html`** ](http://localhost/api/docs/index.html)

## How to Contribute

- Fork this repository
- Create your changes
- Pull Request

## License [![License: MIT](https://img.shields.io/badge/License-MIT-red)](https://opensource.org/license/mit)

This project is licensed under the MIT License

## CONTACTS

[![Email](https://img.shields.io/badge/Email-almarufhidayat99@gmail.com-blue?logo=maildotru&logoColor=red)](mailto:almarufhidayat99@gmail.com)<br>
[![WhatsApp](https://img.shields.io/badge/WhatsApp-6281973779380-blue?logo=whatsapp&logoColor=green)](https://wa.me/6281973779380)

## Related Project

[![github](https://img.shields.io/badge/FRONTEND--blue?logo=googlechrome&logoColor=blue)](https://github.com/almarufh/frontend-shortlink)
