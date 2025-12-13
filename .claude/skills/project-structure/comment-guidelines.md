# Comment Guidelines

## Core Principle

Describe **WHAT THE CONTENT IS ABOUT**, not just restate the name.

Focus on actual functionality, logic, data, or purpose inside the file/directory.

## Length Guidelines

- **Root-level files:** ~8-12 words
- **Directories (all nesting levels):** ~14-23 words (telegraphic, keyword-dense)

These are guidelines, not strict limits. Quality over exact word count.

## Content Guidelines

### DO:
✅ Describe what's INSIDE the file/directory
✅ Focus on actual functionality and logic
✅ Be specific about what code does or contains
✅ Mention technologies, patterns, or domains
✅ For directories, include file types, notable files, and counts
✅ For directories, mention key characteristics and purpose

### DON'T:
❌ Just restate the filename or directory name
❌ Use obvious or redundant descriptions
❌ Be too generic or vague
❌ Forget to actually read the file content
❌ Omit important details about directory contents

## Examples: Root-Level Files (~8-12 words)

### Good Examples

```
package.json                # Node.js dependencies, scripts, and project metadata configuration
README.md                   # Project overview, installation instructions, and development setup guide
tsconfig.json               # TypeScript compiler configuration with strict mode and path aliases
.env.example                # Environment variable template with API keys and database connection settings
docker-compose.yml          # Docker services configuration for local development with PostgreSQL and Redis
```

### Bad Examples (Too Obvious)

```
package.json                # Package file ❌
README.md                   # README ❌
tsconfig.json               # TypeScript config ❌
.env.example                # Environment variables ❌
docker-compose.yml          # Docker compose file ❌
```

## Examples: Directories (~14-23 words)

**IMPORTANT:** ALL directories at ALL nesting levels must have comments, not just top-level directories.

**Style:** Telegraphic with minimal connectors. Core purpose: key items. Tech/patterns. Counts.

### Good Examples (All Levels)

```
# Top-level directories
src/                        # Source: React components (45), utilities (12), API services (8). Entry App.tsx
components/                 # UI components: Button, Card, Form, Modal. Design system, accessibility, tests, stories. 15 modules
api/                        # REST endpoints: user, product operations. Auth middleware, validation. 8 handlers by resource
types/                      # TypeScript definitions: API contracts, domain models, utility types. Request/response schemas, validation
hooks/                      # React hooks: useAuth, useLocalStorage, useFetch. State/effects management. 12 files
services/                   # Business logic: AuthService, UserService, ProductService. Error handling, retry logic, API communication
models/                     # Database entities: User, Product, Order. Sequelize ORM, validation, relationships, query methods
utils/                      # Utilities: formatting, validation, date handling. 18 helper functions for common operations

# Nested subdirectories (also need comments!)
components/Button/          # Button component: primary/secondary/tertiary variants, loading states. ARIA accessibility, tests, Storybook stories
components/Form/            # Form inputs: Input, Select, Checkbox, Radio. Validation, error handling, accessibility. 8 components
api/handlers/               # HTTP handlers: users, products, orders. Auth middleware, validation, error handling
services/auth/              # Auth logic: JWT validation, token refresh, session management. Password hashing, AuthContext integration
```

### Bad Examples (Too Short)

```
components/                 # UI components ❌
api/                        # API endpoints ❌
types/                      # Type definitions ❌
hooks/                      # React hooks ❌
services/                   # Services ❌
utils/                      # Utilities ❌
```

## How to Analyze Directories

**CRITICAL:** ALL directories at ALL nesting levels require ~14-23 word telegraphic comments with purpose, key items, tech, and counts.

This includes:
- Top-level directories (e.g., `src/`, `components/`)
- Nested subdirectories (e.g., `components/Button/`, `api/handlers/`)
- Deeply nested directories (e.g., `components/Button/tests/`, `services/auth/validators/`)

### Step 1: List Directory Contents

Use Glob to see what's inside:

```bash
Glob: "src/components/*"
Glob: "src/services/*.ts"
```

### Step 2: Identify the Collective Purpose and Details

Ask yourself:
- **What type of files?** Components, services, utilities, models, types?
- **What domain?** User management, authentication, products, API endpoints?
- **What technologies?** React components, REST handlers, database models?
- **What patterns or features?** Design system, validation, error handling, accessibility?
- **How many files?** Count the files to include in the comment
- **Notable files?** What are the key/important files to mention by name?
- **Key characteristics?** Tests, stories, styles, variants, states?

### Step 3: Read Representative Files

Read 1-2 files to understand the directory's role:

```bash
Read: src/components/Button/Button.tsx
Read: src/components/Form/FormInput.tsx
```

### Step 4: Write Directory Comment (~14-23 words)

Use telegraphic style: Core purpose: key items. Tech/patterns. Counts.

**Good Examples:**
```
components/                 # UI components: Button, Card, Form, Modal. Design system, accessibility, tests, stories. 15 modules
services/                   # Business logic: AuthService, UserService, ProductService. Error handling, retry logic, API communication
models/                     # Database entities: User, Product, Order. Sequelize ORM, validation, relationships, query methods
handlers/                   # HTTP handlers: REST endpoints. Auth middleware, validation, error handling. 8 files by resource
utils/                      # Utilities: formatting, validation, date handling. 18 helper functions for common operations
```

**Bad Examples:**
```
components/                 # Components ❌ (too short, no keywords)
services/                   # Service files ❌ (too generic, no items listed)
models/                     # Models ❌ (missing tech and items)
handlers/                   # Handlers ❌ (no key items or patterns)
utils/                      # Utilities ❌ (missing specific utilities)
```

### Step 5: Use Telegraphic Style

**Format pattern:** `Purpose: key items. Tech/patterns. Counts`

Always include:
- **Core purpose** using minimal connectors (Auth, UI components, Business logic)
- **Key items** listed with colons or commas (Button, Card, Form OR AuthService, UserService)
- **Tech/patterns** mentioned concisely (JWT, React hooks, Sequelize ORM, middleware)
- **Counts** when relevant (15 modules, 8 files, 12 utilities)

**Style tips:**
- Use colons to introduce lists: "UI components: Button, Card, Form"
- Use periods to separate idea groups: "Design system. Accessibility. Tests"
- Omit articles (the, a, an) and unnecessary verbs (contains, includes, provides)
- List concrete items rather than categories when possible

## How to Analyze Files

### Step 1: Read the File

Use Read tool to examine actual content:

```bash
Read: src/components/Button/Button.tsx
```

### Step 2: Look for Key Elements

- **Main exports:** What classes, functions, or components are exported?
- **Function names:** What operations does the code perform?
- **Class definitions:** What entities or concepts are modeled?
- **Imports:** What dependencies indicate the file's purpose?
- **Comments/docs:** What does existing documentation say?

### Step 3: Identify Specifics

Be specific about:
- **Technologies:** Jest, React Testing Library, CSS Modules, JWT
- **Patterns:** Table-driven tests, dependency injection, observer pattern
- **Domains:** User authentication, product catalog, payment processing
- **Features:** Accessibility, responsive design, validation, caching

### Example Analysis

**File:** `src/components/Button/Button.tsx`

```typescript
import { forwardRef } from 'react'

interface ButtonProps {
  size?: 'sm' | 'md' | 'lg'
  variant?: 'primary' | 'secondary'
  disabled?: boolean
  ariaLabel?: string
}

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ size = 'md', variant = 'primary', ...props }, ref) => {
    // Implementation with ARIA attributes
    return <button ref={ref} aria-label={props.ariaLabel} {...props} />
  }
)
```

**Analysis:**
- Exports: `Button` component
- Features: Multiple size props, variant props, accessibility (ARIA), forwardRef
- Technology: React

**Comment:** "Accessible button component with multiple size variants"

---

**File:** `src/hooks/useAuth.ts`

```typescript
import { useState, useEffect, useContext } from 'react'
import { AuthContext } from '../contexts/AuthContext'

export function useAuth() {
  const [user, setUser] = useState(null)
  const [loading, setLoading] = useState(true)

  // Token refresh logic, session management
  // Login/logout handlers

  return { user, loading, login, logout, refresh }
}
```

**Analysis:**
- Exports: `useAuth` hook
- Features: User state, loading state, login/logout handlers, token refresh
- Purpose: Authentication management

**Comment:** "Authentication state and login/logout handlers with session management"

## Language-Specific Tips

### JavaScript/TypeScript

Look for:
- `export` statements (what's the API?)
- Component props (what features?)
- Hook dependencies (what does it manage?)
- Type definitions (what shape is the data?)

### Python

Look for:
- Class definitions (`class User:`)
- Function signatures (`def process_payment():`)
- Decorators (`@dataclass`, `@validator`)
- Docstrings (existing documentation)

### Go

Look for:
- Exported identifiers (uppercase)
- Package comment (`// Package handlers provides...`)
- Struct definitions (data models)
- Interface definitions (contracts)

### Rust

Look for:
- `pub` items (public API)
- Trait implementations
- Struct definitions
- Module-level comments (`//!`)

## Special Cases

### Configuration Files

Be specific about what they configure:

```
tsconfig.json               # TypeScript compiler configuration with strict mode and path aliases
.eslintrc.js                # ESLint rules enforcing Airbnb style guide with React plugins
jest.config.js              # Jest test configuration with coverage thresholds and setup files
```

### Entry Points

Describe what they initialize:

```
main.tsx                    # Application entry point with React DOM rendering and providers
index.ts                    # Package barrel exporting all public API components
__init__.py                 # Package initialization with database and cache setup
```

### README Files

Describe the section/module they document:

```
README.md                   # Project overview, installation instructions, and development setup guide
api/README.md               # API documentation with endpoint specifications and authentication details
```

## Common Pitfalls

### Pitfall 1: Too Generic

❌ `user.service.ts # User service`
✅ `user.service.ts # User CRUD operations with profile validation and password hashing`

### Pitfall 2: Restating Filename

❌ `useLocalStorage.ts # Hook that uses local storage`
✅ `useLocalStorage.ts # React hook for persistent browser storage management`

### Pitfall 3: Missing Details in Directory Comments

❌ `components/ # UI components`
✅ `components/ # UI components: Button, Card, Form, Modal. Design system, accessibility, tests, stories. 15 modules`

### Pitfall 4: Not Reading the File

Don't guess based on filename. Always read to understand actual content.

❌ Assuming `auth.service.ts` just does "authentication"
✅ Reading and discovering it handles JWT validation, session refresh, and role-based access

## Quality Checklist

Before writing a comment, ask:

**For all comments:**
- [ ] Did I actually read the file/directory content?
- [ ] Am I describing WHAT's inside, not just the name?
- [ ] Did I mention specific technologies or patterns?
- [ ] Is this specific enough to be useful?
- [ ] Would someone unfamiliar with the codebase understand this?
- [ ] Am I avoiding obvious statements?

**For directory comments specifically:**
- [ ] Did I use telegraphic style (Purpose: items. Tech. Counts)?
- [ ] Did I list key items/files by name?
- [ ] Did I mention tech/patterns concisely?
- [ ] Is my comment 14-23 words, keyword-dense?
- [ ] Did I omit unnecessary articles and verbs?

If you can answer "yes" to all applicable questions, your comment is likely good quality.
