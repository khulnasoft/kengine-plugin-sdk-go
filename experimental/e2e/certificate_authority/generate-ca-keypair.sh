#!/bin/bash
set -ex
openssl genrsa -aes256 -passout pass:1234 -out kengine-e2e-ca.key.pem 4096
openssl rsa -passin pass:1234 -in kengine-e2e-ca.key.pem -out kengine-e2e-ca.key.pem
openssl req -config openssl.cnf -key kengine-e2e-ca.key.pem -new -x509 -days 3650 -sha256 -extensions v3_ca -out kengine-e2e-ca.pem