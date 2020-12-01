server-up-test:
	podman run --rm --pod new:meili-go \
		--name meilisearch -p 7700:7700 \
		-v ${PWD}/data.ms:/data.ms:Z \
		getmeili/meilisearch

server-up:
	podman run -d --pod new:meili-go \
		--name meilisearch -p 7700:7700 \
		-v ${PWD}/data.ms:/data.ms:Z \
		getmeili/meilisearch

server-down:
	podman pod rm meili-go --force
	rm -rf data.ms/*

migrate-index:
	go run -v cmd/index/main.go

update-uid-documents:
	go run -v cmd/add/main.go

get-uid-documents:
	go run -v cmd/get/main.go

get-deserts:
	go run -v cmd/desert/main.go
