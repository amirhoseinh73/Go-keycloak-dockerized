FROM golang:alpine
 
RUN mkdir /app

WORKDIR /app

COPY ./ .

RUN go mod init keycloak-go

RUN go get \
    github.com/gin-gonic/gin \
    github.com/golang-jwt/jwt/v4 \
    github.com/joho/godotenv \
    golang.org/x/crypto \
    gorm.io/driver/postgres \
    gorm.io/gorm


EXPOSE 8002

CMD [ "go", "run", "main.go" ]