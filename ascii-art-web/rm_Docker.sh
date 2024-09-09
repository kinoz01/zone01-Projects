docker stop $(docker ps -aq)  # Stops all running containers
docker rm $(docker ps -aq)    # Removes all containers
docker rmi $(docker images -q)  # Removes all images
docker system prune -a # Removes all unused containers, networks, images
