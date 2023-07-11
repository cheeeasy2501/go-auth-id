ARG  BUILDER_IMAGE=golang:alpine3.15
############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates && apk add --virtual build-dependencies build-base gcc

WORKDIR $GOPATH/src/app/
COPY . .
COPY ./env.local /go/auth-id/env.local

# Fetch dependencies.
RUN go get -d -v ./

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/auth-id/main -a -ldflags '-linkmode external -extldflags "-static"' ./

############################
# STEP 2 build a small image
############################
# FROM scratch
FROM golang:alpine3.15
WORKDIR /go/auth-id

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/auth-id/ /go/auth-id/
# COPY --from=builder /go/auth-id/env.local /go/auth-id/.env.local
RUN ls -la
# Run the binary.
ENTRYPOINT ["/go/auth-id/main"]