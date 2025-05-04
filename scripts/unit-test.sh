#!/bin/bash

set -euo pipefail

go_to_project="cd src"

eval "$go_to_project"
for module in "$@"; do
    echo "Running tests in $module"
    go test $module -timeout 30s
done