# Go Auth API

Bu proje, Go ile geliştirilmiş, güvenli ve ölçeklenebilir bir kimlik doğrulama (authentication) sistemidir. Projede JWT yerine modern ve daha güvenli bir standart olan PASETO (v4) kullanılmış, veri erişim katmanı tip güvenliği (type-safety) sağlamak amacıyla SQLC ile inşa edilmiştir.

## Kullanılan Teknolojiler

* **Dil:** Go (Golang)
* **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
* **Veritabanı:** PostgreSQL
* **Veri Erişim Katmanı:** [SQLC](https://sqlc.dev/)
* **Migration:** [golang-migrate](https://github.com/golang-migrate/migrate)
* **Kimlik Doğrulama:** [PASETO v4](https://github.com/aidanwoods/go-paseto)
* **Şifreleme:** Bcrypt
* **Konfigürasyon Yönetimi:** [Viper](https://github.com/spf13/viper)

## Ön Koşullar

Projeyi çalıştırmadan önce sisteminizde aşağıdakilerin kurulu olduğundan emin olun:
* [Go](https://golang.org/doc/install) (1.20+)
* [PostgreSQL](https://www.postgresql.org/download/)
* [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
* [sqlc CLI](https://docs.sqlc.dev/en/latest/overview/install.html) (Geliştirme yapılacaksa)

## Kurulum ve Çalıştırma

**1. Depoyu Klonlayın**
\`\`\`bash
git clone https://github.com/kullanici_adin/repo_adin.git
cd repo_adin
\`\`\`

**2. Ortam Değişkenlerini Ayarlayın**
Proje kök dizininde bir `.env` dosyası oluşturun ve aşağıdaki değişkenleri kendi sisteminize göre doldurun:
\`\`\`env
DB_SOURCE="postgresql://kullanici:sifre@localhost:5432/authdb?sslmode=disable"
SERVER_ADDRESS="0.0.0.0:8080"
PASETO_KEY="32_karakterlik_gizli_anahtar_yaz"
\`\`\`

**3. Veritabanını Hazırlayın**
PostgreSQL üzerinde `authdb` adında bir veritabanı oluşturun ve migration'ları çalıştırın:
\`\`\`bash
migrate -path internal/db/migration -database "$DB_SOURCE" -verbose up
\`\`\`

**4. Sunucuyu Başlatın**
\`\`\`bash
go run cmd/api/main.go
\`\`\`
Sunucu varsayılan olarak `8080` portunda çalışmaya başlayacaktır.

## API Uç Noktaları (Endpoints)

### POST `/register`
Yeni bir kullanıcı oluşturur.
* **Body:** `{"username": "testuser", "password": "secretpassword"}`
* **Başarılı Yanıt (200 OK):** Kullanıcı bilgileri (şifre hash'i hariç).

### POST `/login`
Mevcut bir kullanıcı ile giriş yapar ve PASETO token döndürür.
* **Body:** `{"username": "testuser", "password": "secretpassword"}`
* **Başarılı Yanıt (200 OK):** Token ve kullanıcı bilgileri.

### GET `/anasayfa` (Korumalı)
Sadece geçerli bir PASETO token'a sahip kullanıcıların erişebildiği test uç noktası.
* **Header:** `Authorization: Bearer <token>`
* **Başarılı Yanıt (200 OK):** Karşılama mesajı ve kullanıcı adı.

## Geliştirici Bilgileri

* Proje Mimari Dizaynı: Standard Go Project Layout
* Geliştirici: Bora KALKAN
