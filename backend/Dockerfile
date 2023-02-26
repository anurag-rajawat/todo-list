FROM golang:1.19-alpine AS build
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o /todo-list /app/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=build /app/.env .
COPY --from=build /todo-list /todo-list
ENTRYPOINT ["/todo-list"]
EXPOSE 8000