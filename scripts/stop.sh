#!/bin/bash

echo "Stopping all services..."
docker-compose -f deployments/docker/docker-compose.yml down
echo "All services stopped."
