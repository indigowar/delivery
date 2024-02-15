FROM golang:1.21.6-alpine AS build
WORKDIR /delivery
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o menu ./cmd/menu

FROM scratch
WORKDIR /delivery
COPY --from=build /delivery/menu .
EXPOSE 80
CMD ["./menu"]
