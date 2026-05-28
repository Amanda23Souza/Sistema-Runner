package command

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// defaultHTTPTimeout define o timeout padrão para requisições ao servidor assinador.
const defaultHTTPTimeout = 10 * time.Second

// httpPost realiza um POST JSON com timeout e retorna o body parseado como map.
// Trata explicitamente: timeout, conexão recusada e resposta malformada.
func httpPost(url, body string) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultHTTPTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: defaultHTTPTimeout}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("timeout após %s aguardando resposta do servidor", defaultHTTPTimeout)
		}
		return nil, fmt.Errorf("conexão recusada ou servidor inacessível: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("servidor retornou status HTTP %d (esperado 2xx)", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("resposta do servidor não é JSON válido: %w", err)
	}
	return result, nil
}

// httpGet realiza um GET com timeout e retorna o body parseado como map.
func httpGet(url string) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultHTTPTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	client := &http.Client{Timeout: defaultHTTPTimeout}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("timeout após %s aguardando resposta do servidor", defaultHTTPTimeout)
		}
		return nil, fmt.Errorf("conexão recusada ou servidor inacessível: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("servidor retornou status HTTP %d (esperado 2xx)", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("resposta do servidor não é JSON válido: %w", err)
	}
	return result, nil
}
