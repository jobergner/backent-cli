# Project Structure Examples

**Note:** These examples show the complete `PROJECT-STRUCTURE.md` file as a standalone document. The file is located at the repository root.

## Complete Example: TypeScript/React Application

```markdown
# Project Structure

*Generated on 2025-01-23 14:30 with the project-structure skill*

package.json                 # Node.js dependencies, npm scripts, project metadata
tsconfig.json                # TypeScript config: strict mode, path aliases, module resolution
.eslintrc.js                 # ESLint rules: Airbnb style guide, React, TypeScript plugins
.prettierrc                  # Prettier config: tab width, semicolons, formatting rules
README.md                    # Project overview, installation, development setup, deployment guide
docker-compose.yml           # Docker services: PostgreSQL, Redis for local development
src/
├── components/              # UI components: Button, Card, Form, Modal, Tabs. Design system, accessibility. Tests, Storybook. 15 modules
│   ├── Button/             # Button: primary/secondary/tertiary variants, loading states, icon support. Jest tests, Storybook stories
│   ├── Card/               # Card: header/body/footer sections, elevation variants. CardHeader, CardFooter components. Tests
│   ├── Form/               # Form inputs: Input, Select, Checkbox, Radio, TextArea. Validation, error handling, ARIA. 8 components
│   ├── Modal/              # Modal: backdrop, keyboard nav, focus trap, animation. ModalHeader, ModalBody, ModalFooter
│   └── Tabs/               # Tabs: keyboard nav, active state, lazy loading. Tab, TabPanel components. Tests
├── hooks/                   # React hooks: useAuth, useLocalStorage, useFetch, useDebounce. State, effects, logic. 12 files, tests
│   ├── auth/               # Auth hooks: useAuth, usePermissions, useSession. Login/logout, RBAC, sessions. 3 files, integration tests
│   └── data/               # Data hooks: useFetch, useQuery, useMutation. Loading states, error handling, retry. 5 files, API mocks
├── services/                # Business logic: AuthService, UserService, ProductService, AnalyticsService. Error handling, retry. 8 files
│   ├── api/                # HTTP client: axios interceptors, transformers, error handling, JWT refresh. Base client, typed endpoints
│   └── validators/         # Validation: Zod schemas for forms, API requests. Type-safe runtime checks. 12 files by domain
├── types/                   # TypeScript definitions: API contracts, domain models, utility types. api.types, user.types, product.types. 15 files
├── utils/                   # Utilities: formatDate, formatMoney, validate, debounce, throttle, deepClone. Edge case tests. 18 files
├── App.tsx                  # Root component: React Router, context providers (Auth, Theme, I18n), error boundary
└── main.tsx                 # Entry point: React DOM rendering, StrictMode, global CSS

## Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

### Date & Time
- **date-fns** (v2.30.0) - Modern date utility library with immutable operations
  - Used in: `src/utils/`, `src/services/reporting.ts`, `src/components/Calendar/` (15 files)
  - Common functions: `format()`, `parseISO()`, `differenceInDays()`, `addDays()`, `startOfWeek()`
  - Typical use cases: Date formatting in UI, business day calculations, report date ranges

### Validation & Schemas
- **zod** (v3.22.4) - TypeScript-first schema validation with type inference
  - Used in: `src/schemas/`, `src/api/validators/`, `src/forms/` (23 files)
  - Common patterns: `z.object()`, `z.string()`, `z.number()`, `z.array()`, `.parse()`, `.safeParse()`
  - Typical use cases: API request validation, form schemas with error messages, runtime type checking

### Data Transformation
- **lodash** (v4.17.21) - Comprehensive utility library for arrays, objects, and functions
  - Used in: Throughout codebase (47 files)
  - Common functions: `_.debounce()`, `_.groupBy()`, `_.uniqBy()`, `_.get()`, `_.merge()`, `_.cloneDeep()`
  - Typical use cases: Debouncing search inputs, grouping data for reports, deep object manipulation

### HTTP Client
- **axios** (v1.6.2) - Promise-based HTTP client with interceptors and automatic transforms
  - Used in: `src/services/*.service.ts`, `src/api/` (12 files)
  - Common patterns: Instance creation with interceptors, request/response transformation, error handling
  - Typical use cases: API communication, JWT token refresh, centralized error handling
```

---

## Complete Example: Python/Django Application

```markdown
# Project Structure

*Generated on 2025-01-23 14:30 with the project-structure skill*

manage.py                    # Django CLI: runserver, migrate, createsuperuser commands
wsgi.py                      # WSGI entry point for Gunicorn/uWSGI production deployment
requirements.txt             # Python dependencies for production
requirements-dev.txt         # Dev dependencies: pytest, black, flake8, mypy
pyproject.toml               # Python config: black, isort, pytest, mypy settings
README.md                    # Project overview, setup, architecture, deployment guide
docker-compose.yml           # Docker: PostgreSQL, Redis, Celery for local dev
src/
├── api/                     # REST API: DRF views, serializers, permissions, URLs. User, Product, Order. 8 test files, auth scenarios
│   ├── v1/                 # API v1: User, Product, Order resources. ViewSets, routers, nested serializers. Filtering, pagination, search
│   └── v2/                 # API v2: breaking changes. Enhanced permissions, bulk ops, async support. 12 endpoints
├── models/                  # Django ORM: User, Product, Order, Category. Relationships, validators, managers. 15 files, test fixtures
│   ├── user/               # User models: User, Profile, UserSettings. Auth fields, profile data, prefs. Custom manager, validators
│   └── commerce/           # Commerce models: Product, Order, Payment, Inventory. Business logic, workflows, payment. 8 files
├── services/                # Business logic: AuthService, OrderService, PaymentService, EmailService. Mocked deps in tests. 12 files
│   ├── auth/               # Auth services: JWT, password reset, email verify, social auth. OAuth2 support. 4 files
│   ├── order/              # Order services: inventory checks, payment, fulfillment, notifications. OrderService, FulfillmentService
│   └── integrations/       # Third-party: Stripe, PayPal, SendGrid, analytics. Webhook handlers. 6 files
├── utils/                   # Utilities: date format, validation, string format, file handling. Edge case tests. 18 files
└── settings/                # Django settings by env: base, dev, prod, staging, test. Security, DB, middleware config
    ├── base.py             # Base settings: installed apps, middleware, templates, static files, i18n
    ├── development.py      # Dev settings: debug mode, Debug Toolbar, relaxed security, SQLite
    ├── production.py       # Prod settings: security hardening, PostgreSQL, Redis cache, Cloudflare CDN, Sentry
    └── test.py             # Test settings: in-memory DB, disabled migrations, mocked services

## Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

### HTTP Client
- **requests** (v2.31.0) - Simple, elegant HTTP library with excellent API
  - Used in: `src/services/`, `src/integrations/` (8 files)
  - Common functions: `requests.get()`, `requests.post()`, `.json()`, session management
  - Typical use cases: Third-party API integration, webhook handling, external service calls

### Data Validation
- **pydantic** (v2.5.0) - Data validation using Python type annotations
  - Used in: `src/models/`, `src/api/schemas/` (34 files)
  - Common patterns: `BaseModel`, `Field()`, `validator()`, `model_validate()`, `model_dump()`
  - Typical use cases: API request/response models, configuration validation, data parsing

### Date & Time
- **arrow** (v1.3.0) - Better dates and times for Python with friendly API
  - Used in: `src/utils/`, `src/services/reporting.py` (11 files)
  - Common functions: `arrow.now()`, `.shift()`, `.humanize()`, `.format()`, timezone handling
  - Typical use cases: Date arithmetic, timezone conversions, human-readable date formatting
```

---

## Complete Example: Go Application

```markdown
# Project Structure

*Generated on 2025-01-23 14:30 with the project-structure skill*

go.mod                       # Go module dependencies and versions
go.sum                       # Cryptographic checksums for dependencies
README.md                    # Project overview, build instructions, API docs, deployment guide
Makefile                     # Build automation: build, test, lint, docker, deploy targets
Dockerfile                   # Multi-stage Docker build for optimized production image
.env.example                 # Environment variables: database URLs, API keys
cmd/
├── server/                  # HTTP server: graceful shutdown, signal handling. Main API entrypoint
└── cli/                     # CLI tools: cobra for user management, migrations, admin tasks
internal/
├── handlers/                # HTTP handlers: chi router. User, Auth, Product, Order, Health. 12 test files, mocking, table-driven
│   ├── user/               # User handlers: CRUD, profile updates, password changes. Pagination, filtering, error handling
│   ├── auth/               # Auth handlers: login, logout, token refresh, password reset. JWT generation, validation
│   └── product/            # Product handlers: search, filter, sort, pagination. Inventory, bulk operations
├── services/                # Business logic: User, Auth, Product, Order. 15 test files, table-driven, mocked deps
│   ├── user/               # User service: validation, bcrypt hashing, email verify, profile. UserService interface
│   ├── auth/               # Auth service: JWT, token refresh, sessions, RBAC. Redis token storage
│   └── product/            # Product service: inventory, stock tracking, pricing, search. ProductService interface
├── models/                  # Data models: GORM tags. User, Product, Order, Category. Relationships, validators, hooks. 8 tests
├── database/                # Database: connection pool, golang-migrate, seeders. PostgreSQL config, retry logic
└── middleware/              # Middleware: logging, auth, CORS, rate limiting, request IDs. 6 test files, mock handlers
    ├── auth/               # JWT middleware: token validation, claims extraction, context injection. Expiration, refresh
    ├── logging/            # Logging middleware: structured zap logging. Response time, status codes, request IDs
    └── rate/               # Rate limiting: token bucket algorithm. Per-endpoint limits, Redis backend
pkg/
├── validator/               # Validation utilities: email, phone, URL, UUID. Custom error messages
└── config/                  # Config: viper for env vars, files, defaults

## Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

### HTTP Router
- **gin** (github.com/gin-gonic/gin v1.9.1) - Fast HTTP web framework with middleware support
  - Used in: `internal/handlers/`, `cmd/server/` (18 files)
  - Common patterns: Router groups, middleware chains, JSON binding, validation
  - Typical use cases: REST API routing, request validation, response serialization

### Validation
- **validator** (github.com/go-playground/validator/v10 v10.15.5) - Struct validation with tags
  - Used in: `internal/models/`, `internal/handlers/`, `pkg/validator/` (25 files)
  - Common tags: `validate:"required,email"`, `validate:"min=8,max=32"`, custom validators
  - Typical use cases: Request body validation, struct field validation, custom validation rules

### Database Driver
- **pgx** (github.com/jackc/pgx/v5 v5.5.0) - PostgreSQL driver with connection pooling
  - Used in: `internal/database/`, `internal/services/` (14 files)
  - Common patterns: Connection pooling, prepared statements, transaction management, migrations
  - Typical use cases: Database queries, connection management, transaction handling
```

---

## Complete Example: Rust Application

```markdown
# Project Structure

*Generated on 2025-01-23 14:30 with the project-structure skill*

src/
├── main.rs                  # Binary entry point with Tokio async runtime initialization
├── lib.rs                   # Library root re-exporting public modules
├── api/                     # HTTP API endpoints and routing
│   ├── mod.rs              # API module re-exporting handlers and routes
│   ├── handlers.rs         # Request handlers with validation and error handling
│   ├── routes.rs           # Axum router configuration with middleware
│   ├── middleware.rs       # Authentication and logging middleware
│   └── tests/
│       └── *.rs            # Integration tests with mock HTTP requests (8 files)
├── services/                # Business logic layer with domain operations
│   ├── mod.rs              # Service module exports
│   ├── user.rs             # User service with password hashing and validation
│   ├── auth.rs             # Authentication service with JWT token management
│   ├── product.rs          # Product service with inventory operations
│   └── tests/
│       └── *.rs            # Service tests with mock repositories (12 files)
├── models/                  # Data models with Serde serialization
│   ├── mod.rs              # Model module exports
│   ├── user.rs             # User model with authentication fields
│   ├── product.rs          # Product model with inventory tracking
│   ├── order.rs            # Order model with status workflow
│   └── _internal.rs        # Internal helper types not exposed in mod.rs
├── database/                # Database connection and query operations
│   ├── mod.rs              # Database module exports
│   ├── pool.rs             # Connection pool management with SQLx
│   ├── migrations.rs       # Database migration runner
│   └── queries.rs          # Reusable query functions with type safety
├── utils/                   # Utility functions and helpers
│   ├── mod.rs              # Utility module exports
│   ├── date.rs             # Date formatting and timezone conversion using Chrono
│   ├── validation.rs       # Email and phone validation utilities
│   └── crypto.rs           # Password hashing and token generation utilities
└── config.rs                # Configuration loading from environment with validation

Cargo.toml                   # Rust package manifest with dependencies

## Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

### Async Runtime
- **tokio** (v1.35.1) - Asynchronous runtime for Rust with full feature set
  - Used in: Throughout codebase (56 files)
  - Common patterns: `#[tokio::main]`, `tokio::spawn()`, async/await, channels, timers
  - Typical use cases: HTTP server, concurrent tasks, async I/O, background jobs

### Serialization
- **serde** (v1.0.193) - Framework for serializing and deserializing data structures
  - Used in: `src/models/`, `src/api/`, `src/config.rs` (42 files)
  - Common patterns: `#[derive(Serialize, Deserialize)]`, `serde_json::from_str()`, custom serializers
  - Typical use cases: JSON/YAML/TOML parsing, API serialization, configuration files

### Date & Time
- **chrono** (v0.4.31) - Date and time library with timezone and duration support
  - Used in: `src/models/`, `src/services/`, `src/utils/` (19 files)
  - Common types: `DateTime<Utc>`, `NaiveDateTime`, duration calculations, parsing
  - Typical use cases: Timestamp handling, date arithmetic, parsing ISO 8601 formats

### HTTP Framework
- **axum** (v0.7.3) - Ergonomic web framework built on Tokio and Hyper
  - Used in: `src/api/` (15 files)
  - Common patterns: Router with extractors, JSON responses, middleware layers, state sharing
  - Typical use cases: REST API endpoints, request parsing, response serialization, middleware

### Database
- **sqlx** (v0.7.3) - Async SQL toolkit with compile-time query verification
  - Used in: `src/database/`, `src/services/` (23 files)
  - Common patterns: Connection pooling, compile-time query checking, transactions, migrations
  - Typical use cases: Database queries, connection management, type-safe SQL operations
```

---

## Minimal Example: Small Utility Library

```markdown
# Project Structure

*Generated on 2025-01-23 14:30 with the project-structure skill*

src/
├── index.ts                 # Main entry point exporting all utilities
├── string.ts                # String manipulation and formatting utilities
├── number.ts                # Number formatting and arithmetic utilities
├── array.ts                 # Array transformation and filtering utilities
└── *.test.ts                # Unit tests with edge case coverage (4 files)

package.json                 # Package manifest with zero dependencies
README.md                    # Documentation with API reference and examples
tsconfig.json                # TypeScript compiler configuration with strict mode

## Utility Dependencies

None - This is a zero-dependency utility library.
```

---

## Tree Formatting Examples

### Correct Tree Characters

```
directory/
├── first-item
├── second-item
├── third-item
└── last-item
```

### Nested Tree Structure

```
parent/
├── child-1/
│   ├── grandchild-1
│   └── grandchild-2
├── child-2/
│   ├── grandchild-3
│   ├── grandchild-4
│   └── grandchild-5
└── child-3
```

### Complex Nesting with Files and Directories

```
src/
├── components/              # Component directory
│   ├── Button/             # Nested component
│   │   ├── Button.tsx      # File in nested directory
│   │   ├── *.test.tsx      # Pattern group (4 files)
│   │   └── index.ts        # Another file
│   ├── Form/               # Another nested component
│   │   ├── FormInput.tsx
│   │   └── index.ts
│   └── index.ts            # File at components level
├── hooks/                   # Another top-level directory
│   ├── useAuth.ts
│   └── useFetch.ts
└── App.tsx                  # File at src level
```

## Common Mistakes to Avoid

### ❌ Wrong: Inconsistent Characters

```
directory/
├── first
├── second
├── third          # Should be └── for last item
```

### ❌ Wrong: Missing Vertical Continuation

```
parent/
├── child-1/
  ├── grandchild   # Missing │ for vertical line
```

### ❌ Wrong: Obvious Comments

```
Button.tsx         # Button component
auth.service.ts    # Auth service
*.test.ts          # Tests
```

### ✅ Correct: Meaningful Comments

```
Button.tsx         # Accessible button with size and variant props
auth.service.ts    # JWT token validation and session management
*.test.ts          # Unit tests using Jest with React Testing Library (8 files)
```
