FROM alpine:3.21.0@sha256:21dc6063fd678b478f57c0e13f47560d0ea4eeba26dfc947b2a4f81f686b9f45

RUN apk add --no-cache git

COPY releaseros /usr/local/bin/releaseros

ENTRYPOINT ["releaseros"]
