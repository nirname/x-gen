all: config

# compile
build:
	@docker build . -t x-gen >&2

start:
	sudo nginx -c "$(CURDIR)/nginx.conf"

# generate nginx config
config: build
	@docker run --rm -e SERVICES_PATH=$(SERVICES_PATH) -e LOCATIONS_PATH=$(LOCATIONS_PATH) -v "`pwd`:/app" x-gen

.PHONY: clean

clean:
	docker images -f "dangling=true" -q | xargs docker rmi -f
	docker images x-gen -q | xargs docker rmi -f

todo: /tmp/todo.html
	open /tmp/todo.html

/tmp/todo.html: todo.md
	pandoc -f markdown -t html todo.md > /tmp/todo.html
