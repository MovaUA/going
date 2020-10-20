#!/bin/sh

set -ex

DOMAIN='localhost.test'

DIR=$(dirname $0)

# Generate the CA Key and Certificate:
# 
openssl req -x509 -sha256 -newkey rsa:4096 -keyout "${DIR}/ca.key" -out "${DIR}/ca.crt" -days 3650 -nodes -subj '/CN=GRPC Cert Authority'

# Generate the Server Key, and Certificate and Sign with the CA Certificate:
# 
openssl req -new -newkey rsa:4096 -keyout "${DIR}/server.key" -out "${DIR}/server.csr" -nodes -subj '/CN='"${DOMAIN}"
openssl x509 -req -sha256 -days 365 -in "${DIR}/server.csr" -CA "${DIR}/ca.crt" -CAkey "${DIR}/ca.key" -set_serial 01 -out "${DIR}/server.crt"