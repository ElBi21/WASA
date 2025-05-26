test_api: doc/api.yaml
	lint-openapi doc/api.yaml -r spectral.js

build:
	./open-node.sh | {yarn run build-embed && exit} \
    go build -tags webui ./cmd/webapi/

run:
	./open-node.sh \
    yarn run dev