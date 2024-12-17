# Build image
FROM golang:alpine3.20 AS build

WORKDIR /app

# maybe add go.sum in the future
COPY go.mod ./

RUN go mod download && go mod verify

COPY ./src ./src

RUN CGO_ENABLED=0 GOOS=linux go build -v -o pipeline-app ./src/

RUN apk add --no-cache ca-certificates


# Application image scratch
FROM scratch AS prd

WORKDIR /app

EXPOSE 8000

ENV PORT=8000

COPY --from=build /app/pipeline-app /app/

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "/app/pipeline-app" ]