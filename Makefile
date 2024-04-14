run:
	go run API/cmd/main.go
dc_up:
	docker-compose up -d
	docker-compose down db_test
dc_down:
	docker-compose down
dc_build:
	docker-compose build
test:
	go test ./API/tests/banner_test.go
	go test ./API/tests/bannerGet_test.go
	go test ./API/tests/error_json_test.go
	go test ./API/tests/params_test.go
integration_test_user_banner:
	docker-compose up -d db_test

	go test ./API/tests/int_user_banner_test.go

	docker-compose down db_test

integration_test_get_banner:
	docker-compose up -d db_test

	go test ./API/tests/int_get_banner_test.go

	docker-compose down db_test
ping_db:
	docker-compose up -d db_test
	go test ./API/tests/ping_bd_test.go
	docker-compose down db_test
