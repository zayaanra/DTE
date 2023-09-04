FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download
RUN apt-get update && apt-get install -y libgtk-3-dev libcairo2-dev libglib2.0-dev

CMD ["./cmd/server/server"]