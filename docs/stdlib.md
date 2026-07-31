# Standard Library

Serv ships with 46 reusable modules in `stdlib/`. Import what you need:

```serv
import { ok, notFound } from "../stdlib/response.pnr"
import { requireAuth } from "../stdlib/auth.pnr"
```

## Quick Reference

### Security
| Module | Key Exports |
|--------|-------------|
| `auth.pnr` | `bearerToken`, `basicAuth`, `requireAuth` |
| `crypto.pnr` | `hashPassword`, `verifyPassword`, `randomToken`, `hmacSign` |
| `jwt.pnr` | `jwtEncode`, `jwtDecode`, `jwtIsExpired` |
| `sanitize.pnr` | `escapeHTML`, `stripTags`, `escapeSQL`, `sanitizeFilename` |
| `ratelimit.pnr` | `createLimiter`, `isAllowed`, `remaining`, `resetLimiter` |
| `mask.pnr` | `maskEmail`, `maskPhone`, `maskCard`, `maskString`, `redact` |
| `ip.pnr` | `extractIP`, `isPrivate`, `isTrustedProxy`, `anonymizeIP` |

### HTTP
| Module | Key Exports |
|--------|-------------|
| `response.pnr` | `ok`, `created`, `badRequest`, `notFound`, `serverError` |
| `pagination.pnr` | `offset`, `pageResponse`, `parsePageParams` |
| `pagination_cursor.pnr` | `encodeCursor`, `decodeCursor`, `cursorResponse` |
| `middleware.pnr` | `corsHeaders`, `requestId`, `logRequest` |
| `http_client.pnr` | `getJSON`, `postJSON`, `isSuccess`, `isClientError` |
| `url.pnr` | `encodeURI`, `parseQuery`, `buildQuery`, `joinPath` |
| `cors.pnr` | `allowOrigin`, `allowAll`, `preflightResponse` |

### Utilities
| Module | Key Exports |
|--------|-------------|
| `datetime.pnr` | `now`, `timestamp`, `isExpired`, `formatDuration` |
| `strings_util.pnr` | `slugify`, `truncate`, `capitalize`, `isEmpty` |
| `math.pnr` | `min`, `max`, `clamp`, `abs`, `percent`, `sum`, `average` |
| `sort.pnr` | `reverse`, `minOf`, `maxOf` |
| `collections.pnr` | `unique`, `flatten`, `chunk`, `first`, `last`, `countWhere` |

### Data
| Module | Key Exports |
|--------|-------------|
| `csv.pnr` | `parseCSV`, `parseRow`, `toCSV` |
| `base64.pnr` | `encode`, `decode`, `isValid` |
| `diff.pnr` | `hasChanged`, `fieldChanged`, `changeRecord` |

### Config
| Module | Key Exports |
|--------|-------------|
| `env.pnr` | `requireEnv`, `envOrDefault`, `envInt`, `envBool` |
| `config.pnr` | `getConfig`, `requireConfig`, `configBool`, `configList` |
| `feature_flags.pnr` | `enableFlag`, `disableFlag`, `isEnabled`, `toggleFlag` |

### Resilience
| Module | Key Exports |
|--------|-------------|
| `retry.pnr` | `backoffDelay`, `defaultMaxRetries` |
| `circuit_breaker.pnr` | `createBreaker`, `isOpen`, `recordSuccess`, `recordFailure` |
| `timeout.pnr` | `withDeadline`, `isTimedOut`, `remainingTime`, `elapsed` |
| `queue.pnr` | `createQueue`, `enqueue`, `dequeue`, `queueSize` |

### Concurrency
| Module | Key Exports |
|--------|-------------|
| `semaphore.pnr` | `createSemaphore`, `tryAcquire`, `release`, `available` |
| `batch.pnr` | `createBatch`, `addToBatch`, `isBatchFull`, `flushBatch` |

### Processing
| Module | Key Exports |
|--------|-------------|
| `job.pnr` | `createJob`, `startJob`, `completeJob`, `failJob` |
| `scheduler.pnr` | `scheduleAfter`, `isScheduled`, `cancelSchedule` |

### Reliability
| Module | Key Exports |
|--------|-------------|
| `idempotency.pnr` | `checkIdempotency`, `markProcessed`, `isProcessed` |
| `dlq.pnr` | `createDLQ`, `sendToDLQ`, `dlqSize`, `clearDLQ` |

### Integration
| Module | Key Exports |
|--------|-------------|
| `webhook.pnr` | `buildPayload`, `sendWebhook`, `verifySignature` |
| `events.pnr` | `on`, `emit`, `hasHandler` |

### Observability
| Module | Key Exports |
|--------|-------------|
| `metrics.pnr` | `counter`, `gauge`, `recordLatency`, `trackRequest` |
| `tracing.pnr` | `traceId`, `startSpan`, `endSpan`, `traceContext` |

### Multi-tenancy
| Module | Key Exports |
|--------|-------------|
| `tenant.pnr` | `extractTenant`, `tenantConfig`, `isTenantActive`, `tenantFilter` |

### Compliance
| Module | Key Exports |
|--------|-------------|
| `audit.pnr` | `auditLog`, `auditAction`, `auditAccess`, `auditAuth`, `auditDenied` |

### Operations
| Module | Key Exports |
|--------|-------------|
| `health.pnr` | `healthy`, `unhealthy`, `degraded`, `buildHealthResponse` |
| `graceful.pnr` | `initShutdown`, `isShuttingDown`, `isDrained` |
| `cache_patterns.pnr` | `cacheKey`, `cacheGet`, `cacheSet`, `invalidate`, `computeIfAbsent` |

### Testing
| Module | Key Exports |
|--------|-------------|
| `testing_helpers.pnr` | `assertEqual`, `assertNotNil`, `assertContains`, `assertTrue` |

## Usage Example

```serv
import { requireAuth, bearerToken } from "../stdlib/auth.pnr"
import { ok, badRequest } from "../stdlib/response.pnr"
import { maskEmail } from "../stdlib/mask.pnr"
import { auditLog } from "../stdlib/audit.pnr"

server "8080"

route "GET" "/api/profile" (req) {
    let authErr = requireAuth(req)
    if authErr != nil { return authErr }

    let token = bearerToken(req)
    auditLog(token, "view", "profile", nil)

    return ok({
        "email": maskEmail("alice@example.com"),
        "role": "admin"
    })
}
```

Full module documentation: see comments at the top of each file in `stdlib/`.
