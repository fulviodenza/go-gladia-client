# Go Gladia Client

This project is a Go client for the Gladia.io transcription service. It provides an easy-to-use interface for interacting with the transcription API, allowing users to send audio files for transcription and receive the results.

## Project Structure

```
go-gladia-client
├── pkg
│   ├── gladia
│   │   ├── client.go      # Gladia client structure and methods
│   │   ├── models.go      # Data models for transcription requests and responses
│   │   └── transcription.go # Functions for sending transcription requests
│   └── errors
│       └── errors.go      # Custom error types and handling functions
├── go.mod                 # Module definition and dependencies
├── go.sum                 # Checksums for module dependencies
└── README.md              # Project documentation
```

An example of how to use this library is here:
notion-echo bot: https://github.com/fulviodenza/notion-echo/blob/main/adapters/gladia/gladia.go
