# Dream Back End

## Access Path

```sh
# See blue
http://localhost:8000/hello

# Echo name
http://localhost:8000/hello/{name}

# See Swagger UI
http://localhost:8000/swagger/index.html

```

## CI POC

```sh
# Pull image
docker pull famedocker/go-back:{$ on build}

# Run container
docker run -d --name go-back -p 8000:8000 famedocker/go-back:{$ on build}
```