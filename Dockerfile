# FROM --platform=linux/arm64/v8 golang:1.22
FROM --platform=linux/arm64v8 golang:1.22
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
COPY *.html ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /argodemo
EXPOSE 8001
CMD ["/argodemo"]