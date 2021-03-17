##Run MongoDB in docker:
```bash
docker run -d -p 27017:27017 mongo
```

##Run in docker:
```bash
docker build -t go-grader . && docker run -it -d -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock go-grader
```