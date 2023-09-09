# Use golang latest image as the base image
FROM golang:latest

# Set working directory
WORKDIR /app

# Copy the application code into the container
COPY ./src .

# Build the Golang application

RUN go build -o main .

# Expose port 5000
EXPOSE 5000

#Run the application
CMD ["./main"]