FROM golang:1.21.6-alpine AS build
WORKDIR /delivery
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o accounts ./cmd/accounts

FROM scratch
WORKDIR /delivery
COPY --from=build /delivery/accounts .
EXPOSE 80
CMD ["./accounts"]