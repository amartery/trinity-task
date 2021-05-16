FROM golang:latest

LABEL maintainer="amartery@gmail.com"

ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# install migrate
RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install -y migrate

# make wait-for-postgres.sh executable
RUN chmod +x ./scripts/wait-for-postgres.sh

# build go app
RUN go mod download
RUN make build

EXPOSE 8080

CMD ["./main"]