# Security Hardening & Compliance Guide

Production security settings for networking, identity validation, and logging.

## 1. Network Hardening (mTLS)

To restrict endpoints to verified inter-service callers, configure mutual TLS (mTLS) inside your Docker compose environment:

1. Generate client and server certificates via `Pranor Mesh` root CA:
   ```bash
   curl -X POST http://localhost:8089/api/csr -d '{"service":"my-backend", "csr":"..."}'
   ```
2. Enable mTLS in service config files:
   ```yaml
   security:
     mtls_enabled: true
     root_ca_path: "/certs/ca.pem"
     client_cert_path: "/certs/cert.pem"
     client_key_path: "/certs/key.pem"
   ```

## 2. JWT Signature Verification

* Always verify that `PRANOR_JWT_SECRET` is at least 32 cryptographically random bytes.
* Do not expose `/readyz` or `/healthz` endpoints to public IP ranges; restrict ingress routing in `Pranor Gate`.

## 3. Log Redaction

The regex-based log sanitizer in `Pranor Core/pkg/middleware/log.go` automatically redacts sensitive tokens:
```go
// Output is scrubbed automatically:
log.info("Processing login request with password: " + req.Password)
// Output: [INFO] Processing login request with password: [REDACTED]
```
Ensure all custom handlers route logs through `Pranor Core.SanitizeLog(msg)` before emission.
