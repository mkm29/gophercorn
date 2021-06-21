# builder image
FROM golang:1.16-alpine3.13 as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o golang-memtest .


# generate clean, final image for end users
FROM alpine:3.14.0
COPY --from=builder /build/golang-memtest .

CMD [ "go", "run", "server.go" ]