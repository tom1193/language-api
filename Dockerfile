FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/tom1193/language-api

# Get google language api creds
COPY creds/gcp_nlp_creds.json /creds/
ENV GOOGLE_APPLICATION_CREDENTIALS /creds/gcp_nlp_creds.json

RUN cd /go/src/github.com/tom1193/language-api && go get ./
RUN go install github.com/tom1193/language-api

ENTRYPOINT /go/bin/language-api

# Document that the service listens on port 8080.
EXPOSE 8080