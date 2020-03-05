FROM golang

WORKDIR /go/src/go_docker

###開発環境ではvolumeさせる(現在はさせない)
# RUN go get github.com/f-masanori/mynikki_dev/docker/go/conf


RUN go get github.com/go-sql-driver/mysql;
RUN go get github.com/gorilla/mux;
RUN go get -v github.com/rubenv/sql-migrate/... 
# RUN go get github.com/DATA-DOG/go-sqlmock
RUN go get github.com/gorilla/handlers
RUN go get firebase.google.com/go
RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]
# Dockerfile更新したら　docker-compose up --build　でビルドし直しを忘れない
EXPOSE 8080

CMD ["go", "run", "/go/src/go_docker/main.go"]