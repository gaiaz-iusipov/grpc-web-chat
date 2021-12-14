FROM golang:1.17 as build-env
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app .

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /app
ENV PORT=80
EXPOSE $PORT
CMD ["/app"]
