# syntax=docker/dockerfile:1

FROM golang:1.19 AS build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# TODO; avoid copying non-go files?
COPY . ./
#COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /watertemp

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /watertemp /watertemp

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/watertemp"]

# Run
CMD ["/watertemp"]
