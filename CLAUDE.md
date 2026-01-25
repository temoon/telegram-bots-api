# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Go implementation of the Telegram Bots API (currently version 9.3.0), providing a type-safe client library for building Telegram bots. The implementation closely follows the official Telegram Bot API documentation at https://core.telegram.org/bots/api.

## Development Commands

### Build and Test
```bash
# Build the package
go build

# Run all tests
go test ./...

# Run tests for a specific package
go test ./requests

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

### Module Management
```bash
# Download dependencies
go mod download

# Verify dependencies
go mod verify

# Tidy up go.mod and go.sum
go mod tidy
```

## Architecture

### Core Components

**`bot.go`** - Central HTTP client and request dispatcher
- `Bot` struct encapsulates HTTP client and configuration
- `BotOpts` supports production/test environments via `Env` field (use `EnvProduction` or `EnvTest`)
- `CallMethod()` handles all API calls, automatically switching between form-urlencoded and multipart/form-data based on file presence
- Request/response cycle: serializes request → HTTP POST → deserializes Telegram's JSON response

**`types.go`** - Complete type definitions (~3000+ lines)
- Exhaustive Go structs for all Telegram API types
- All types follow Telegram's exact field naming (JSON tags use snake_case)
- Pointer fields (`*string`, `*int64`, `*bool`) represent optional parameters

**`constants.go`** - API constants
- Parse modes: `ParseModeMarkdown`, `ParseModeMarkdownV2`, `ParseModeHTML`
- Chat types, message entity types, chat member statuses, etc.
- Always use these constants instead of string literals

**`requests/` directory** - One file per API method (~150+ methods)
- Each request file implements the `Request` interface with three methods:
  - `Call(ctx, bot)` - executes the request and returns typed response
  - `GetValues()` - serializes struct fields to `map[string]string` for HTTP form data
  - `GetFiles()` - returns file attachments for multipart uploads (or nil)
- Request structs have fields matching Telegram API parameters exactly

### Custom Types

**`ChatId`** (type_chat_id.go)
- Unified type handling both numeric IDs (`int64`) and usernames (`string` starting with `@`)
- Use `NewChatId(id, name)` to construct
- Automatically serializes correctly for API calls

**`InputFile`** (type_input_file.go)
- Handles three file input methods: file_id (string), HTTP URL, or file upload
- For uploads, use `NewInputFile("", reader, fileName)`
- For existing files, use `NewInputFile(fileId, nil, "")`
- Implements `attach://` protocol for multipart uploads automatically

### Request Pattern

All requests follow this pattern:
```go
type SomeRequest struct {
    RequiredField   string
    OptionalField   *string  // pointer = optional
}

func (r *SomeRequest) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
    response = new(telegram.SomeResponseType)
    err = b.CallMethod(ctx, "someMethod", r, response)
    return
}

func (r *SomeRequest) GetValues() (values map[string]string, err error) {
    values = make(map[string]string)
    values["required_field"] = r.RequiredField
    if r.OptionalField != nil {
        values["optional_field"] = *r.OptionalField
    }
    // Booleans: "1" for true, "0" for false
    // Complex types: JSON-marshal to string
    return
}

func (r *SomeRequest) GetFiles() (files map[string]io.Reader) {
    // Return nil if no files, or map[fieldName]reader for uploads
    return
}
```

### Key Implementation Details

1. **Boolean serialization**: Telegram API expects booleans as `"1"` (true) or `"0"` (false) strings
2. **Complex parameters**: Objects/arrays are JSON-marshaled to strings in `GetValues()`
3. **File handling**: Files in `GetFiles()` automatically trigger multipart/form-data encoding
4. **Error handling**: API errors are returned as Go errors with Telegram's description message
5. **Context support**: All requests accept `context.Context` for cancellation/timeout control

## Code Generation

**CRITICAL**: Most of this codebase is automatically generated from the official Telegram Bot API documentation using https://github.com/temoon/telegram-bots-api-generator

**DO NOT** manually edit the following files - all changes will be overwritten on next generation:
- `types.go` - all type definitions
- `requests/*.go` - all API method implementations

**Files safe to modify manually:**
- `bot.go` - core bot implementation (NOT generated)
- `constants.go` - API constants (NOT generated)
- `type_chat_id.go` - custom ChatId type
- `type_input_file.go` - custom InputFile type
- `go.mod` / `go.sum` - dependency management
- `NOTICE` / `LICENSE` - copyright and licensing

**To update when Telegram releases new API versions:**
1. See documentation at https://github.com/temoon/telegram-bots-api-generator
2. Run the generator against the updated Telegram API documentation
3. The generator will regenerate `types.go` and `requests/*.go`
4. Manually update `constants.go` with any new constants if needed
5. Update version in `constants.go` to match new Telegram Bot API version

## Version Updates

The library version in `constants.go` reflects the Telegram Bot API version. Current version: 9.3.0

When Telegram updates their API, use the generator to create a new version rather than making manual changes.
