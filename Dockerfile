FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# Expose port
EXPOSE ${HTTP_PORT}

CMD [ "tail", "-f", "/dev/null" ]