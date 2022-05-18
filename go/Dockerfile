FROM golang:1.17 AS builder

WORKDIR /user/src/app

# Separate dependency commandas to better use cache
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# Build specifying OS and ARCH, and also statically linking (tags netgo to use Go Net lib, and ldflags -s), and stripping debug symbols
# The result will be a binary stripped of unecessary data and focused on a single OS and ARCH - thus, smaller
# Also, with all dependencies linked statically, so we can run it without anything else on the system
RUN GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags="-w -s" -v -o app

FROM scratch

USER 1000

# Copy statically linked binary and configure startup
COPY --from=builder /user/src/app/app /usr/local/bin/app
EXPOSE 8899
ENTRYPOINT ["/usr/local/bin/app"]