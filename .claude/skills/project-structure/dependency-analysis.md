# Dependency Analysis Guide

## Purpose

Document utility libraries to help developers avoid reimplementing existing functionality.

**Focus:** Production dependencies only (exclude dev/test dependencies).

**Important:** The Utility Dependencies section should ONLY be included in the output if meaningful production dependencies are found (3+ file usage or critical utilities like JWT, validation, etc.). If no meaningful dependencies exist, omit the entire section.

## Step 1: Detect Package Manager

Identify which package manager(s) the project uses:

```bash
# Check for Node.js
test -f package.json && echo "node_found"

# Check for Python
test -f pyproject.toml && echo "python_toml_found"
test -f requirements.txt && echo "python_requirements_found"

# Check for Rust
test -f Cargo.toml && echo "rust_found"

# Check for Go
test -f go.mod && echo "go_found"

# Check for Java
test -f pom.xml && echo "maven_found"
test -f build.gradle && echo "gradle_found"

# Check for C#
ls *.csproj 2>/dev/null && echo "csproj_found"
```

## Step 2: Read Production Dependencies

Extract ONLY production dependencies from package files:

### Node.js (package.json)

```json
{
  "dependencies": {          // ← Read this section ONLY
    "lodash": "^4.17.21",
    "date-fns": "^2.30.0",
    "zod": "^3.22.4"
  },
  "devDependencies": {       // ← SKIP this section
    "vitest": "^1.0.0",
    "eslint": "^8.0.0"
  }
}
```

**Read:** `dependencies` section only

### Python (pyproject.toml)

```toml
[project]
dependencies = [            # ← Read this
    "requests>=2.31.0",
    "pandas>=2.0.0",
    "pydantic>=2.0.0"
]

[project.optional-dependencies]
dev = [                     # ← SKIP this
    "pytest>=7.0.0",
    "black>=23.0.0"
]

[tool.poetry.dependencies]  # ← Read this (Poetry format)
python = "^3.11"
requests = "^2.31.0"

[tool.poetry.dev-dependencies]  # ← SKIP this
pytest = "^7.0.0"
```

**Read:** `[project.dependencies]` or `[tool.poetry.dependencies]`
**Skip:** `[project.optional-dependencies]`, dev groups

### Python (requirements.txt)

```
requests>=2.31.0           # ← All are production (no dev distinction)
pandas>=2.0.0
pydantic>=2.0.0
```

**Note:** requirements.txt doesn't separate dev/prod. Assume all are production unless there's a separate `requirements-dev.txt`.

### Rust (Cargo.toml)

```toml
[dependencies]              # ← Read this
serde = "1.0"
tokio = { version = "1.35", features = ["full"] }
chrono = "0.4"

[dev-dependencies]          # ← SKIP this
criterion = "0.5"
```

**Read:** `[dependencies]` only

### Go (go.mod)

```
module example.com/myapp

go 1.21

require (                   # ← Read all (Go doesn't separate dev/prod)
    github.com/gin-gonic/gin v1.9.1
    github.com/go-playground/validator/v10 v10.15.5
)
```

**Read:** All `require` directives (Go doesn't distinguish dev/prod)

### Java (pom.xml)

```xml
<dependencies>
    <!-- Production dependency -->
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
        <version>2.15.0</version>
    </dependency>

    <!-- Dev dependency (skip) -->
    <dependency>
        <groupId>junit</groupId>
        <artifactId>junit</artifactId>
        <version>4.13.2</version>
        <scope>test</scope>  <!-- ← SKIP if scope is test -->
    </dependency>
</dependencies>
```

**Read:** Dependencies WITHOUT `<scope>test</scope>`

## Step 3: Categorize Dependencies

Common utility categories:

### Date & Time
- **Node.js:** date-fns, dayjs, luxon, moment (deprecated)
- **Python:** arrow, pendulum, python-dateutil
- **Rust:** chrono, time
- **Go:** time (stdlib), go-dateutil
- **Java:** java.time (stdlib), Joda-Time

### Validation
- **Node.js:** zod, yup, joi, ajv, class-validator
- **Python:** pydantic, marshmallow, cerberus, voluptuous
- **Rust:** validator, garde
- **Go:** go-playground/validator
- **Java:** javax.validation, Hibernate Validator

### Data Transformation
- **Node.js:** lodash, ramda, fp-ts
- **Python:** pandas, more-itertools, toolz
- **Rust:** itertools, rayon
- **Go:** go-funk, lo (lodash-like)

### HTTP/API Client
- **Node.js:** axios, ky, got, node-fetch
- **Python:** requests, httpx, aiohttp
- **Rust:** reqwest, hyper
- **Go:** net/http (stdlib), resty
- **Java:** OkHttp, Apache HttpClient

### String Manipulation
- **Node.js:** string-case libraries, slugify, chalk (formatting)
- **Python:** inflection, python-slugify, unidecode
- **Rust:** heck (case conversion), slug
- **Go:** strings (stdlib), go-slugify

### Functional Utilities
- **Node.js:** lodash, ramda, fp-ts
- **Python:** toolz, funcy, fn.py
- **Rust:** itertools, tap
- **Go:** lo, go-funk

### Parsing
- **Node.js:** qs, csv-parse, xml2js, yaml
- **Python:** csv (stdlib), PyYAML, xmltodict
- **Rust:** serde (json/yaml/toml), csv
- **Go:** encoding/json (stdlib), yaml.v3

### Crypto/Security
- **Node.js:** bcrypt, jsonwebtoken, uuid, crypto-js
- **Python:** cryptography, PyJWT, bcrypt, secrets (stdlib)
- **Rust:** bcrypt, jsonwebtoken, uuid, ring
- **Go:** crypto (stdlib), golang-jwt, uuid

## Exclude from Documentation

**Core frameworks:** React, Vue, Angular, Django, Flask, Express, FastAPI, Spring Boot

**Build tools:** webpack, vite, rollup, babel, tsc, cargo, maven-plugins

**Dev/test dependencies:** jest, pytest, vitest, mocha, junit, eslint, prettier

**Type definitions:** @types/* packages (TypeScript)

**Standard library:** Don't document built-in language features

## Step 4: Analyze Actual Usage

For each utility dependency identified:

### 4A: Find Import Patterns

Use Grep to locate imports:

```bash
# Node.js/TypeScript
Grep: "from ['\"](lodash|date-fns|zod)" --output-mode files_with_matches

# Python
Grep: "^import (pandas|requests|pydantic)" --output-mode files_with_matches
Grep: "^from (pandas|requests|pydantic)" --output-mode files_with_matches

# Rust
Grep: "use (serde|tokio|chrono)" --output-mode files_with_matches

# Go
Grep: "\"github\\.com/gin-gonic/gin\"" --output-mode files_with_matches
```

### 4B: Count Usage Locations

Count how many files import each dependency:
- **3+ files:** Document it (meaningful usage)
- **1-2 files:** Skip (minimal usage)
- **Exception:** Critical utilities (e.g., JWT, validation) document even if few files

### 4C: Identify Common Functions

Read 2-3 representative files to understand usage patterns:

```bash
# Example: Find common lodash functions
Grep: "import.*from ['\""]lodash" --output-mode content -C 2

# Example: Find zod schema patterns
Grep: "z\\.(object|string|number|array)" --output-mode content -n

# Example: Find date-fns usage
Grep: "(format|parseISO|differenceInDays)" --output-mode content -n
```

Look for:
- **Specific functions:** `_.debounce()`, `_.groupBy()`, `format()`, `parseISO()`
- **Usage patterns:** "Form validation", "API requests", "Date formatting in reports"

## Step 5: Document Format

Create "Utility Dependencies" subsection within the "## Project Structure" section of CLAUDE.md:

```markdown
### Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

#### [Category Name]
- **[package-name]** (v[version]) - [Brief description]
  - Used in: [key directories or patterns] ([N] files)
  - Common functions: `function1()`, `function2()`, `function3()`
  - Typical use cases: [patterns like "API request handling", "form validation"]
```

**Important:** This is a subsection (###) within the "## Project Structure" section. Category headings use #### (level 4).

## Real-World Examples

### Node.js/TypeScript Project

```markdown
### Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

#### Date & Time
- **date-fns** (v2.30.0) - Modern date utility library with immutable operations
  - Used in: `src/utils/`, `src/services/reporting.ts`, `src/components/Calendar/` (15 files)
  - Common functions: `format()`, `parseISO()`, `differenceInDays()`, `addDays()`, `startOfWeek()`
  - Typical use cases: Date formatting in UI, business day calculations, report date ranges

#### Validation & Schemas
- **zod** (v3.22.4) - TypeScript-first schema validation with type inference
  - Used in: `src/schemas/`, `src/api/validators/`, `src/forms/` (23 files)
  - Common patterns: `z.object()`, `z.string()`, `z.number()`, `.parse()`, `.safeParse()`
  - Typical use cases: API request validation, form schemas, runtime type checking

#### Data Transformation
- **lodash** (v4.17.21) - Comprehensive utility library for arrays, objects, functions
  - Used in: Throughout codebase (47 files)
  - Common functions: `_.debounce()`, `_.groupBy()`, `_.uniqBy()`, `_.get()`, `_.cloneDeep()`
  - Typical use cases: Debouncing search, grouping report data, deep object manipulation

#### HTTP Client
- **axios** (v1.6.2) - Promise-based HTTP client with interceptors and transforms
  - Used in: `src/services/*.service.ts`, `src/api/` (12 files)
  - Common patterns: Instance creation with interceptors, request/response transformation
  - Typical use cases: API communication, JWT token refresh, centralized error handling
```

### Python Project

```markdown
### Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

#### HTTP Client
- **requests** (v2.31.0) - Simple, elegant HTTP library with excellent API
  - Used in: `src/services/`, `src/integrations/` (8 files)
  - Common functions: `requests.get()`, `requests.post()`, `.json()`, session management
  - Typical use cases: Third-party API integration, webhook handling

#### Data Validation
- **pydantic** (v2.5.0) - Data validation using Python type annotations
  - Used in: `src/models/`, `src/api/schemas/` (34 files)
  - Common patterns: `BaseModel`, `Field()`, `validator()`, `model_validate()`
  - Typical use cases: API request/response models, configuration validation, data parsing

#### Data Processing
- **pandas** (v2.1.4) - Powerful data analysis and manipulation library
  - Used in: `src/analytics/`, `src/reports/` (12 files)
  - Common functions: `read_csv()`, `DataFrame`, `groupby()`, `merge()`, `to_excel()`
  - Typical use cases: Report generation, data aggregation, CSV/Excel processing
```

### Rust Project

```markdown
### Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

#### Serialization
- **serde** (v1.0.193) - Framework for serializing and deserializing Rust data structures
  - Used in: Throughout codebase (56 files)
  - Common patterns: `#[derive(Serialize, Deserialize)]`, `serde_json::from_str()`
  - Typical use cases: JSON/YAML/TOML parsing, API serialization, config files

#### Async Runtime
- **tokio** (v1.35.1) - Asynchronous runtime for Rust with full features
  - Used in: `src/server/`, `src/services/`, `src/db/` (28 files)
  - Common patterns: `#[tokio::main]`, `tokio::spawn()`, async/await, channels
  - Typical use cases: HTTP server, concurrent tasks, async I/O operations

#### Date & Time
- **chrono** (v0.4.31) - Date and time library with timezone support
  - Used in: `src/models/`, `src/services/` (19 files)
  - Common types: `DateTime<Utc>`, `NaiveDateTime`, duration calculations
  - Typical use cases: Timestamp handling, date arithmetic, parsing ISO 8601
```

## Analysis Workflow

1. **Read dependency file** → Get list of production dependencies
2. **Categorize** → Group by utility type (date, validation, HTTP, etc.)
3. **For each utility dependency:**
   - Use Grep to find import locations (get file count)
   - If 3+ files or critical utility:
     - Read 2-3 files to understand usage
     - Extract common functions/patterns
     - Document in appropriate category
4. **Skip if:**
   - Dev/test dependency
   - Framework or build tool
   - Type definitions only
   - Used in fewer than 3 files (unless critical)

## Performance Tips

- **Batch Grep searches:** Search for multiple packages at once
- **Focus on high-usage:** Prioritize dependencies used in 10+ files
- **Read strategically:** Don't read every file, sample 2-3 representative files
- **Cache categories:** Build mental model of common utility types
