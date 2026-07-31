# Pranor Notify — Multi-Channel Notification Provider

> **Status:** 🟡 Stable | **Port:** 8094 | **Repository:** [Pranor Notify](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Notify)

## Overview

Pranor Notify is a multi-channel notification provider supporting SMTP email, Slack webhooks, and SMS delivery. It features Go-template rendering, delivery tracking, dead-letter queue retry via Pranor Pulse, and per-channel rate limiting.

## Key Features

- SMTP email delivery with TLS support
- Slack webhook notifications
- SMS delivery via configurable providers
- Go-template rendering for message bodies
- Delivery tracking with status callbacks
- Dead-letter queue retry via Pranor Pulse
- Per-channel rate limiting
- Template management and versioning

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | HTTP listen port | `8094` |
| `SERVMAIL_SMTP_HOST` | SMTP server host | (required for email) |
| `SERVMAIL_SLACK_WEBHOOK` | Slack webhook URL | (required for Slack) |
| `PRANOR_VAULT_URL` | Pranor Vault URL for template storage | (optional) |
| `PRANOR_OTLP_ENDPOINT` | OTel collector URL | (disabled) |

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `POST /api/v1/send` | Send a notification |
| `GET /api/v1/templates` | List available templates |
| `GET /api/v1/tracking/{id}` | Get delivery status |

## Pranor Integration

```srv
mail.send(to, template, data)
notify("slack", msg)
```
