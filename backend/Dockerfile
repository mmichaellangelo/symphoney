##################################
# BASE STAGE
FROM golang:alpine3.21 as base
WORKDIR /app
COPY go.mod go.sum /
RUN go mod download
COPY . .
RUN go build -o app

##################################
# DEV TARGET
FROM base as dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]

##################################
# PROD TARGET
FROM gcr.io/distroless/static as prod
COPY --from=base /app/app /app
EXPOSE 8080
CMD ["/app"]