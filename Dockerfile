# Start from a minimal Go runtime image
FROM golang:1.15

# Create a new directory for the program
RUN mkdir -p /app

# Set the working directory to the new directory
WORKDIR /app

# Copy the program source code into the working directory
COPY . .

# Compile the program into a single binary
RUN go build -o app

# Expose the default HTTP port (8080)
EXPOSE 8080

# Run the program when the container starts
CMD ["./app"]
