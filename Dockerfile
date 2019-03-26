FROM golang:1.12 as builder

# Change cwd
WORKDIR /go/src/registrationapi/

# Copy files
COPY . .

# Installs dependencies
RUN go get -d -v ./...

# Start api
CMD go run registration.go 
