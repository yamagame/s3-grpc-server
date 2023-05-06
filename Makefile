all: \
	client \
	server \
	fxclient \
	fxserver \
	migrate \
	droptable

client: cmd/client/main.go
	go build -o build/client cmd/client/main.go

server: cmd/server/main.go
	go build -o build/server cmd/server/main.go

fxclient: cmd/fxclient/main.go
	go build -o build/fxclient cmd/fxclient/main.go

fxserver: cmd/fxserver/main.go
	go build -o build/fxserver cmd/fxserver/main.go

migrate: cmd/migrate/main.go
	go build -o build/migrate cmd/migrate/main.go

droptable: cmd/droptable/main.go
	go build -o build/droptable cmd/droptable/main.go
