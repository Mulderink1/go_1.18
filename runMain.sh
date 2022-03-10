#!/bin/bash

echo ""
cd ./fuzz
go1.18beta2 test ./fuzz_test.go

echo ""
go1.18beta2 test -fuzz ./... -fuzztime=3s

# shellcheck disable=SC2103
cd ..
echo ""

# shellcheck disable=SC2028
go1.18beta2 run main.go
echo ""