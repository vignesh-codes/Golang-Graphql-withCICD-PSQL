FROM golang:1.14.0-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modchache
RUN go build -o main .
CMD ["/app/main"]
