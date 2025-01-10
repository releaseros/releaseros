FROM alpine:3.21.2@sha256:56fa17d2a7e7f168a043a2712e63aed1f8543aeafdcee47c58dcffe38ed51099

RUN apk add --no-cache git

COPY releaseros /usr/local/bin/releaseros

ENTRYPOINT ["releaseros"]
