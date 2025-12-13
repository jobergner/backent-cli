# Project Structure

*Generated on November 25, 2025 9:05 AM CET with the project-structure skill*

```
.claude/                       # Claude Code configuration: custom commands, principles, skills for AI-assisted development
├── commands/                  # Custom slash commands: bootstrap, learn, why for project workflows
├── principles/                # AI assistant principles documentation guiding behavior, code style
└── skills/                    # Custom skills definitions including project-structure documentation generator
    └── project-structure/     # Project structure skill: README, guidelines, dependency analysis, comment examples
assets/                        # Static assets: gopher.png logo for README documentation
docker-compose.yml             # Docker Compose configuration for end-to-end testing environment setup
example.config.json            # Example configuration defining state types, actions, responses for code generation
examples/                      # Generated example code: client, server, state engine, webclient demonstrating output
├── client/                    # Client example: actions, controller, message routing. WebSocket client implementation. 6 files
├── connect/                   # Connection utilities: WebSocket connection helpers for client-server communication
├── logging/                   # Logging utilities: structured logging helpers using zerolog for debugging
├── message/                   # Message protocol: action kinds, parameters, responses for WebSocket communication. 4 files
├── server/                    # Server example: client management, lobby, rooms, ticking, trigger actions. 10 files
├── state/                     # State engine example: CRUD operations, assembling, tree structures, references. 44+ generated files
└── webclient/                 # TypeScript/JavaScript webclient: config, tests, type definitions, NPM package. Jest testing
generate/                      # Code generation tools: decltostring snapshot tool, static code generator, TypeScript test declarations
├── decltostring/              # Declaration snapshot tool: converts Go declarations to string representations for testing
│   └── testdata/              # Test fixtures: input Go files, expected output for decltostring validation
│       ├── actual_output/     # Actual generated output from decltostring for comparison with expected
│       ├── expected_output/   # Expected golden files: reference output for decltostring correctness verification
│       └── input/             # Sample input Go declarations: book, fax, letter types for testing
├── static_code/               # Static code generator: produces boilerplate code templates embedded in packages
└── typescript_test_decls/     # TypeScript test declarations generator: creates type definitions for webclient testing
generate.sh                    # Generation orchestration script: runs decltostring, static code generators in sequence
go.mod                         # Go module definition: dependencies jennifer, easyjson, zerolog, websocket, testing frameworks
go.sum                         # Go dependency checksums: integrity verification for module dependencies
main.go                        # CLI entry point: parses config, generates packages, orchestrates code generation pipeline
pkg/                           # Core packages: AST parsing, config reading, factories, validation, code generation
├── ast/                       # Abstract Syntax Tree: parses config.json into structured representation. Types, actions, fields
├── config/                    # Configuration reader: loads JSON config files, validates syntax, provides parsed data
├── env/                       # Environment utilities: directory creation, path manipulation for output management
├── factory/                   # Code generation factories: client, server, state, webclient, message protocol generators
│   ├── client/                # Client factory: generates Go client code, actions, controller. 6+ modules
│   ├── configs/               # Factory configurations: actions, responses, state config objects for generation
│   ├── jumpstart/             # Jumpstart factory: initial project setup, example generation for quick start
│   ├── message/               # Message factory: generates message types, action kinds, parameters, responses. 5+ modules
│   ├── server/                # Server factory: generates server code, controller, trigger actions. 5+ modules
│   ├── state/                 # State factory: generates engine, CRUD, assemblers, tree, pools. 17+ writer modules
│   ├── testutils/             # Test utilities: declaration helpers, comparison functions for factory testing
│   ├── utils/                 # Factory utilities: code formatting, import management, common generation helpers
│   └── webclient/             # Webclient factory: generates TypeScript client, types, message handling. 13+ modules
├── marshallers/               # JSON marshaller generation: easyjson integration, import file creation for performance
├── module/                    # Go module utilities: finding module root, running go mod tidy operations
├── packages/                  # Package management: defines output packages, static code templates, language handlers
├── typescript/                # TypeScript code generation: AST-like helpers for generating type-safe TypeScript code
└── validator/                 # Configuration validation: 21+ validation rules, error messages, syntax checking
README.md                      # Project documentation: overview, installation, usage, API reference, architecture explanations
test/                          # End-to-end testing: mock controller, integration tests validating generated code behavior
tmp/                           # Temporary output directory: default target for generated code during development
```

## Utility Dependencies

**Note:** Before implementing custom utilities, check if functionality exists in these libraries.

### Code Generation
- **github.com/dave/jennifer** (v1.4.1) - Go code generation with type-safe AST building. Main factory tool for generating all Go output files.
- **github.com/mailru/easyjson** (v0.7.7) - Fast JSON marshaller/unmarshaller generation. Creates optimized JSON serialization for state objects.

### Data Handling
- **github.com/google/uuid** (v1.2.0) - UUID generation for entity IDs in state engine. Used in generated code for unique identifiers.
- **github.com/gertd/go-pluralize** (v0.1.7) - Pluralization utility. Converts singular type names to plural forms for collections, slices.

### WebSocket Communication
- **nhooyr.io/websocket** (v1.8.6) - WebSocket library for real-time client-server communication. Core transport layer for state broadcasting.

### Logging & Output
- **github.com/rs/zerolog** (v1.26.1) - Structured logging with zero allocations. Used in examples, server code for debug output.

### Testing & Comparison
- **github.com/stretchr/testify** (v1.7.0) - Test assertions and mock utilities. Used extensively across 100+ test files.
- **github.com/golang/mock** (v1.6.0) - Mock generation for interfaces in tests.
- **github.com/sergi/go-diff** (v1.2.0) - Text diffing for comparing expected vs actual generated code.
- **github.com/yudai/gojsondiff** (v1.0.0) - JSON diffing for state comparison in integration tests.

### Development Tools
- **golang.org/x/tools** (v0.1.7) - Go tools including imports formatter. Used for post-processing generated code formatting.
