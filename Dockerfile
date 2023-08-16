FROM golang:1.20-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY db/migration ./migration
COPY start.sh .
COPY wait-for.sh .
RUN sed -i 's/\r$//' start.sh  && \  
        sed -i 's/\r$//' wait-for.sh && \
        chmod +x start.sh && \ 
        chmod +x wait-for.sh
COPY docker-compose.yaml .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]