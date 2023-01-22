VERSION := 1.0

test:
	go test ./... -count=1

generate-private-key:
	openssl genpkey -algorithm RSA -out key/private.pem -pkeyopt rsa_keygen_bits:2048

generate-public-key:
	openssl rsa -pubout -in key/private.pem -out key/public.pem

start: build dermaservice-up

build:
	docker build -f zarf/docker/Dockerfile.service -t dermacare-service:$(VERSION) .

dermaservice-up: 
	docker-compose -f zarf/docker/docker-compose.yaml --env-file zarf/docker/.env.dev up

dermaservice-stop: 
	docker-compose -f zarf/docker/docker-compose.yaml stop


