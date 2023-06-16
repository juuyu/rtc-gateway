// Package srs
// @title
// @description
// @author njy
// @since 2023/6/1 15:50
package srs

import "fmt"

func genConf() []string {
	return []string{
		"",
		"# For SSL/TLS config.",
		"listen       443 ssl http2 default_server;",
		"listen       [::]:443 ssl http2 default_server;",
		"ssl_session_cache shared:SSL:1m;",
		"ssl_session_timeout  10m;",
		"ssl_ciphers HIGH:!aNULL:!MD5;",
		"ssl_prefer_server_ciphers on;",
		fmt.Sprintf(`ssl_certificate "%v/containers/ssl/nginx.crt";`, ""),
		fmt.Sprintf(`ssl_certificate_key "%v/containers/ssl/nginx.key";`, ""),
		"",
		"# For automatic HTTPS.",
		"location /.well-known/acme-challenge/ {",
		"  proxy_pass http://127.0.0.1:2022$request_uri;",
		"}",
	}
}
