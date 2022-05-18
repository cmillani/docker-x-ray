all: docker


docker:
	docker build . -t hello-go

run:
	docker run -p 8899:8899 hello-go 

dockertwo:
	docker build -f DockerfileTwo . -t hello-go-two

runtwo:
	docker run -p 8899:8899 hello-go-two