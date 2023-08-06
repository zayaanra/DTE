FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download
RUN apt-get update && apt-get install -y libgl1-mesa-dev xorg-dev x11-apps dbus-x11 xvfb make
RUN make

ENTRYPOINT ["./cmd/server/server"]