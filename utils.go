package main

import (
	"html"
	"net/http"
	"strings"
)

// return sanitized value
func GetSanitizeHeader(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, " ", "", -1)
	return html.EscapeString(str)
}

// get user ip from request
func GetIp(r *http.Request) (ip string) {

	ip = GetSanitizeHeader(r.Header.Get("X-Real-IP"))
	if ip == "" {
		ip = GetSanitizeHeader(r.Header.Get("X-Forwarded-For"))
	}
	if ip == "" {
		ip = GetSanitizeHeader(r.RemoteAddr)
	}
	// if multiple ips, get the first
	ip = strings.Split(ip, ",")[0]
	// extact IP from <ip>:<port> with ipv6 in mind
	splittedIp := strings.Split(ip, ":")
	if len(splittedIp) > 1 {
		ip = strings.Join(splittedIp[:len(splittedIp)-1], ":")
	}
	return
}
