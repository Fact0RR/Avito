run:
	go run -C API/ /cmd/main.go
dc_up:
	docker-compose up -d
	docker-compose down db_test
dc_down:
	docker-compose down
dc_build:
	docker-compose build
test:
	go test -C API/ ./tests/banner_test.go
	go test -C API/ ./tests/bannerGet_test.go
	go test -C API/ ./tests/error_json_test.go
	go test -C API/ ./tests/params_test.go
integration_test_user_banner:
	docker-compose up -d db_test
	go test -C API/ ./tests/int_user_banner_test.go
	docker-compose down db_test

integration_test_get_banner:
	docker-compose up -d db_test
	go test -C API/ ./tests/int_get_banner_test.go
	docker-compose down db_test
ping_db:
	docker-compose up -d db_test
	go test -C API/ ./tests/ping_bd_test.go
	docker-compose down db_test

integration_test_delete_banner:
	docker-compose up -d db_test
	go test -C API/ ./tests/int_delete_banner_test.go
	docker-compose down db_test

integration_test_patch_banner:
	docker-compose up -d db_test
	go test -C API/ ./tests/int_patch_banner_test.go
	docker-compose down db_test

integration_test_post_banner:
	docker-compose up -d db_test
	go test -C API/ ./tests/int_post_banner_test.go
	docker-compose down db_test

integration_tests:
	docker-compose up -d db_test
	go test -C API/ ./tests/int_user_banner_test.go
	go test -C API/ ./tests/int_get_banner_test.go
	go test -C API/ ./tests/int_patch_banner_test.go
	go test -C API/ ./tests/int_post_banner_test.go
	go test -C API/ ./tests/int_delete_banner_test.go
	docker-compose down db_test