docker run --rm -e POSTGRES_USER=aws -e POSTGRES_PASSWORD=aws -p 5432:5432 --name aws --volume pgsql-data:/var/lib/pgsql -d postgres:9.6-alpine
