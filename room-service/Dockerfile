FROM golang:1.22-alpine AS FirstBuild
WORKDIR /room-service/
COPY ./ ./
RUN go mod download
RUN go build -o ./cmd/build ./cmd/main.go
RUN apk add -U --no-cache ca-certificates

FROM scratch
COPY --from=FirstBuild /room-service/cmd/build /
COPY --from=FirstBuild /room-service/.env /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 50056
CMD [ "/build" ]