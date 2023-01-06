# FROM postgres:alpine
# ADD scripts/1_init.sql /docker-entrypoint-initdb.d
# # ADD scripts/2_copy.sql /docker-entrypoint-initdb.d

# RUN chmod a+r /docker-entrypoint-initdb.d/*

# EXPOSE 6666
# RUN export dockerhost=$(docker-machine ip)


# The base go-image
FROM golang:alpine

# RUN mkdir -p /go/src/github.com/pavel/gorm-psql
# WORKDIR /go/src/github.com/pavel/gorm-psql

# ADD . /go/src/github.com/pavel/gorm-psql

# RUN go get -v



 #--------------
#Create a directory for the app
RUN mkdir /app
 
# Copy all files from the current directory to the app directory
COPY . /app
 
# Set working directory
WORKDIR /app

# RUN find . -type f -name "*.go"
# RUN pwd
# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o server cmd/main.go
 
# Run the server executable

CMD [ "/app/server" ]
#---------------

# # Stage 1
# FROM golang:alpine as builder
# RUN apk update && apk add --no-cache git
# RUN mkdir /build 
# ADD . /build/
# WORKDIR /build
# RUN go get -d -v
# RUN go build -o clover-app .
# # Stage 2
# FROM alpine
# RUN adduser -S -D -H -h /app appuser
# USER appuser
# COPY --from=builder /build/ /app/
# WORKDIR /app
# CMD ["./clover-app"]