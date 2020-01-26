
# Two stage docker build to reduce image overhead
FROM golang:1.13.6-alpine3.11 as build
WORKDIR /app
COPY . /app
RUN apk add make
RUN make build-example

FROM alpine:3.11.3
WORKDIR /app
COPY --from=build  /app/graphql-example .
EXPOSE 8080
ENTRYPOINT ["/app/graphql-example"]
