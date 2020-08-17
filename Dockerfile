FROM golang
WORKDIR /usr/app

COPY go.mod ./
RUN go mod download

COPY ./ ./

# prod ver. compiled code
# RUN go build src/main.go
# CMD [ "./main" ]

# dev ver.
CMD [ "go", "run", "src/main.go" ]