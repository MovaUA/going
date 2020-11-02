#!/bin/sh

set -ex

DIR=$(dirname $0)

kubectl create secret tls tls-secret-http --key "${DIR}/server.key" --cert "${DIR}/chain.crt"

