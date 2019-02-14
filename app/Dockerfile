FROM golang:latest AS builder

COPY . .

RUN go get -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /main .

FROM scratch
COPY --from=builder /main ./
EXPOSE 8080
ENTRYPOINT ["./main"]