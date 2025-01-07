package http_client

import "net/http"

func IsFronteggResponseError(resp *http.Response) bool {
	return resp.StatusCode < 200 || resp.StatusCode >= 300
}
