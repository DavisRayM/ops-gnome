FROM golang:1.20 AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /ops-server ./cmd/server

FROM ubuntu:22.10 AS release

ENV DEBIAN_FRONTEND noninteractive

WORKDIR /

COPY --from=build /ops-server /ops-server

CMD ["/ops-server"]
