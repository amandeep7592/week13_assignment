# csdd1008_LB

docker swarm init

docker service create --name sdd1008_lb --publish 8080:8080 --replicas 3 sdd1008_lb
