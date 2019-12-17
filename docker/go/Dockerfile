FROM golang

WORKDIR /go/src/go_docker

# RUN git config --global url."https://0e3360c428d1f1f6e6bfcfbb43d5a0776896ad77:x-oauth-basic@github.com/".insteadOf "https://github.com/"
###開発環境ではvolumeさせる(現在はさせない)
# RUN go get github.com/f-masanori/mynikki_dev/docker/go/conf
RUN go get github.com/go-sql-driver/mysql;
RUN go get github.com/gorilla/mux;
RUN go get -v github.com/rubenv/sql-migrate/... 
RUN go get github.com/DATA-DOG/go-sqlmock
RUN go get github.com/gorilla/handlers
RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]

EXPOSE 8080

CMD ["go", "run", "/go/src/go_docker/main.go"]