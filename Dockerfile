
#BUILDER
FROM golang as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./tmp

RUN cd ./tmp && go build -v -o /app/valid-password

#RUNNER
FROM ubuntu as runner

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=America/Sao_Paulo

WORKDIR /app

COPY --from=builder /app/valid-password /app/valid-password

ENTRYPOINT [ "./valid-password" ]