<h1 align="center">
DesignSphere Backend
</h1>

Backend server for DesignSphere written in Go. Handles storing and deployment for application infrastructure resources.

# Build & Run

## Docker Container

### Toolchain

- Require Docker
- Install Docker from [here](https://docs.docker.com/get-docker)

Run the following script to build and run the server in the container.

```
./docker_run.sh start
```

Run the following script to stop the server container or you can do it manually.

```
./docker_run.sh stop
```

## Locally

### Toolchain

- Require Go toolchain
- Install Go from [here](https://go.dev/doc/install)

Run the following script to build the server.

```
./app_build.sh
```

Run the following script to run the server.

```
./app_run.sh
```

---

#### Application will be up at [`http://localhost:8080/app`](http://localhost:8080/app/)

---
