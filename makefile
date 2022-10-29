generate-private-key:
	openssl genpkey -algorithm RSA -out key/private.pem -pkeyopt rsa_keygen_bits:2048

generate-public-key:
	openssl rsa -pubout -in key/private.pem -out key/public.pem