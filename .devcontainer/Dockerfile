FROM registry.coregame.de/core/core-event-0:latest as devimage

FROM golang:latest

COPY --from=devimage /core/con_lib.h /core/con_lib.h
COPY --from=devimage /core/con_lib.a /core/con_lib.a
