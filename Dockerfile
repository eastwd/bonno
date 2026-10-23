FROM golang:1.27-trixie AS build

WORKDIR /bonno

COPY tools/revel-cmd/go.mod tools/revel-cmd/go.sum tools/revel-cmd/tools.go ./tools/revel-cmd/
RUN cd tools/revel-cmd && go install github.com/revel/cmd/revel

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN revel build -a . -t /build -m prod

FROM debian:trixie-slim
LABEL maintainer="bando"

COPY --from=build /build /app
WORKDIR /app/src/bonno

CMD ["/app/run.sh"]
EXPOSE 9000
