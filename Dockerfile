FROM golang:1.13 AS build  

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/src/app
COPY main.go .
RUN go build -o main main.go

FROM scratch
COPY --from=build /usr/src/app/main /app
EXPOSE 8081  
ENTRYPOINT [ "/app" ]   