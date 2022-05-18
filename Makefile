clean:
	rm -rf runc/config.json runc/rootfs
	rm -f capture/*txt*
	rm -rf hello-world/layered hello-world/static hello-world/rootfs hello-world/main hello-world/main-static hello-world/main-static-simple
	docker image rm -f hello-go hello-go-two hellotar hello-world hello-docker hello-docker-scratch
