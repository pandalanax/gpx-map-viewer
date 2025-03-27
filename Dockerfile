FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build -o gpx-viewer .

EXPOSE 8080

CMD ["./gpx-viewer"]

