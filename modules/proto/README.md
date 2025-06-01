# Proto Module

This module contains all Protocol Buffer definitions for the forum project.

## Structure

```
proto/
├── auth/           # Auth service proto files
├── forum/          # Forum service proto files
├── chat/           # Chat service proto files
└── common/         # Common proto definitions
```

## Usage

To use these proto files in your service:

1. Import the module:
```go
import "github.com/yourusername/forum-modules/proto"
```

2. Use the generated code in your service.

## Generation

To generate Go code from proto files:

```bash
make proto
```

This will generate all necessary Go files in the appropriate directories. 