package btcpayserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) request(method, path string, query url.Values, payload any, dest any) error {
	u, errJoinPath := url.JoinPath(c.server, path)
	if errJoinPath != nil {
		return fmt.Errorf("failed to join path: %w", errJoinPath)
	}

	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	var body io.Reader = http.NoBody
	if payload != nil {
		b, errMarshal := json.Marshal(&payload)
		if errMarshal != nil {
			return fmt.Errorf("failed to marshal payload: %w", errMarshal)
		}
		body = bytes.NewReader(b)
	}

	req, errReq := http.NewRequest(method, u, body)
	if errReq != nil {
		return fmt.Errorf("failed to create request: %w", errReq)
	}
	req.Header.Add("Authorization", "token "+c.token)

	if payload != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	res, errDo := c.httpClient.Do(req)
	if errDo != nil {
		return fmt.Errorf("failed to do request: %w", errDo)
	}

	defer res.Body.Close()

	respBody, errReadAll := io.ReadAll(res.Body)
	if errReadAll != nil {
		return fmt.Errorf("failed to read response body: %w", errReadAll)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%d: %s", res.StatusCode, string(respBody))
	}

	if dest == nil {
		return nil
	}

	errDecode := json.Unmarshal(respBody, dest)
	if errDecode != nil {
		return fmt.Errorf("failed to decode response: %w", errDecode)
	}

	return nil
}
