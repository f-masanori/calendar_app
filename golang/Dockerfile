FROM golang:1.14.0

WORKDIR /go/src/golang

RUN apt-get update && apt-get upgrade
RUN apt-get install git
# マイグレーション用
RUN go get github.com/rubenv/sql-migrate/...

# Dockerfile更新したら　docker-compose up --build　でビルドし直しを忘れない
EXPOSE 8080

CMD ["go", "run", "/go/src/golang/main.go"]