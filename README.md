# SSS Wallet Core

A distributed cryptographic wallet system using Shamir's Secret Sharing (SSS) to securely manage private keys. The system splits a private key into multiple shares and can recover it using only a subset of shares.

## Features

- **Distributed Key Management**: Split private key into 5 shares (recoverable with 3)
- **Secure Signing**: Sign messages using randomly selected shares
- **Signature Verification**: ECDSA signature validation
- **REST API**: Simple HTTP-based API

## Quick Start

```bash
# Install dependencies
go mod download

# Run server
go run cmd/server/main.go
```

Server runs on `http://localhost:8080`

## API Endpoints

### GET `/wallet`
Get wallet public key information.

**Response:**
```json
{
  "x": "1234567890abcdef...",
  "y": "fedcba0987654321..."
}
```

### POST `/sign`
Sign a message using distributed shares.

**Request:**
```json
{
  "message": "Hello, World!"
}
```

**Response:**
```json
{
  "r": "a1b2c3d4e5f6...",
  "s": "f6e5d4c3b2a1..."
}
```

### POST `/verify`
Verify a signature.

**Request:**
```json
{
  "message": "Hello, World!",
  "r": "a1b2c3d4e5f6...",
  "s": "f6e5d4c3b2a1..."
}
```

**Response:**
```json
{
  "valid": true
}
```

## Security Features

- **Shamir's Secret Sharing**: 5 shares with threshold 3
- **ECDSA**: P-256 curve with SHA-256 hashing
- **Random Selection**: Different share combinations for each signature
- **Fault Tolerance**: System remains secure even if some shares are lost

## Architecture

```
sss-wallet-core/
├── cmd/server/main.go    # Application entry point
├── core/                 # Core cryptographic logic
│   ├── keygen.go        # Key generation & splitting
│   ├── sss.go           # Shamir reconstruction
│   ├── sign.go          # ECDSA signing
│   └── verify.go        # Signature verification
└── handlers/            # HTTP request handlers
    ├── wallet.go        # Wallet info handler
    ├── sign.go          # Signing handler
    └── verify.go        # Verification handler
```

## Requirements

- Go 1.24.3+
- Hashicorp Vault library

## Use Cases

- Blockchain wallet key management
- Multi-signature systems
- Secure key backup and recovery
- Distributed systems with fault tolerance

## License

MIT 