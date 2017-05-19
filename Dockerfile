FROM scratch
ADD token-svc /
ENTRYPOINT ["/./token-svc"]
