From golang:1.14.3
RUN mkdir /app
WORKDIR /app
COPY main.go .
COPY form.html .
COPY animal.html .
COPY information.html .
RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 app .
CMD ["./main"]
