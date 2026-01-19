#!/bin/bash

# Blog Microservices Startup Script

echo "Starting Blog Microservices System..."

# Check if etcd is running
if ! pgrep -x "etcd" > /dev/null; then
    echo "Error: ETCD is not running. Please start ETCD first:"
    echo "  etcd"
    exit 1
fi

echo "ETCD is running ✓"

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Build services
echo "Building services..."
cd "$SCRIPT_DIR/user" && go build -o user-server . || exit 1
cd "$SCRIPT_DIR/blog" && go build -o blog-server . || exit 1
cd "$SCRIPT_DIR/comment" && go build -o comment-server . || exit 1
cd "$SCRIPT_DIR/gateway" && go build -o gateway-server . || exit 1

echo "Services built successfully ✓"

# Start services in background
echo "Starting User Service..."
cd "$SCRIPT_DIR/user" && ./user-server -f etc/user.yaml &
USER_PID=$!

echo "Starting Blog Service..."
cd "$SCRIPT_DIR/blog" && ./blog-server -f etc/blog.yaml &
BLOG_PID=$!

echo "Starting Comment Service..."
cd "$SCRIPT_DIR/comment" && ./comment-server -f etc/comment.yaml &
COMMENT_PID=$!

echo "Starting Gateway Service..."
cd "$SCRIPT_DIR/gateway" && ./gateway-server -f etc/gateway-api.yaml &
GATEWAY_PID=$!

echo ""
echo "Services started successfully!"
echo "User Service PID: $USER_PID"
echo "Blog Service PID: $BLOG_PID"
echo "Comment Service PID: $COMMENT_PID"
echo "Gateway Service PID: $GATEWAY_PID"
echo ""
echo "API Gateway available at: http://localhost:8888"
echo ""
echo "Press Ctrl+C to stop all services"

# Wait for interrupt
trap 'echo ""; echo "Stopping services..."; kill $USER_PID $BLOG_PID $COMMENT_PID $GATEWAY_PID 2>/dev/null; exit' INT

wait