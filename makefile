format:
	gofmt -w .

run-redis:
	docker run --name redis -p 6379:6379 -d redis:latest

redis-write-json:
	go run ./cases/rw-json/redis/write.go

run-mongodb:
	docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-community-server:latest

mongodb-write-json:
	go run ./cases/rw-json/mongodb/write.go