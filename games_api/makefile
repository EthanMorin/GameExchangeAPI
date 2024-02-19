codegen:
	oapi-codegen \
	-generate gin,spec \
	-package api config/games_api.yaml > api/api.gen.go
	oapi-codegen \
	-generate types \
	-package models config/games_api.yaml > models/models.gen.go
	go mod tidy
