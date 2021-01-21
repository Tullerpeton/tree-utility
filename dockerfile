FROM golang:1.9.2
COPY . .
CMD go test -v