# Switchblade

A post exploitation framework based on the framework built in Blackhat-Go

# Setup

Generate certs following this guide:

<code> openssl genrsa -out ca.key 2048 openssl req -new -x509 -days 365 -key ca.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=Acme Root CA" -out ca.crt</code>

<code> openssl req -newkey rsa:2048 -nodes -keyout server.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=_*.example.com" -out server.csr openssl x509 -req -extfile <(printf "subjectAltName=DNS:example.com,DNS:www.example.com") -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt </code>
  
Ensure you specify a proper SANS, otherwise you will encounter errors.

Taken from here https://security.stackexchange.com/questions/74345/provide-subjectaltname-to-openssl-directly-on-the-command-line

To generate implant.pb.go from implant.proto ensure protobuf is installed as well as protoc-gen-go with <code>go get github.com/golang/protobuf/protoc-gen-go</code>
  
Then run <code> protoc -I . implant.proto --go_out=plugins=grpc:./ </code> in the directory of the implant.proto file.
