#!/bin/bash
go mod vendor
gcloud functions deploy roll --runtime go111 --trigger-http --memory 128MB
rm -rf vendor