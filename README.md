# Encryption Service
[![CI](https://github.com/vpnbeast/encryption-service/workflows/CI/badge.svg?event=push)](https://github.com/vpnbeast/encryption-service/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/vpnbeast/encryption-service)](https://hub.docker.com/r/vpnbeast/encryption-service/)
[![Go Report Card](https://goreportcard.com/badge/github.com/vpnbeast/encryption-service)](https://goreportcard.com/report/github.com/vpnbeast/encryption-service)

This is a web server created for encrypting strings based on [AES-256](https://www.solarwindsmsp.com/blog/aes-256-encryption-algorithm) 
and returns the encrypted response as `JSON`. This service also has capability of checking the clear and encrypted strings, 
takes clear and encrypted strings as JSON and returns a JSON response to check if both strings are compatible with each other.
