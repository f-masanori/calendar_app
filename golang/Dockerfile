FROM golang:1.14.0

WORKDIR /go/src/golang

###開発環境ではvolumeさせる(現在はさせない)
# RUN go get github.com/f-masanori/mynikki_dev/docker/go/conf
RUN apt-get update && apt-get upgrade
RUN apt-get install git


# RUN apt add git

RUN go get github.com/go-sql-driver/mysql;
RUN go get github.com/gorilla/mux;

# Dockerfile更新したら　docker-compose up --build　でビルドし直しを忘れない
EXPOSE 8080

CMD ["go", "run", "/go/src/golang/main.go"]