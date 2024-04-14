FROM golang:latest AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

COPY . .
RUN go mod download

WORKDIR /cmd/tink-server

RUN go build .

###########################################################################
# runtime stage

# # Build final image using nothing but the binary
# # run `make image` to build the binary + container
# # if you're using `make build` this Dockerfile will not find the binary
# # and you probably want `make smee-linux-amd64`
FROM alpine:3.19

COPY --from=builder /cmd/smee/tink-server /usr/bin/tink-server

# Command to run
EXPOSE 42113 42114
RUN apk add --update --upgrade --no-cache ca-certificates

ENTRYPOINT ["/usr/bin/tink-server"]
