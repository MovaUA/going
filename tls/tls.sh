#!/bin/sh

set -ex

DIR=$(dirname $0)

kubectl create secret tls tls-secret --key "${DIR}/server.key" --cert "${DIR}/chain.crt" -n ingress-nginx

