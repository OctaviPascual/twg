package stripe_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type response struct {
	StatusCode int    `json:"status_code"`
	Body       []byte `json:"body"`
}

type recorderClient struct {
	t         *testing.T
	responses []response
}

func (rc *recorderClient) Do(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		rc.t.Fatalf("http request failed. err = %v", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		rc.t.Fatalf("failed to read the response body. err = %v", err)
	}
	rc.responses = append(rc.responses, response{
		StatusCode: res.StatusCode,
		Body:       body,
	})
	res.Body = io.NopCloser(bytes.NewReader(body))
	return res, err
}
