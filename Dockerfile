FROM public.ecr.aws/docker/library/alpine:3.18.4

RUN mkdir /app

COPY go-binary /app/go-binary

WORKDIR /app

EXPOSE 8080

USER 65534:65534

ENTRYPOINT ["/app/go-binary"]
