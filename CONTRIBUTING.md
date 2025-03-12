# Contributing

Welcome! 🎉 We're excited you're interested in contributing to this project. Please take a moment to review this document to make the process smooth for everyone.

## Table of Contents

- [Getting Started](#getting-started)
- [How to Contribute](#how-to-contribute)
- [Code Style](#code-style)
- [Commit Messages](#commit-messages)
- [Pull Requests](#pull-requests)
- [Reporting Issues](#reporting-issues)
- [Community](#community)

## Getting Started

1. **Fork** the repository.
2. **Clone** your fork:

   ```bash
   git clone https://github.com/issamoxix/Comit.git
   cd Comit
   ```

3. **Install dependencies** (if any):

   ```bash
   go mod tidy
   ```

4. **Run the tests** to make sure everything is working:

   ```bash
   go test ./...
   ```

## How to Contribute

- 💡 Fix bugs
- ✨ Add new features
- 🧪 Improve tests or add coverage
- 🧹 Refactor or improve existing code
- 📚 Improve documentation

If you're not sure where to start, check out the issues labeled [`good first issue`](https://github.com/issamoxix/Comit/issues).

## Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html).
- Use `go fmt`, `go vet`, and `golint` (if used in the project).
- Keep functions short and focused.
- Write clear, self-explanatory code.

To format your code automatically:

```bash
go fmt ./...
```

## Commit Messages

use the Comit
```
comit

# fix(parser): handle nil pointer in token parser

```

Types include: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`.

## Pull Requests

1. Make sure all tests pass.
2. Make sure your changes are linted and formatted.
3. Keep PRs focused and small.
4. Link to any relevant issues.
5. Add tests for new functionality if possible.

## Reporting Issues

When reporting a bug, please include:

- A clear description
- Steps to reproduce
- Expected vs actual behavior
- Relevant logs or error messages
- Your Go version and OS

## Community

Be kind, respectful, and constructive. We follow the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md).

---

Thanks for helping make this project better! 🎈
```

---

Let me know if you'd like to add sections for GitHub Actions, specific tools like `golangci-lint`, or examples of tests!
