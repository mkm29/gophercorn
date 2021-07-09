# builder image
FROM golang:1.16-alpine3.13 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o gophercorn .


# generate clean, final image for end users
FROM alpine:3.14.0
COPY --from=builder /build/gophercorn .

EXPOSE 8080

CMD ["./gophercorn"]