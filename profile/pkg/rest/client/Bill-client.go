package client

import (
	"io"
	"net/http"
)

func ExecuteNodeA0() ([]byte, error) {
	response, err := http.Get("Bill:8084/ping")

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
