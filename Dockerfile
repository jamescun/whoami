FROM scratch

COPY whoami /

ENTRYPOINT ["/whoami"]
EXPOSE 8080

