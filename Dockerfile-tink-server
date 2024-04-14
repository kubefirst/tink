FROM golang:latest AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN echo hello1

COPY . .
RUN go mod download

RUN go build -o /cmd/tink-server/tink-server ./cmd/tink-server

###########################################################################
# runtime stage

# build final image using nothing but the binary
FROM alpine:3.19

COPY --from=builder /cmd/tink-server/tink-server /usr/bin/tink-server

# Command to run
EXPOSE 42113 42114
RUN apk add --update --upgrade --no-cache ca-certificates

ENTRYPOINT ["/usr/bin/tink-server"]
