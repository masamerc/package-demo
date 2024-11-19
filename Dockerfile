ARG GO_VERSION=1.22.6

FROM golang:${GO_VERSION} AS build

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/server .


FROM alpine:latest AS final

# add an unprivileged user for security best practices
# https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#user
ARG UID=1001
RUN adduser \
    --disabled-password \
    --gecos "" \ 
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

COPY --from=build /bin/server /bin/server

ARG PORT=8080
EXPOSE ${PORT}

ENTRYPOINT ["/bin/server"]
