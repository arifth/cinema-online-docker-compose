# Create an image from golang official and stage it as base 

FROM golang:1.19 as base

FROM base as  dev

# Use air binaries for live reloading 

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /final-task

ADD . .

RUN go mod tidy

CMD ["air"]

# ENTRYPOINT go build  && ./final-task
