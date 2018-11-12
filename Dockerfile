FROM golang:1.11.2-stretch AS build
WORKDIR /go/src/github.com/johscheuer/go-hello-webserver
COPY hello-webserver.go /go/src/github.com/johscheuer/go-hello-webserver/hello-webserver.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-webserver .


FROM scratch
COPY --from=build /go/src/github.com/johscheuer/go-hello-webserver/hello-webserver /hello-webserver
CMD ["/hello-webserver"]
EXPOSE 8000
