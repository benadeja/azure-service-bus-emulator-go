![Demo](https://img.shields.io/badge/demo-working-brightgreen)
![Local Only](https://img.shields.io/badge/cloud-none-success)
![Zero Cost](https://img.shields.io/badge/cost-%240.00-blue)

# Azure Service Bus Emulator with Testcontainers-Go 🐳✨

[![Go](https://img.shields.io/badge/Go-1.22%2B-blue?logo=go)](https://go.dev)
[![Testcontainers](https://img.shields.io/badge/Testcontainers-v0.33.0-0066cc?logo=docker)](https://testcontainers.com)
[![Azure Service Bus](https://img.shields.io/badge/Azure_Service_Bus_Emulator-1.1.2-0078D4?logo=microsoft-azure)](https://learn.microsoft.com/en-us/azure/service-bus-messaging/)

**Run a full Azure Service Bus namespace locally — no Azure account required.**  
Perfect for integration tests, local development, CI/CD, and learning.

This project demonstrates how to run the **official Microsoft Azure Service Bus Emulator** using **Testcontainers-Go**, with a pre-configured queue and working send/receive in pure Go.

Zero cloud. Zero cost. Zero hassle.

## Features

- Fully local Service Bus namespace (`sbemulator`)
- Pre-created queue: `queue.1`
- Automatic SQL Server + emulator setup via Testcontainers
- Embedded config with proper logging (fixes common startup errors)
- Real send/receive using the official `azservicebus` SDK

## Prerequisites

- Docker (Docker Desktop on macOS/Windows, Docker Engine on Linux)
- Go 1.22 or later

## Quick Start

```bash
git clone https://github.com/yourusername/azure-servicebus-emulator-testcontainers-go.git
cd azure-servicebus-emulator-testcontainers-go

go run .
```

## Expected Output

```bash
Connected to emulator: Endpoint=sb://localhost;SharedAccessKeyName=RootManageSharedAccessKey;...
Message sent
Received 1 message(s)
Body: Hello message sent!
Demo completed successfully!
```

## Project Structure

```bash
.
├── send-and-receive-from-asb-emulator.go    → Complete working demo
├── servicebus-config.json                   → Embedded emulator configuration
├── go.mod                                   → Dependencies
└── README.md                               
```

## Why This Exists

The official Azure Service Bus Emulator is powerful but tricky to run locally:
- Requires SQL Server
- Complex config JSON
- Logging fields must have Level or it crashes
- Hard to wire up networking

This repo solves all of that in <100 lines of clean Go using Testcontainers.Use it as:
- A template for integration tests
- A local dev environment
- A CI-friendly Service Bus mock

## Use in Tests
Just copy main.go into your test suite and wrap it:

```go
func TestIntegration_WithServiceBus(t *testing.T) {
    client := startEmulator(t) // returns *azservicebus.Client
    // Your real integration tests here
}
```

## Connection String
The emulator uses the development connection string:
```bash
Endpoint=sb://localhost;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;
```
No need to manage keys — it's all local and safe.

## License
MIT © [Jaco Benadé]






