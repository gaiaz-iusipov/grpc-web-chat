FROM golang:1.20 as build-env
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/main.go

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /app
CMD ["/app"]
