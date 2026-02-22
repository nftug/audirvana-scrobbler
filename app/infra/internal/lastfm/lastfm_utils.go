package lastfm

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const apiURL = "https://ws.audioscrobbler.com/2.0/"

type lastFMConfig struct {
	apiKey    string
	apiSecret string
}

func (l *lastFMAPIImpl) callPostWithoutSk(ctx context.Context, method string, args map[string]string) (map[string]any, error) {
	params := url.Values{}

	params.Set("method", method)
	for k, v := range args {
		params.Set(k, v)
	}
	params.Set("api_key", l.cfg.apiKey)
	params.Set("api_sig", l.createSignature(params))

	result, err := invokeAPICall(ctx, http.MethodPost, params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (l *lastFMAPIImpl) callPostWithSk(ctx context.Context, method string, args map[string]string) (map[string]any, error) {
	if l.sessionKey == "" {
		return nil, errors.New("not logged in")
	}
	args["sk"] = l.sessionKey
	return l.callPostWithoutSk(ctx, method, args)
}

func (l *lastFMAPIImpl) createSignature(params url.Values) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sigPlain strings.Builder
	for _, k := range keys {
		sigPlain.WriteString(k + params.Get(k))
	}
	sigPlain.WriteString(l.cfg.apiSecret)

	hasher := md5.New()
	hasher.Write([]byte(sigPlain.String()))
	return hex.EncodeToString(hasher.Sum(nil))
}

func invokeAPICall(ctx context.Context, httpMethod string, params url.Values) (map[string]any, error) {
	params.Add("format", "json")

	req, err := http.NewRequestWithContext(ctx, httpMethod, apiURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if errCode, exists := result["error"]; exists {
		return nil, fmt.Errorf(
			"API error: %v, message: %s",
			errCode,
			result["message"],
		)
	}

	return result, nil
}
