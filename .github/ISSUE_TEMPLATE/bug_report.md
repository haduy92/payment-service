---
name: Bug Report
about: Create a report to help us improve
title: '[BUG] '
labels: ['bug', 'needs-triage']
assignees: ''

---

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Send request to '...'
2. With payload '...'
3. Expected response '...'
4. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Actual behavior**
A clear and concise description of what actually happened.

**API Request/Response**
```json
// Request
{
  "user_id": "example",
  "amount": 100.00,
  "transaction_id": "txn_123"
}

// Response
{
  "error": "example error message"
}
```

**Environment:**
 - OS: [e.g. Linux, Windows, macOS]
 - Go Version: [e.g. 1.23]
 - Service Version: [e.g. v1.0.0]
 - Docker: [Yes/No]

**Additional context**
Add any other context about the problem here.

**Logs**
```
Add relevant log output here
```
