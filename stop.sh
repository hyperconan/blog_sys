#!/bin/bash

# Blog Microservices Stop Script

echo "Stopping Blog Microservices System..."

# Kill all blog services
pkill -f "user-server" 2>/dev/null && echo "User service stopped"
pkill -f "blog-server" 2>/dev/null && echo "Blog service stopped"
pkill -f "comment-server" 2>/dev/null && echo "Comment service stopped"
pkill -f "gateway-server" 2>/dev/null && echo "Gateway service stopped"

echo "All services stopped"