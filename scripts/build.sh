echo "Building the Docker image..."
docker-compose -f deployments/docker/docker-compose.yml up --build -d 
echo "Done."