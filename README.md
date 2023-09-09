# README
## How to Run

### Build and Run Docker Image

1. Build the Docker image (local):

```bash
docker build -t testing/welcome:local .
```
2. Run the Docker image (local port 8000):

```bash
docker run -p 8000:5000 testing/welcome:local
```
3. Run the application using docker-compose (using ghcr and local tag):
- for prod (using ghcr)
```bash
docker-compose up prod
```
- for local
```bash
docker-compose up local
```

**You can access the application at http://localhost:8000**