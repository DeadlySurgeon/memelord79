FROM golang:1.14-alpine AS builder

COPY . /go/src/deadly.surgery/memerlord79
WORKDIR /go/src/deadly.surgery/memerlord79

# Update our certificates incase our builder is out of date.
RUN apk --no-cache add ca-certificates upx git build-base gcc abuild binutils binutils-doc gcc-doc

# Make output directory.
RUN mkdir /.bin

RUN GOOS=linux GARCH=amd64 \
    go build \
        -o /.bin/memelord \
         -ldflags  "-s -w -installsuffix nocgo -linkmode external -extldflags -static" \
        ./cmd/main.go

# Stub in UPX here I guess
RUN upx --lzma /.bin/memelord

# We might not use Scratch in the end. For now though it's going to be used.
FROM scratch

# Copy the certs from our build container so stuff isn't broken.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /.bin/memelord /memelord

CMD ["/memelord"]