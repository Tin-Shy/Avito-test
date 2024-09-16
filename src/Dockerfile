# Строим приложение на базе Go
FROM golang:1.20 AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server main.go

# Используем Debian для поддержки CGO
FROM alpine:3.16
RUN apk add --no-cache bash curl && \
    curl -Lo /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    curl -Lo glibc-2.34-r0.apk https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.34-r0/glibc-2.34-r0.apk && \
    apk add --no-cache --allow-untrusted --force-overwrite glibc-2.34-r0.apk && \
    rm -f glibc-2.34-r0.apk



COPY --from=build /app/server /app/server
EXPOSE 8080
CMD ["/app/server"]
