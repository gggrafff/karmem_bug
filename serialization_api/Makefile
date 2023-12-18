all: build

build: api_generated.go

api_generated.go: serialization_api.km
	go run karmem.org/cmd/karmem build --golang -o "." serialization_api.km