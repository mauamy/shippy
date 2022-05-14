FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /app
ENV GO111MODULE=on

COPY . .

WORKDIR /app/shippy-cli-consignment
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-cli-consignment

RUN pwd
RUN ls -lah

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
ADD shippy-cli-consignment/consignment.json /app/consignment.json
COPY --from=builder /app/shippy-cli-consignment/shippy-cli-consignment .

RUN pwd
RUN ls -lah

CMD ["/app/shippy-cli-consignment"]
