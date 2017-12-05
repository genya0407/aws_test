docker run -e POSTGRES_USER=aws -e POSTGRES_PASSWORD=aws -p 5432:5432 --name aws --volume pgsql-data:/var/lib/pgsql/data -d postgres:9.6-alpine
