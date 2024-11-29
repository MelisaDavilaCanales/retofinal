package models

import "github.com/golang-jwt/jwt/v5"

type Payload struct {
	jwt.MapClaims        // ExpiryAt, IssueAt
	Session       string `json:"session"`
	SearchFilters string `json:"search_filters"`
}
