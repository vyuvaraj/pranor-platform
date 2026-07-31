module github.com/vyuvaraj/pranor-platform

go 1.26.4

require (
	github.com/vyuvaraj/pranor/core v0.0.0
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/vyuvaraj/pranor/core => ../pranor/core

require github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
