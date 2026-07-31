# Approachable Issues for New Contributors (`good-first-issue`)

We have curated a list of **20+ approachable issues** to help new contributors get started with the Pranor ecosystem. Each issue is tagged as a `good-first-issue` and contains step-by-step implementation guidance.

---

### 💻 Pranor Console (Web UI)

#### 1. Add tooltips to the auto-scaling status indicators
- **Repository**: `Pranor Console`
- **Description**: Add hover tooltips explaining "running", "paused", "stopped" states.
- **Difficulty**: Easy
- **Guidance**: Update the status badge components in the HTML template and add a standard CSS tooltip helper.

#### 2. Implement light/dark theme toggle persistence
- **Repository**: `Pranor Console`
- **Description**: Persist user theme choices using localStorage.
- **Difficulty**: Easy
- **Guidance**: Read/write `theme` key to localStorage on toggle and apply the appropriate classes during onload.

#### 3. Display memory unit conversions (bytes to MB/GB)
- **Repository**: `Pranor Console`
- **Description**: Format raw byte metrics from Pranor Deploy into human-readable MB/GB.
- **Difficulty**: Easy
- **Guidance**: Add a formatting utility in the javascript view layer to render bytes dynamically.

#### 4. Add clear search filter button to metrics view
- **Repository**: `Pranor Console`
- **Description**: Add an "x" button to clear search input fields.
- **Difficulty**: Easy
- **Guidance**: Bind an onClick event handler to clear inputs and trigger state updates.

---

### 📧 Pranor Notify (Communications)

#### 5. Add custom subject line length warning
- **Repository**: `Pranor Notify`
- **Description**: Return a warning header if a subject line exceeds 78 characters.
- **Difficulty**: Easy
- **Guidance**: Check subject length in `pkg/delivery` and log warning.

#### 6. Support HTML email content-type headers
- **Repository**: `Pranor Notify`
- **Description**: Enable setting explicit content-type headers for HTML delivery.
- **Difficulty**: Easy
- **Guidance**: Expose Content-Type mapping inside SMTP sender routines.

#### 7. Add validation for target phone formats (SMS)
- **Repository**: `Pranor Notify`
- **Description**: Validate SMS target phone formats (E.164 standard).
- **Difficulty**: Medium
- **Guidance**: Use regex validation inside target parsing function before enqueueing.

#### 8. Implement basic retry backoff configuration limits
- **Repository**: `Pranor Notify`
- **Description**: Expose max backoff interval duration via env configuration.
- **Difficulty**: Medium
- **Guidance**: Read environment configs inside DLQ retry loop to cap backoff delays.

---

### 🔄 Pranor Flow (Workflows)

#### 9. Add DAG cycle check command line validation
- **Repository**: `Pranor Flow`
- **Description**: Expose cycle checker validation as a CLI tool check.
- **Difficulty**: Medium
- **Guidance**: Update `main.go` flags to parse files and trigger `HasCycle` directly.

#### 10. Implement custom names for workflow checkpoints
- **Repository**: `Pranor Flow`
- **Description**: Allow workflows to assign custom names to checkpoint states.
- **Difficulty**: Easy
- **Guidance**: Expose name field in checkpoint schema and persist to disk.

#### 11. Add human approval gate timeout expiration limits
- **Repository**: `Pranor Flow`
- **Description**: Auto-expire approval gates after a specified duration.
- **Difficulty**: Medium
- **Guidance**: Check creation timestamps in execution loops and transition to failed state on timeout.

#### 12. Log compensation action execution steps
- **Repository**: `Pranor Flow`
- **Description**: Add telemetry/logs when saga compensations trigger.
- **Difficulty**: Easy
- **Guidance**: Insert log logs inside saga execution methods.

---

### 🗄️ Pranor Pool (Database proxy)

#### 13. Expose connection checkout duration metrics
- **Repository**: `Pranor Pool`
- **Description**: Track lease times for checked-out connection pools.
- **Difficulty**: Medium
- **Guidance**: Measure time elapsed between Acquire and Release.

#### 14. Add sqlite database path configuration validation
- **Repository**: `Pranor Pool`
- **Description**: Validate SQLite file paths before creating connection pools.
- **Difficulty**: Easy
- **Guidance**: Verify directories exist or create them cleanly inside setup.

#### 15. Support query statistics filtering by type
- **Repository**: `Pranor Pool`
- **Description**: Filter analytics output by SELECT/INSERT type.
- **Difficulty**: Medium
- **Guidance**: Parse query prefixes in analytics package to group statistics.

#### 16. Log deadlocked lease reclaims
- **Repository**: `Pranor Pool`
- **Description**: Print structured warning messages when janitor reclaims connection leases.
- **Difficulty**: Easy
- **Guidance**: Insert log lines in pool janitor reclaim loop.

---

### ⚡ Pranor Cache (Caching)

#### 17. Add custom key eviction alerts
- **Repository**: `Pranor Cache`
- **Description**: Log a warning when high rate of cache key evictions happens.
- **Difficulty**: Medium
- **Guidance**: Trigger alert metrics when evictions exceed limit.

#### 18. Support cache namespaces listings
- **Repository**: `Pranor Cache`
- **Description**: Expose an endpoint to list all registered cache namespaces.
- **Difficulty**: Easy
- **Guidance**: Expose new handler returning active namespaces keys.

---

### 🛠️ Compiler & Tooling (`Pranor`)

#### 19. Add warning on unused function arguments
- **Repository**: `Pranor`
- **Description**: Print warnings in semantic analyzer for unused parameters.
- **Difficulty**: Medium
- **Guidance**: Track argument usage mappings in typechecker ast.

#### 20. Implement colorized CLI error messages
- **Repository**: `Pranor`
- **Description**: Print compiler errors with ANSI escape colors.
- **Difficulty**: Easy
- **Guidance**: Add helper functions to format outputs using red/yellow prefixes.

#### 21. Document package import examples
- **Repository**: `Pranor`
- **Description**: Add example files for package structure imports.
- **Difficulty**: Easy
- **Guidance**: Document multi-file library imports in docs and examples.
