#!/bin/bash

# Install revive
echo "Installing revive..."
go get -u github.com/mgechev/revive

# Install golangci-lint
echo "Installing golangci-lint..."
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

# Install pre-commit hook
echo "Installing pre-commit hook..."
cp hooks/pre-commit .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit

echo "Dependencies installation completed."
