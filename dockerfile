
FROM golang:1.23.4 AS builder

WORKDIR /app


RUN apt-get update && apt-get install -y \
    gcc \
    make


COPY . .


RUN go mod tidy


RUN GOOS=linux GOARCH=amd64 go build -o mygame-linux main.go
RUN GOOS=windows GOARCH=amd64 go build -o mygame.exe main.go
RUN GOOS=darwin GOARCH=amd64 go build -o mygame-macos main.go

# Etapa Final
FROM ubuntu:22.04

WORKDIR /app


RUN apt-get update && apt-get install -y \
    libx11-6 \
    libxrandr2 \
    libxcursor1 \
    libxi6 \
    libxinerama1 \
    libxxf86vm1 \
    libgl1-mesa-glx \
    libxext6 \
    libxfixes3


COPY --from=builder /app/mygame-linux ./mygame-linux
COPY --from=builder /app/mygame.exe ./mygame.exe
COPY --from=builder /app/mygame-macos ./mygame-macos


CMD ["./mygame-linux"]
