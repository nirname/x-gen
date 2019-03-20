FROM golang:1.8
WORKDIR /app
COPY *.go docker.makefile ./
RUN make -f docker.makefile

FROM alpine
COPY --from=0 /app/x-gen /bin
ENTRYPOINT x-gen
WORKDIR /app