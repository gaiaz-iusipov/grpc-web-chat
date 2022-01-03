FROM golang:1.17 as build-env
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/main.go

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /app
ENV GRPC_PORT=3000 \
    DEBUG_PORT=6060
EXPOSE $DEBUG_PORT/tcp
CMD ["/app"]
