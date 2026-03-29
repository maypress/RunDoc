
# RunDoc
**Executable documentation that never lies.**

RunDoc is a CLI tool that turns your Markdown documentation into executable specifications. It finds code blocks, runs them, and validates the output — ensuring your documentation is always truthful and up to date.

## The Problem
Documentation gets stale the moment you write it. Commands stop working, code examples have typos, and new team members waste hours figuring out which parts of the docs are actually correct.

## The Solution
RunDoc transforms your Markdown files into living documents that verify themselves. Add a simple annotation to any code block, and RunDoc will execute it, compare the output against your expectations, and tell you exactly what's broken.

## Features
- Multi-language support — Bash, Go, Python, and more with extensible runners
- Output validation — exact match, regular expressions, and exit code checking
- Windows and Unix support — works on any platform
- Verbose mode — detailed execution logs for debugging
- Update mode — automatically refresh expected outputs in your docs
- CI/CD ready — fails the build when documentation is out of date

## Installation

### From source
```bash run
git clone https://github.com/maypress/RunDoc.git
cd RunDoc
go build -o rundoc cmd/rundoc/main.go
```

### Using go install
```bash run
go install github.com/maypress/RunDoc/cmd/rundoc@latest
```

## Quick Start
Create a file named `example.md`:

```markdown
# My Documentation

## Check the version
```bash run
echo "Hello, World!"
# expect: Hello, World!
```
```

Now run it:
```bash run
rundoc example.md
```

**Output:**
```
📄 example.md
  ✓ echo "Hello, World!" (bash) — 2.3ms

📊 Result: 1 of 1 blocks passed
```

## Usage Guide

### Basic Commands
Check a documentation file:
```bash run
rundoc README.md
```

Run with verbose output to see execution details:
```bash run
rundoc --verbose README.md
```

Update expected outputs in your documentation:
```bash run
rundoc --update README.md
```

Combine flags:
```bash run
rundoc --update --verbose README.md
```

### Writing Documentation
RunDoc looks for code blocks with the `run` annotation. Add it right after the language name:

````markdown
```bash run
your-command-here
```
````

### Validating Output
Use comments inside your code block to tell RunDoc what to expect:

#### Exact Output Match
````markdown
```bash run
echo "Hello"
# expect: Hello
```
````

#### Regular Expression Match
````markdown
```python run
import datetime
print(datetime.datetime.now().year)
# expect-regex: \d{4}
```
````

#### Exit Code Validation
````markdown
```bash run
cat non-existent-file.txt
# expect-exit: 1
```
````

## Supported Languages
| Language | Annotation |
|----------|------------|
| Bash | `bash run` or `sh run` |
| Go | `go run` |
| Python | `python run` or `py run` |

## Adding New Languages
RunDoc uses a runner interface. To add a new language:
1. Create a new file in `internal/runner/extensions/`
2. Implement the `Run` method
3. Add your language to the switch in `runner.go`

## Example Workflow

### Step 1: Write documentation with tests
Create `docs/api.md`:

````markdown
# API Setup

## Start the server
```bash run
curl -s http://localhost:3000/health
# expect: OK
```
````

### Step 2: Run the validation
```bash run
rundoc docs/api.md
```

### Step 3: See what's broken
```
📄 docs/api.md
  ✗ Start the server (bash) — 1.2s
    💡 output mismatch:
    Expected:
    OK
    Got:
    Connection refused

📊 Result: 0 of 1 blocks passed
💡 Update documentation: rundoc docs/api.md --update
```

### Step 4: Fix your docs or your code
Either fix the actual API server, or update the documentation with the correct expected output:
```bash run
rundoc docs/api.md --update
```

## CI/CD Integration
RunDoc returns exit code 1 when any block fails, making it perfect for CI pipelines.

### GitHub Actions Example
```yaml
name: Validate Documentation

on:
  push:
    paths:
      - '**.md'
      - '**.go'

jobs:
  validate-docs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      - name: Install RunDoc
        run: go install github.com/maypress/RunDoc/cmd/rundoc@latest
      - name: Run Documentation Tests
        run: rundoc --verbose ./docs/*.md
```

### GitLab CI Example
```yaml
validate-docs:
  stage: test
  script:
    - go install github.com/maypress/RunDoc/cmd/rundoc@latest
    - rundoc --verbose docs/*.md
  only:
    changes:
      - "**/*.md"
```

## Command Reference
| Flag | Description |
|------|-------------|
| `--update` | Update expected outputs in the Markdown file |
| `--verbose` | Show detailed execution information |

## Project Structure
```
RunDoc/
├── cmd/
│   └── rundoc/
│       └── main.go           # CLI entry point
├── internal/
│   ├── parser/               # Markdown parser
│   │   └── parser.go
│   ├── runner/               # Code execution
│   │   ├── runner.go         # Runner interface
│   │   └── extensions/       # Language implementations
│   │       ├── bash.go
│   │       ├── go.go
│   │       └── python.go
│   ├── validator/            # Output validation
│   │   └── validator.go
│   └── reporter/             # Console output
│       └── reporter.go
├── testdata/                 # Test files
│   └── sample.md
└── go.mod
```

## Contributing
Contributions are welcome! Here's how to get started:
1. Fork the repository
2. Create a feature branch
3. Add your changes
4. Run tests: `go test ./...`
5. Submit a pull request

### Development Requirements
- Go 1.21 or higher
- Git

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements
RunDoc was inspired by tools like [cram](https://bitheap.org/cram/), [mdtest](https://github.com/rain-1/mdtest), and Python's [doctest](https://docs.python.org/3/library/doctest.html) — bringing the concept of executable documentation to modern multi-language projects.
```
