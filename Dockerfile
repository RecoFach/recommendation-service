FROM golang:1.15 AS build

ENV GO111MODULE=on
# Set the Current Working Directory inside the container
WORKDIR /go/src/recommendation_go-service/

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Install the package
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o reco *.go
# This container exposes port 8080 to the outside world

FROM scratch
WORKDIR /go/src/recommendation_go-service/
COPY --from=build /go/src/recommendation_go-service/ /go/src/recommendation_go-service/
#COPY --from=build /go/src/recommendation_go-service/subjects.csv /go/src/recommendation_go-service
EXPOSE 2000
# Run the executable
CMD ["./reco"]