# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

MAINTAINER DevNet Cloudy Team

LABEL Description="DevNet thousandeyes-operator image"

WORKDIR /
COPY ./bin/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]