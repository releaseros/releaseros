FROM alpine:3.18.4@sha256:eece025e432126ce23f223450a0326fbebde39cdf496a85d8c016293fc851978

RUN apk add --no-cache git

COPY releaseros /usr/local/bin/releaseros

ENTRYPOINT ["releaseros"]
