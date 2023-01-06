FROM golang:alpine
 
RUN mkdir /app

WORKDIR /app

COPY ./ .

RUN go get \
    github.com/gin-gonic/gin \
    github.com/golang-jwt/jwt/v4 \
    github.com/joho/godotenv \
    golang.org/x/crypto \
    gorm.io/driver/postgres \
    gorm.io/gorm


EXPOSE 5000

CMD [ "go", "run", "main.go" ]