# Use official Golang image as the base
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN go build -o z3ter-birthday-app

# Expose the application port
EXPOSE 8080

# Run the app
CMD ["./z3ter-birthday-app"]
