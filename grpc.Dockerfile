FROM golang:1.25.7-trixie AS builder

WORKDIR /app

# Copying go.mod and go.sum enables the build to leverage Docker's cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o . ./...

FROM scratch
WORKDIR /

COPY --from=builder /app/grpc /grpc

# Command to run the executable when the container starts
CMD ["/grpc"]
