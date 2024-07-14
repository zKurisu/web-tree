clean:
	rm ~/.local/bin/web-tree

install:
	go build .
	cp ./web-tree ~/.local/bin/web-tree
