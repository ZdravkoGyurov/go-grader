start-mongo:
	docker run -d -p 27017:27017 mongo

start-app:
	go run cmd/go-grader/main.go

start-app-in-docker:
	docker build -t go-grader . && docker run --net gograder-mongo-cluster -it -d -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock go-grader