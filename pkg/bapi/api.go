package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	APP_KEY    = "aiaa"
	APP_SECRET = "blog-service"
)

type AccessToken struct {
	Token string `json:"token"`
}

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

// 所有 API 请求都需要带上的 AccessToken 的获取
func (a *API) getAccessToken(ctx context.Context) (string, error) {
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APP_KEY, APP_SECRET))
	if err != nil {
		return "", err
	}

	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

// API SDK 统一的 HTTP GET 的请求方法
func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

// GetTagList 获取多个tag
func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, err
	}

	return body, nil
}
