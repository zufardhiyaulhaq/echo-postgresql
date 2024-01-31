# Echo Postgresql
Echo server that connect to Postgresql server to test connectivity.

### Usage
1. Run
```
source .env.example
make run
```

2. Telnet & use echo
```
telnet localhost 5000
Trying ::1...
Connected to localhost.
Escape character is '^]'.

echo
echo

curl http://localhost:80/postgresql/test
```
