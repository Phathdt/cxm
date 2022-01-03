FROM golang:1.17.5-alpine as build-env
WORKDIR /app
RUN apk update && apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash
WORKDIR /app
RUN chown nobody:nobody /app
USER nobody:nobody
COPY --from=build-env --chown=nobody:nobody /app/bin ./bin
COPY --from=build-env --chown=nobody:nobody /app/run.sh .

ENTRYPOINT sh run.sh
