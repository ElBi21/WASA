# Start from Go
FROM golang:latest

# Set WORKDIR
WORKDIR /src/

# Copy all files
COPY . .

# Build backend
RUN go mod vendor
RUN go build ./cmd/webapi/

# Expose PORT 3000
EXPOSE 3000

# Set webapi to be the main program
CMD ["/src/cmd/webapi"]