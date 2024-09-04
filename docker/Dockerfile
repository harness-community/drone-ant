# Use Alpine Linux as the base image for the final stage
FROM alpine:latest

# Install Ant using apk
RUN apk update && \
    apk add --no-cache openjdk8 apache-ant

# Copy the built Go executable from the builder stage to the final stage
COPY release/linux/amd64/drone-ant-plugin /bin/

# Set the entrypoint to the Go executable
ENTRYPOINT ["/bin/drone-ant-plugin"]
