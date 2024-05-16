# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22.3 AS build-stage

WORKDIR /app

COPY main.go go.mod go.sum ./
RUN go mod download

COPY ./internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o garnet-userapi

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/garnet-userapi /garnet-userapi

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/garnet-userapi"]
