![image](https://github.com/user-attachments/assets/0695d85f-e5af-4343-9d13-3e5610422a23)# Real-time Chat - Golang w/ Gin, Gorilla & TypeScript w/ Next.js

![Go](https://img.shields.io/badge/Go-1.23-blue.svg)
![Next.js](https://img.shields.io/badge/Next.js-14.2-blue.svg)
![TailwindCSS](https://img.shields.io/badge/TailwindCSS-3.4.10-green.svg)
![TypeScript](https://img.shields.io/badge/TypeScript-5.5.4-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)

## Project Overview

This project is a full-stack web application that uses Golang on the backend for WebSocket communication and Next.js with Tailwind CSS and TypeScript on the frontend. The backend is designed to handle real-time communication via WebSockets, while the frontend provides a modern, responsive user interface.

## Project Structure

### Server (`server/`)

The backend of the project is a Golang server structured as follows:

```
server/
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

#### How to Run the Backend

To start the backend server, navigate to the `server/cmd/` directory and run:

```bash
go run main.go
```

Backend server runs on `http://localhost:8080/`.

Ensure your `.env` file is properly configured with the necessary environment variables.

### Client (`client/`)

The frontend of the project is built with Next.js, using the pages router:

```
client/
├── modules/                    # Modules directory
├── constants/                  # Constants directory
├── components/                 # Components directory
├── pages/                      # Next.js pages routing
│   ├── index.tsx               # Main entry point
│   └── ...                     # Other pages and components
├── public/                     # Static files
├── styles/                     # Global styles, including TailwindCSS
├── tsconfig.json               # TypeScript configuration
├── tailwind.config..ts               # Tailwind configuration
└── package.json                # NPM package configuration
```

#### How to Run the Frontend

Navigate to the `client/` directory and install dependencies:

```bash
npm install
```
or
```bash
pnpm install
```

To start the development server:

```bash
npm run dev
```
or
```bash
pnpm dev
```

This will start the Next.js application, typically running at `http://localhost:3000`.

## Hub architecture

![image](https://github.com/user-attachments/assets/811629b9-3386-4420-9b92-2d6767c3f1bf)

Each client has a `writeMessage` and a `readMessage` methods. The `readMessage` method reads incoming messages from the client's WebSocket connection and forwards them to the hub's Broadcast channel. The hub then broadcasts these messages to all clients within the same room. The `writeMessage` method for each client sends messages through its WebSocket connection, which the frontend handles to display the messages appropriately.

## Key Features

- **Golang WebSocket Server:** Real-time communication support using WebSockets.
- **Next.js Frontend:** Modern, SSR-capable frontend with pages router.
- **Tailwind CSS:** For rapid UI development with utility-first CSS.
- **TypeScript:** Type-safe codebase on both client and server.

## Environment Setup

1. **Backend:** Ensure Go is installed and set up on your system. Set up environment variables in the `.env` file based on `.env.example`.
2. **Frontend:** Ensure Node.js and npm are installed. Configure environment variables in `.env.local` based on `.env.example`.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. Ensure the code passes linting with `golangci-lint` and `eslint`.
4. Submit a pull request!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Golang](https://golang.org/)
- [Next.js](https://nextjs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [TypeScript](https://www.typescriptlang.org/)

---

Feel free to raise any issues or suggest improvements. Enjoy coding!
