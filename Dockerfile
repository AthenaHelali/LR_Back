# Build stage
FROM golang:1.19 AS build-stage

WORKDIR /opt/app

COPY ./ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o lr-app

# Run stage
FROM alpine:3.19 AS run-stage

WORKDIR /bin/

COPY --from=build-stage /opt/app/ ./ 

COPY --from=build-stage /opt/app/lr-app /bin/lr-app

EXPOSE 8088

ENTRYPOINT ["/bin/lr-app"]

