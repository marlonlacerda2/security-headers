package main

import (
	"fmt"
	"net/http"
)

func checkHeaders(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer a requisição: %v\n", err)
		return
	}
	defer resp.Body.Close()

	securityHeaders := []string{
		"Content-Security-Policy",
		"Strict-Transport-Security",
		"X-Content-Type-Options",
		"X-Frame-Options",
		"X-XSS-Protection",
		"Referrer-Policy",
		"Permissions-Policy",
	}

	fmt.Printf("Verificando headers de segurança para %s:\n", url)
	for _, header := range securityHeaders {
		if value := resp.Header.Get(header); value != "" {
			fmt.Printf("%s: %s\n", header, value)
		} else {
			fmt.Printf("%s: Não presente\n", header)
		}
	}
}

func main() {
	// URL da aplicação a ser testada
	url := "https://studiovisual.com.br"
	checkHeaders(url)
}
