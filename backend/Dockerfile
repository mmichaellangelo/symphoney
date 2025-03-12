FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy all src files
COPY . .

# Install dependencies
RUN apk add --update --no-cache \
    go

# Build application
RUN go build -o app

# Expose port
EXPOSE 8080

# Run the server
# CMD ["./app"]