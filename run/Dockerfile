FROM golang:1.19

# Set destination
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /authyreapi ./cmd

# Bind to a TCP port
EXPOSE 8080

# Run
CMD ["/authyreapi"]