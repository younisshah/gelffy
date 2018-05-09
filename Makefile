build-bin:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o gelffy-linux-amd64
	rm gelffy-linux-amd64

build:
	docker build -t hisholiness/gelffy:latest .

run:
	docker run \
	-p 12123:12123/udp \
	--name gelffy \
	--rm \
	hisholiness/gelffy:latest