FROM golang:1.21.2-alpine AS builder

WORKDIR /snapvault

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Link statically with CGO_ENABLED=0 for use in a minimal image
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /snapvault/snapvault .

FROM scratch

WORKDIR /snapvault
COPY --from=builder /snapvault/snapvault .

# Override with 'docker run -e DATABSE_URL=postgres://<user>:<password>@localhost:5432/<db>'
ENV DATABASE_URL="postgres://acmeuser:acmepassword@localhost:5432/acmedb"

ENTRYPOINT ["./snapvault"]
