gen-mock:
	mockgen -source letscloud.go -destination httpclient/http_client_mock.go -package httpclient Requester

test:
	go test -cover -run=$TestClient

format:
	go fmt github.com/letscloud-community/letscloud-go/...