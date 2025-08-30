# Contributing to This Go Library

### 📦 Versioning

We follow [Semantic Versioning](https://semver.org/) (`MAJOR.MINOR.PATCH`) for this project. Every release is marked with a Git tag, and these tags are used by Go Modules to resolve versions.

- **MAJOR**: Increment when making backward-incompatible changes.
  - For `v2+`, update the module path in `go.mod` (e.g., `module github.com/yourname/yourlib/v2`) and in all import statements.
- **MINOR**: Increment when adding functionality in a backward-compatible manner.
- **PATCH**: Increment for backward-compatible bug fixes or small improvements.

#### Release Process
1. Commit all changes and ensure tests pass.
2. Create a version tag:
   ```bash
   git tag vX.Y.Z
   git push origin vX.Y.Z
   ```
