#!/bin/sh

set -ex

DOMAIN='localhost.test'

DIR=$(dirname $0)

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout "${DIR}/tls.key" -out "${DIR}/tls.crt" -subj "/CN=${DOMAIN}/O=${DOMAIN}"

kubectl create secret tls tls-secret --key "${DIR}/tls.key" --cert "${DIR}/tls.crt"

