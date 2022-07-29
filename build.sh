#!/bin/bash
VERSION=$(git rev-parse --short HEAD)
docker build . -t hughonfinality/onf-test-cli:rev-$VERSION
docker push hughonfinality/onf-test-cli:rev-$VERSION