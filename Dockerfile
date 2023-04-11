FROM golang:1.20-alpine

ADD . /home
        
WORKDIR /home

CMD ["go","run","main.go"]