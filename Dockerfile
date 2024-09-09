FROM golang:alpine

WORKDIR /src

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o valorantAgentMeta .

CMD ["valorantAgentMeta"]

ENTRYPOINT ["./valorantAgentMeta"]