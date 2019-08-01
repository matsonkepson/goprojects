# genarate keys for the https server

run locally 

```
$ openssl genrsa -out server.key 2048
$ openssl ecparam -genkey -name secp384r1 -out server.key
$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
to access tls webserver use 
> https://localhost:8081