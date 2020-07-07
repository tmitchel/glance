FROM golang:1.13
WORKDIR /go/src/glance
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD [ "glance" ]

EXPOSE 8080