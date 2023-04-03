FROM alpine
ADD pod /pod
COPY config /root/.kube/config
ENTRYPOINT [ "/pod" ]