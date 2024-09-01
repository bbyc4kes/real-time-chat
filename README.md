# Real-time Chat - Golang w/ Gin, Gorilla 

![Go](https://img.shields.io/badge/Go-1.23-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)

## Project Overview

This is the project in golang  that is designed to handle real-time communication via WebSockets.

## Project Structure

### Server

The project structured as follows:

```
├── cmd/
│   └── main.go                 # Main entry point for the server
├── db/
│   ├── migrations/             # Database migration files
│   └── db.go                   # Database connection and setup
├── internal/
│   ├── user/                   # User domain logic
│   │   ├── user_handler.go     # HTTP handlers for user routes
│   │   ├── user_repository.go  # User data access layer
│   │   ├── user_service.go     # User business logic
│   │   └── user.go             # User model
│   └── ws/                     # WebSocket logic
│       ├── client.go           # WebSocket client management
│       ├── hub.go              # WebSocket hub for broadcasting messages
│       └── ws_handler.go       # WebSocket HTTP handler
├── router/
│   └── router.go               # HTTP router setup
├── util/
│   └── password.go             # Password hashing utilities
├── .env                        # Environment variables
├── go.mod                      # Go module file
└── go.sum                      # Go dependencies checksum
```

#### How to Run The Project

To start the server, navigate to the `cmd/` directory and run:

```bash
go run main.go
```

Backend server runs on `http://localhost:8080/`.

Ensure your `.env` file is properly configured with the necessary environment variables.

## Hub architecture

![image](https://github.com/user-attachments/assets/811629b9-3386-4420-9b92-2d6767c3f1bf)

Each client has a `writeMessage` and a `readMessage` methods. The `readMessage` method reads incoming messages from the client's WebSocket connection and forwards them to the hub's Broadcast channel. The hub then broadcasts these messages to all clients within the same room. The `writeMessage` method for each client sends messages through its WebSocket connection, which the frontend handles to display the messages appropriately.

## Key Features

- **Golang WebSocket Server:** Real-time communication support using WebSockets.

## Environment Setup

Ensure Go is installed and set up on your system. Set up environment variables in the `.env` file based on `.env.example`.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repo and create your branch from `main`.
2. Follow naming as **docs: removed spaces**/**feat: establish websocket connection**/**fix: client connection error*** in imperative mood. 
3. If you've added code that should be tested, add tests.
4. Submit a pull request!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Golang](https://golang.org/)

---

Feel free to raise any issues or suggest improvements. Enjoy coding!
