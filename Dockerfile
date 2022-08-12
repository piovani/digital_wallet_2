FROM golang:1.19-alpine 

WORKDIR /src

ENV PATH="go/bin:${PATH}"

ENV CGO_ENABLED=0

COPY . .

RUN go mod download

RUN go build -o service

# CMD ["tail", "-f", "/dev/null"]
