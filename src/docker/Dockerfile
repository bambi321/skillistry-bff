FROM golang@sha256:30baaea08c5d1e858329c50f29fe381e9b7d7bced11a0f5f1f69a1504cdfbf5e AS builder

ARG project_controller

WORKDIR /build

COPY . .
RUN go mod download -x

RUN GOOS=linux GOARCH=amd64 go build -o skillistry-bff $project_controller

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/skillistry-bff .
EXPOSE 8080
CMD ["./skillistry-bff"]