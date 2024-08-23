# image
FROM golang:1.23

# where we are working
WORKDIR /app

# copy all files in this directory to working directory inside container
COPY . /app

# install dependencies
RUN go mod download

# build
RUN go build -o main ./src

# run
#CMD ["go", "run", "./src/main.go", "&"]
#CMD ["go", "run", "./src/apicomm.go", "&"]
CMD ["./main"]
