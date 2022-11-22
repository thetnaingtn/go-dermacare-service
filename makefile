VERSION := 1.0

generate-private-key:
	openssl genpkey -algorithm RSA -out key/private.pem -pkeyopt rsa_keygen_bits:2048

generate-public-key:
	openssl rsa -pubout -in key/private.pem -out key/public.pem

dermaservice:
	docker build -f zarf/docker/Dockerfile.service -t dermacare-service:$(VERSION) .

dermaservice-up: dermaservice start

dermaservice-stop: docker-compose -f zarf/docker/docker-compose.yaml stop

start:
	docker-compose -f zarf/docker/docker-compose.yaml up

	
