FROM golang:1.15.1-buster as build
WORKDIR /app
COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download
COPY src ./
RUN CGO_ENABLED=0 go build -o server .
RUN useradd -u 10001 appuser

FROM scratch AS bin-linux
COPY --from=build /app/server /app/server
COPY --from=build /etc/passwd /etc/passwd
USER appuser
EXPOSE 8801
CMD ["/app/server"]