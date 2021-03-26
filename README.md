## Run MongoDB in docker:
```bash
docker run -d -p 27017:27017 mongo
```

## Run in docker:
```bash
docker build -t go-grader . && docker run -it -d -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock go-grader
```

## Start mongo replica set
1. docker pull mongo
2. docker network create gograder-mongo-cluster
3. docker run -d --net gograder-mongo-cluster -p 27017:27017 --name mongo1 mongo mongod --replSet gograder-mongo-set --port 27017
4. docker run -d --net gograder-mongo-cluster -p 27018:27018 --name mongo2 mongo mongod --replSet gograder-mongo-set --port 27018
5. docker run -d --net gograder-mongo-cluster -p 27019:27019 --name mongo3 mongo mongod --replSet gograder-mongo-set --port 27019
6. 127.0.0.1       mongo1 mongo2 mongo3
7. docker exec -it mongo1 mongo
8. db = (new Mongo('localhost:27017')).getDB('test')
9. config={"_id":"gograder-mongo-set","members":[{"_id":0,"host":"mongo1:27017"},{"_id":1,"host":"mongo2:27018"},{"_id":2,"host":"mongo3:27019"}]}
10. rs.initiate(config)
11. mongodb://localhost:27017,localhost:27018,localhost:27019/{db}?replicaSet=gograder-mongo-set