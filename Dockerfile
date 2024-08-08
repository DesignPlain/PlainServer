# syntax=docker/dockerfile:1

FROM golang:1.22

# Add image label
LABEL image="ds"

# Set destination for COPY
WORKDIR /app

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux ./app_build.sh

EXPOSE 8080

# Run
CMD ["./app_run.sh"]
