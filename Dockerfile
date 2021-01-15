FROM golang:1.15.6 as builder
WORKDIR /go/src/github.com/voiceapps
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
WORKDIR /
COPY --from=builder /go/src/github.com/voiceapps/app .

# Health Check for the service
HEALTHCHECK --timeout=5s --interval=3s --retries=3 CMD curl --fail http://localhost:8080/v1/health || exit 1

# Expose the application on port 8080.
# This should be the same as in the app.conf file
EXPOSE 8080

# Set the entry point of the container to the application executable
CMD ["/app"]
