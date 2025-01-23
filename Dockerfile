FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /kirana-club-assignment
RUN ls -l /app

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/kirana-club-assignment .

RUN chmod +x /root/kirana-club-assignment
RUN ls -l /root/

COPY .env .
CMD ["./kirana-club-assignment"]
