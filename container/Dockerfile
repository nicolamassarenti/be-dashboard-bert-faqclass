# Build stage
FROM golang:1.14 as build
WORKDIR /go/src/app
COPY package/ .
RUN go build -v -o /app .

# Copying from build stage
FROM gcr.io/distroless/base
COPY --from=build /app /app
CMD ["/app"]