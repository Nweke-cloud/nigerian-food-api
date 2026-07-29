 # Operating System Exploration

## What I learned

### 1. Processes

- I observed many running processes on my Ubuntu WSL2 system, including system services, Docker services, Grafana, Prometheus, Redis, Nginx, and container runtime processes.
- The top process by memory usage was **Grafana Server**, using approximately **124 MB** of memory.
- Other high-memory processes included **Prometheus** and **Docker Engine (dockerd)**.
- **Why it matters for containers:**
  Containers are simply isolated Linux processes. They share the host operating system kernel while running independently with their own processes, networking, and filesystem. Docker manages these processes through containerd and the Docker daemon.

---

### 2. Network

Listening ports on my system included:

- 80 (Nginx)
- 3000
- 3001
- 3002
- 3003
- 3004
- 9090
- 9091
- 9092
- 9100
- 9101
- DNS (53)

- **Port 22 (SSH):** Not listening.
- **Why containers need ports exposed:**
  Containers have their own isolated network. Exposing ports maps a port on the host machine to a port inside the container so users and other applications can access services running inside the container.

---

### 3. Environment Variables

- Current PATH: (checked using `env`; PATH is one of the environment variables used by Linux to locate executables.)
- Go version:

  ```text
  go version go1.22.2 linux/amd64
  ```

- Go binary location (from `which go`):

  ```bash
  /usr/bin/go
  ```

- **Why applications need environment variables:**
  Environment variables keeps my code generic, my secrets hidden, and my app adaptable to any environment it runs in.

---

### 4. Storage

- Root filesystem size: **1007 GB**
- Available space: **938 GB**
- Memory:

  - Total RAM: **3.7 GiB**
  - Available RAM: **1.6 GiB**

- **Why this is important for containers:**
  Containers use the host's storage for images, writable layers, and persistent volumes. They also share the host's memory. Monitoring available disk space and RAM is important to ensure containers have enough resources to run reliably.

---

## Summary

This exercise helped me understand that Docker containers are not virtual machines. They are isolated Linux processes that share the host operating system kernel while maintaining separate processes, networking, storage, and environment configurations. This isolation allows applications to run consistently across different environments, helping solve the "works on my machine" problem.
