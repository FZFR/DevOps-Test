# README
## How to Run

### Build and Run Docker Image

1. Build the Docker image (local):

```bash
docker build -t localtesting/welcome .
```
2. Run the Docker image (local port 8000):

```bash
docker run -p 8000:5000 localtesting/welcome
```
3. Run the application locally using docker-compose:
```bash
docker-compose up
```

**You can access the application at http://localhost:8000**