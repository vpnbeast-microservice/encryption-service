# Encryption Service
[![CI](https://github.com/vpnbeast/encryption-service/workflows/CI/badge.svg?event=push)](https://github.com/vpnbeast/encryption-service/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/vpnbeast/encryption-service)](https://hub.docker.com/r/vpnbeast/encryption-service/)
[![Go Report Card](https://goreportcard.com/badge/github.com/vpnbeast/encryption-service)](https://goreportcard.com/report/github.com/vpnbeast/encryption-service)
[![codecov](https://codecov.io/gh/vpnbeast/encryption-service/branch/master/graph/badge.svg)](https://codecov.io/gh/vpnbeast/encryption-service)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=vpnbeast_encryption-service&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=vpnbeast_encryption-service)
[![Go version](https://img.shields.io/github/go-mod/go-version/vpnbeast/encryption-service)](https://github.com/vpnbeast/encryption-service)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is a web server created for encrypting strings based on [AES-256](https://www.solarwindsmsp.com/blog/aes-256-encryption-algorithm)
and returns the encrypted response as `JSON`. This service also has capability of checking the clear and encrypted strings,
takes clear and encrypted strings as JSON and returns a JSON response to check if both strings are compatible with each other.

## Prerequisites
encryption-service requires [vpnbeast/config-service](https://github.com/vpnbeast/config-service) to fetch configuration. Configurations
are stored at [vpnbeast/properties](https://github.com/vpnbeast/properties).

## Configuration
This project fetches the configuration from [config-service](https://github.com/vpnbeast/config-service).
But you can still override them with environment variables:
```
SERVER_PORT             Port number of web server
METRICS_PORT            Port number to expose Prometheus metrics
METRICS_ENDPOINT        Endpoint to expose Prometheus metrics
WRITE_TIMEOUT_SECONDS   Write timeout seconds of the both web server and metrics server
READ_TIMEOUT_SECONDS    Read timeout seconds of the both web server and metrics server
```

## Development
This project requires below tools while developing:
- [Golang 1.16](https://golang.org/doc/go1.16)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)

After installed [pre-commit](https://pre-commit.com/), make sure that you've completed the below final installation steps:
- Make sure that you've installed [pre-commit](https://pre-commit.com/) for our git repository in root directory of the project:
  ```shell
  $ pre-commit install
  ```
- Add below custom variables to `.git/hooks/pre-commit` in the root of our git repository:
  ```python
  # custom variable definition for local development
  os.environ["ACTIVE_PROFILE"] = "unit-test"
  os.environ["CONFIG_PATH"] = "./../../"
  ```

## License
Apache License 2.0
