# Use rust base image
FROM public.ecr.aws/docker/library/rust:1.86-alpine3.20 AS builder

# Set the working directory
WORKDIR /app

# Install libraries
RUN apk add libc-dev 

# Copy the entire project
COPY ./healthcheck .

# Build the application
RUN cargo build --release

# Second stage for a smaller image
FROM public.ecr.aws/docker/library/alpine:3.21.3

# Set the working directory
WORKDIR /app

# Update libraries
RUN apk update

# Copy the built binary from the builder stage, binary name check when test build on local
COPY --from=builder /app/target/release/healthcheck .

# Create directory
RUN mkdir /tmp/ready

# Expose port
EXPOSE 8080

# Set the entrypoint to the Rust app
ENTRYPOINT ["./healthcheck"]