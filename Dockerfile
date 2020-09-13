FROM golang:1.15.1-buster
ARG gitUser
ARG gitPassword
WORKDIR /app
COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download
COPY src ./
RUN go build -o server .
EXPOSE 8801
CMD ["./server"]