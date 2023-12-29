FROM golang:1.21.5-alpine AS builder
WORKDIR /delivery
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o delivery ./cmd/delivery

FROM scratch
WORKDIR /delivery
COPY --from=builder /delivery/delivery .
EXPOSE 3000
CMD [ "./delivery" ]