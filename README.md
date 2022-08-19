
```
docker run -d \
	--name postgres \
    -p 5432:5432 \
	-e POSTGRES_PASSWORD=mysecretpassword \
	-v postgres-data:/var/lib/postgresql/data \
	postgres:14.5-alpine
```