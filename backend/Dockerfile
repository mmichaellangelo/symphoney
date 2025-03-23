FROM golang:alpine3.21 as build

# Set working directory
WORKDIR /app

# Copy necessary files and download modules
COPY go.mod go.sum /
RUN go mod download

# Copy rest of source code
COPY . .

# Build application
RUN go build -o app

# Copy to distroless
FROM gcr.io/distroless/static
COPY --from=build /app /

# Set environment variables
ENV API_PORT 8080

# Expose port
EXPOSE 8080

# Run the server
CMD ["/app"]