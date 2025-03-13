FROM alpine:3.21

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache go

# Copy necessary files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of source code
COPY . .

# Build application
RUN go build -o app

# Expose port
EXPOSE 8080

# Run the server
CMD ["./app"]