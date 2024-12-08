# Инициализируем проект
go mod init myblog-backend

# Устанавливаем фреймворк для API (Gin)
go get -u github.com/gin-gonic/gin

# Устанавливаем GORM для работы с PostgreSQL
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# Устанавливаем bcrypt для хеширования паролей
go get -u golang.org/x/crypto/bcrypt

# Устанавливаем JWT для токенов
go get -u github.com/dgrijalva/jwt-go