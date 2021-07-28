FROM golang:1.16-alpine3.13 as build
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o netWork_api

FROM gcr.io/distroless/static:nonroot
COPY --from=build /app/netWork_api /netWork_api
EXPOSE 8080/tcp
USER nonroot:nonroot
ENTRYPOINT ["/netWork_api"]