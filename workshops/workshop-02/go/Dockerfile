# Build the application from source
FROM public.ecr.aws/docker/library/golang:1.24.2-alpine3.21 AS builder

WORKDIR /app

# Copy the rest of the application code
COPY . ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /healthcheck

# Run the tests in the container
FROM builder AS run-test
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM public.ecr.aws/docker/library/alpine:3.21.3 AS build-release-stage

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /healthcheck /healthcheck

# Create directory
RUN mkdir /tmp/ready

# Expose port 8080 for the application
EXPOSE 8080

# Set the entrypoint to the Go app
ENTRYPOINT ["/healthcheck"]
