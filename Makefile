all: \
	client \
	server \
	migrate \
	droptable

client: cmd/client/main.go
	go build -o build/client cmd/client/main.go

server: cmd/server/main.go
	go build -o build/server cmd/server/main.go

migrate: cmd/migrate/main.go
	go build -o build/migrate cmd/migrate/main.go

droptable: cmd/droptable/main.go
	go build -o build/droptable cmd/droptable/main.go
