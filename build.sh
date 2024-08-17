set -xe

buf lint
rm -rf api
buf generate