run:
	go run cmd/main.go

tidy:
	go mod tidy

lint:
	golangci-lint run

get-all-tasks:
	curl http://localhost:8080/tasks

get-task-by-id:
	curl http://localhost:8080/tasks/$(id)

create-task:
	curl http://localhost:8080/tasks --include --header "Content-Type: application/json" --request "POST" --data '{"title":"$(title)", "description":"$(description)"}'

update-task:
	curl http://localhost:8080/tasks --include --header "Content-Type: application/json" --request "PUT" --data '{"id":"$(id)", "title":"$(title)", "description":"$(description)"}'

delete-task:
	curl http://localhost:8080/tasks/$(id) --include --header "Content-Type: application/json" --request "DELETE"