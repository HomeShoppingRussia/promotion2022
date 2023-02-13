FROM golang:latest AS build

# install migrate to perform DB migration
WORKDIR /app
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz && mv migrate /usr/bin/

# install go-swagger
RUN curl -o /usr/local/bin/swagger -L'#' "https://github.com/go-swagger/go-swagger/releases/download/v0.29.0/swagger_linux_amd64" \
    && chmod +x /usr/local/bin/swagger

# copy module files first so that they don't need to be downloaded again if no change
COPY go.* ./
RUN go mod download
RUN go mod verify

# copy source files and build the binary
COPY .. .

RUN make swagger

RUN make build
CMD ["/app/loto"]