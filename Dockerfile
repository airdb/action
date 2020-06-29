# Stage 1: Builder
FROM golang AS builder
MAINTAINER info@airdb.com
WORKDIR $GOPATH/src/github.com/airdb/github

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 go build

RUN pwd && ls -l

# Stage 2: Release the binary from the builder stage
FROM scratch
COPY --from=builder /go/src/github.com/airdb/github/github /bin/github

ENTRYPOINT ["/bin/github"]
