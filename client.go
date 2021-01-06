package godaddygo

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/oze4/godaddygo/internal/exception"
)

// makeDo makes an http.Request and sends it
func makeDo(ctx context.Context, config *Config, method, path string, body io.Reader, expectStatus int) ([]byte, error) {
	version := config.version.String()
	if version == "" {
		return nil, exception.InvalidValue("version value not allowed")
	}

	urlBase := config.baseURL.String()
	if urlBase == "" {
		return nil, exception.InvalidValue("urlBase value not allowed")
	}

	fullURL := urlBase + "/" + version + path

	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "sso-key "+config.key+":"+config.secret)
	req.Header.Set("Content-Type", "application/json")

	reqWithCtx := req.WithContext(ctx)

	resp, err := config.client.Do(reqWithCtx)
	if err != nil {
		return nil, exception.SendingRequest(err)
	}
	if resp.StatusCode != expectStatus {
		strerr, err := copyAndCloseBody(resp.Body)
		if err != nil {
			strerr = []byte(err.Error())
		}
		return nil, exception.InvalidStatusCode(expectStatus, resp.StatusCode, string(strerr))
	}
	return copyAndCloseBody(resp.Body)
}

func copyAndCloseBody(r io.ReadCloser) ([]byte, error) {
	response, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	r.Close()
	return response, nil
}
