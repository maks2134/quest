ARG TARGETARCH=amd64
FROM golang:1.25-alpine AS build-stage

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o /opt/app ./cmd/quest/main.go

FROM alpine:latest

WORKDIR /

RUN apk add --no-cache ca-certificates tzdata

RUN addgroup -g 65532 nonroot && \
    adduser -D -u 65532 -G nonroot nonroot

COPY --from=build-stage /opt/app /opt/app

COPY --from=build-stage --chown=65532:65532 /app/migrations /migrations

COPY --from=build-stage --chown=65532:65532 /app/docs /docs

COPY --from=build-stage --chown=65532:65532 /app/swagger-ui /swagger-ui

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["/opt/app"]
