FROM golang:1.15.1-buster as build
WORKDIR /app
COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download
COPY src ./
RUN go build -o server .

FROM scratch AS bin-linux
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=build /app/server /
EXPOSE 8801
CMD ["/server"]