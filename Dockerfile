FROM golang:1.20-alpine

ADD . /home
        
WORKDIR /home

ARG BUILD
ENV build_version=$BUILD
RUN go build -o api-server -ldflags "-X main.version=${build_version}" -v main.go 

EXPOSE 8080

CMD ["./api-server"]