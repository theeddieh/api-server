FROM golang:1.13-alpine

ADD . /home
        
WORKDIR /home

CMD ["go","run","main.go"]