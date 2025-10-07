#!/bin/bash

# GAuth Educational Web Interface Startup Script
# ⚠️ Educational Implementation - Not for Production Use

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🎓 GAuth Educational Web Interface${NC}"
echo -e "${YELLOW}⚠️  EDUCATIONAL IMPLEMENTATION ONLY${NC}"
echo -e "${BLUE}📚 RFC-0150 Go Learning Environment${NC}"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed or not in PATH${NC}"
    echo "Please install Go 1.21 or higher to run the educational demo"
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | grep -oE 'go[0-9]+\.[0-9]+' | grep -oE '[0-9]+\.[0-9]+')
echo -e "${GREEN}✓ Go version: ${GO_VERSION}${NC}"

# Ensure we're in the right directory
if [ ! -f "go.mod" ]; then
    echo -e "${RED}❌ Please run this script from the project root directory${NC}"
    exit 1
fi

# Check if web server exists, build if not
if [ ! -f "web-server" ]; then
    echo -e "${BLUE}🔨 Building educational web server...${NC}"
    go build -o web-server web/server.go
    echo -e "${GREEN}✓ Web server built successfully${NC}"
else
    echo -e "${GREEN}✓ Web server binary found${NC}"
fi

# Default port
PORT=${1:-8080}

echo ""
echo -e "${BLUE}🚀 Starting Educational Demo Server...${NC}"
echo -e "${GREEN}🌐 Server will start on: http://localhost:${PORT}${NC}"
echo -e "${GREEN}📖 Documentation: http://localhost:${PORT}/docs/${NC}"
echo -e "${GREEN}🔧 Health Check: http://localhost:${PORT}/api/v1/educational/health${NC}"
echo ""
echo -e "${YELLOW}⚠️  Educational Notice:${NC}"
echo "   This is a learning implementation for RFC-0150 concepts"
echo "   NOT intended for production use or real security applications"
echo ""
echo -e "${BLUE}Press Ctrl+C to stop the educational demo server${NC}"
echo ""

# Start the server
./web-server "$PORT"