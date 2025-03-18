FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY ./go-http-server/go.mod ./
COPY ./go-http-server/go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY ../../go-http-server/ .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
# NOT USING BECAUSE I'M SETTING NETWORK TO HOST!
#EXPOSE 9096

ENV DATABASE_URL="postgres://postgres:postgres@localhost:2222/postgres?sslmode=disable"
ENV PORT=9096
RUN echo "DATABASE_URL is $DATABASE_URL"

# Command to run the executable
CMD ["./main"]
