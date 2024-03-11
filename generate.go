package main

import (
	_ "github.com/ogen-go/ogen"
)

// Generate code for the unmodified OpenAI API.
//go:generate go run github.com/ogen-go/ogen/cmd/ogen -v --clean --target pkg/ogen/generated/openai openapi/openai/openapi.yaml
