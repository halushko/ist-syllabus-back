FROM golang:1.25-alpine AS build
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/app ./main.go

FROM alpine:3.20
WORKDIR /app
RUN adduser -D -H appuser
USER appuser

COPY --from=build /out/app /app/app
RUN mkdir "logs"

EXPOSE 8080
ENTRYPOINT ["/app/app"]