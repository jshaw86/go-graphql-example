
# Two stage docker build to reduce image overhead
FROM golang:1.13.6-alpine3.11 as build
WORKDIR /app
COPY . /app
RUN ls -la
RUN apk add make
RUN make build-graphql

FROM alpine:3.11.3
WORKDIR /app
COPY --from=build  /app/graphql .
EXPOSE 8080
ENTRYPOINT ["/app/graphql"]
