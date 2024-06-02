FROM golang:1.22-alpine AS FirstBuild
WORKDIR /api-gateway/
COPY ./ ./
RUN go mod download
RUN go build -o ./cmd/build ./cmd/main.go
RUN apk add -U --no-cache ca-certificates

FROM scratch
COPY --from=FirstBuild /api-gateway/cmd/build /
COPY --from=FirstBuild /api-gateway/template/ /template/
COPY --from=FirstBuild /api-gateway/static/ /static/
COPY --from=FirstBuild /api-gateway/.env /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 3000
CMD [ "/build" ]