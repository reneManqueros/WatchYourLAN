DUSER=aceberg
DNAME=watchyourlan

IFACE=virbr-bw
DBPATH=data/hosts.db
SHOUTRRR_URL=gotify://192.168.2.1:8083/AwQqpAae.rrl5Ob/?title=Unknown host detected&DisableTLS=yes
THEME=darkly

mod:
	cd src && \
	rm go.mod && \
	go mod init $(DNAME) && \
	go mod tidy

run:
	cd src && \
	sudo \
	env IFACE=$(IFACE) DBPATH=$(DBPATH) THEME=$(THEME) \
	go run .

go-build:
	cd src && \
	go build .

docker-build:
	docker build -t $(DUSER)/$(DNAME) .

docker-run:
	docker rm wyl || true
	docker run --name wyl \
	-e "IFACE=$(IFACE)" \
	-e "SHOUTRRR_URL=$(SHOUTRRR_URL)" \
	-e "TZ=Asia/Novosibirsk" \
	--network="host" \
	-v ~/.dockerdata/wyl:/data \
	$(DUSER)/$(DNAME)

clean:
	rm src/$(DNAME) || true
	docker rmi -f $(DUSER)/$(DNAME)

release:
	cd src; curl -sfL https://goreleaser.com/static/run | bash -s -- release

dev: docker-build docker-run