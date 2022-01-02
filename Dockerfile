FROM golang:alpine as builder

RUN mkdir /build
ADD . /build
WORKDIR /build

# Statically links the application to make it small.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o links .

FROM scratch
COPY --from=builder /build/links /links/
WORKDIR /links/
CMD ["./links"]
