all:
	gcc main.c -o main

docker:
	docker build . -t hello-docker

##################

export:
	mkdir -p rootfs
	docker export `docker create hello-docker` | tar -C rootfs -xvf -

save:
	mkdir -p layered
	docker save hello-docker | tar -C layered -xvf -

##################

scratch:
	docker build -f DockerfileScratch . -t hello-docker-scratch

##################

staticsimple:
	gcc -Os -static -o main-static-simple main.c

static: 
	gcc -Os -fdata-sections -ffunction-sections -s -Wl,--gc-sections -static -o main-static main.c

strip: static
	strip --strip-all --remove-section=.comment main-static

dockertar: strip
	mkdir -p static
	cp main-static static
	tar -C static -c . | docker import - hellotar