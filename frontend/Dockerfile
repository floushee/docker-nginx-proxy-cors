FROM golang:alpine
EXPOSE 3000
WORKDIR /
COPY main.go /
COPY static /static
CMD ["go", "run", "main.go"]
