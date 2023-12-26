# Build stage
FROM golang:1.19 AS build-stage

WORKDIR /opt/app

COPY backend/ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o lr-app

# Run stage
FROM alpine:3.19 AS run-stage

COPY --from=build-stage /opt/app/lr-app /bin/lr-app

EXPOSE 8080

ENTRYPOINT ["/bin/lr-app"]

CMD ["go", "run", "main.go"]

