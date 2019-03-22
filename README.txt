# localy handle and redirect outlook safelinks

## setup

* in hosts file handle your safelinks domain and point it to 127.0.0.1 - now you should not be able to visit links from outlook

* install golang

* download https://github.com/golang/go/blob/master/src/crypto/tls/generate_cert.go

* `go run generate_cert.go -host=127.0.0.1,localhost,*.safelinks.protection.outlook.com -ca -ecdsa-curve=P384`

* `go build rewrite.go`

* create startup link or task scheduler configuration to automatically start rewrite script

* add pernament exception in your browser for your certificate

Now you skiped one of m$ tracking "feature" and speed up opening links from outlook.