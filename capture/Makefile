dockerd:
	sudo strace -s 80 -ff -o dockerd.txt -p `pgrep dockerd`

containerd:
	sudo strace -ff -o containerd.txt -p `pgrep containerd | head -n 1`

docker:
	sudo strace -s 80 -ff -o docker.txt docker run hello-docker

