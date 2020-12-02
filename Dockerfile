# Not compressed lol
FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o service .
CMD ["./service"]