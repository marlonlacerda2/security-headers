package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
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
	var url string

	var rootCmd = &cobra.Command{
		Use:   "security-checker",
		Short: "Security Checker verifica os headers de segurança de uma aplicação web.",
		Run: func(cmd *cobra.Command, args []string) {
			if url == "" {
				fmt.Println("Você deve fornecer uma URL. Use o comando 'security-checker --url <URL>'")
				return
			}
			checkHeaders(url)
		},
	}

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL da aplicação a ser testada")

	rootCmd.Execute()
}
