FROM golang:1.25-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app ./main.go

FROM alpine:3.20
WORKDIR /app
RUN adduser -D -H appuser && mkdir -p ./logs && chown -R appuser:appuser /app
COPY --from=build /out/app /app/app
USER appuser
EXPOSE 8080
ENTRYPOINT ["/app/app"]