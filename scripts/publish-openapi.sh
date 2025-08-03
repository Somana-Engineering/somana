#!/bin/bash

# Publish OpenAPI specification to GitHub releases
# Usage: ./scripts/publish-openapi.sh [version]

set -e

VERSION=${1:-$(git describe --tags --abbrev=0 2>/dev/null || echo "v1.0.0")}
SPEC_FILE="api/openapi.yaml"
RELEASE_NAME="Somana API OpenAPI Specification $VERSION"

echo "Publishing OpenAPI specification version $VERSION..."

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo "Error: GitHub CLI (gh) is not installed."
    echo "Please install it from: https://cli.github.com/"
    exit 1
fi

# Check if we're authenticated
if ! gh auth status &> /dev/null; then
    echo "Error: Not authenticated with GitHub CLI."
    echo "Please run: gh auth login"
    exit 1
fi

# Check if spec file exists
if [ ! -f "$SPEC_FILE" ]; then
    echo "Error: OpenAPI specification file not found: $SPEC_FILE"
    exit 1
fi

# Create release with OpenAPI spec
gh release create "$VERSION" \
    "$SPEC_FILE" \
    --title "$RELEASE_NAME" \
    --notes "OpenAPI 3.0 specification for the Somana API.

## Usage

Download the specification and use it with your preferred OpenAPI code generator:

### Using curl:
\`\`\`bash
curl -O https://github.com/lukewing/somana/releases/download/$VERSION/openapi.yaml
\`\`\`

### Using wget:
\`\`\`bash
wget https://github.com/lukewing/somana/releases/download/$VERSION/openapi.yaml
\`\`\`

### In other projects:
- **Go**: Use with \`oapi-codegen\`
- **Python**: Use with \`openapi-generator-cli\`
- **JavaScript**: Use with \`openapi-typescript-codegen\`
- **Java**: Use with \`openapi-generator-cli\`

## API Endpoints

- Health Check: \`GET /health\`
- Hosts: \`GET /api/v1/hosts\`
- Host Management: \`GET/POST/PUT/DELETE /api/v1/hosts/{id}\`
- Heartbeat: \`POST /api/v1/hosts/{id}/heartbeat\`

## Authentication

This API uses API key authentication via the \`Authorization\` header."

echo "✅ OpenAPI specification published as release $VERSION"
echo "📄 Download URL: https://github.com/lukewing/somana/releases/download/$VERSION/openapi.yaml" 