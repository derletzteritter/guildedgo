package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	nethttp "net/http"
	"time"
)

type apiError struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Meta    json.RawMessage `json:"meta"`
}

var (
	ErrMarshal    = errors.New("failed to marshal data")
	ErrNewRequest = errors.New("failed to create new request")
	ErrTimeout    = errors.New("request timed out")
)

func (r *Client) PerformRequest(method, url string, data any) (io.ReadCloser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	marshalData, err := json.Marshal(data)
	if err != nil {
		return nil, ErrMarshal
	}

	var buf io.Reader
	if data == nil {
		buf = nil
	} else {
		buf = bytes.NewBuffer(marshalData)
	}

	req, err := nethttp.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, ErrNewRequest
	}

	req.Header.Set("Authorization", "Bearer "+r.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &nethttp.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, ErrTimeout
		}

		return nil, err
	}

	switch resp.StatusCode {
	case nethttp.StatusOK, nethttp.StatusCreated:
		return resp.Body, nil
	default:
		var apiErr apiError
		json.NewDecoder(resp.Body).Decode(&apiErr)

		var message string
		if apiErr.Message != "" {
			message = apiErr.Message
		} else {
			message = "no message"
		}

		return nil, fmt.Errorf("api error (%d, %s): %s. meta: %v", resp.StatusCode, apiErr.Code, message, apiErr.Meta)
	}
}

func (r *Client) Decode(data io.ReadCloser, v any) error {
	return json.NewDecoder(data).Decode(v)
}
