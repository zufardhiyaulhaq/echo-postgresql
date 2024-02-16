# Echo Postgresql
Echo server that connect to Postgresql server to test connectivity.

### Usage
1. Run
```
source .env.example
make postgresql.up
make run
```

2. test curl
```
curl http://localhost:80/postgresql/test
```
