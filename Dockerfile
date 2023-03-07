ARG GOLANG_VERSION=1.20

### GOLANG BUILD
FROM golang:${GOLANG_VERSION}-alpine as build

RUN apk --no-cache add ca-certificates
RUN apk --no-cache add \
    git \
    git-lfs \
    ;

WORKDIR /go/src/social

COPY . .

ENV CGO_ENABLED 0

RUN go mod download && go mod verify && go install -v ./... ;

### GOLANG DEV
FROM build as dev

CMD ["sleep", "86400"]

### GOLANG PROD
FROM alpine:latest AS prod

WORKDIR /go/src/social

RUN apk add --no-cache tzdata;
###Added user 1000uid/guid and use it###
ENV USER=app
ENV UID=1000
ENV GID=1000
###Added user sensei and use it###
RUN addgroup --gid "$GID" "$USER"
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)"\
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"
USER $USER

COPY --from=build --chown=$USER:$USER /go/src/social/configs /go/src/social/configs

### API PROD
FROM prod AS api_prod
COPY --from=build --chown=$USER:$USER /go/bin/api /go/bin/api
EXPOSE 9000
CMD ["/go/bin/api"]
