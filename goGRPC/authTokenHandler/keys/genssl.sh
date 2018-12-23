#!/bin/bash

mkdir keys && cd keys && openssl genrsa -aes256 -out private_key.pem 2048 && openssl rsa -pubout -in private_key.pem -out public_key.pem && cd ..