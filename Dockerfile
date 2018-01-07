FROM alpine
COPY gopath/bin/spinnaker-demo-hub /go/bin/spinnaker-demo-hub
COPY static /static
ENTRYPOINT /go/bin/spinnaker-demo-hub