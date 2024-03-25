FROM golang:1.22 as build
WORKDIR /app
COPY . .
COPY index.html /app/html
RUN CGO_ENABLED=0 go build -o server main.go

FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/server .
COPY index.html /app/server/html
CMD ["./server"]