FROM golang:1.14 as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build 

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app/argo-events-custom-trigger .
CMD ["./argo-events-custom-trigger"]  