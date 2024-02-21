# syntax=docker/dockerfile:1

FROM golang:latest

# Set destination for COP
WORKDIR /api

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy code
COPY . .

# Build
# RUN go get -d -v
# RUN go build -o /gamesapi
RUN CGO_ENABLED=0 GOOS=linux go build -o /gamesapi

EXPOSE 8080

# Run
CMD [ "/gamesapi" ]
