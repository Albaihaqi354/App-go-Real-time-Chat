# App Go Real-time Chat

A real-time chat application built using Go and WebSocket for seamless and efficient communication. This application is designed to provide real-time messaging with a focus on performance and scalability.

## Features

- **Real-time Communication**: Leveraging WebSocket for instant messaging.
- **Scalability**: Designed to handle multiple users and conversations efficiently.
- **Simple and Clean Architecture**: Easy to understand and extend.
- **User Authentication**: Secure authentication for users.
- **Message Persistence**: Messages are stored for later retrieval.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) 1.16 or later
- [MySQL](https://www.mysql.com/downloads/) or [PostgreSQL](https://www.postgresql.org/download/)
- [Node.js](https://nodejs.org/en/download/) (optional, if you plan to work on the frontend)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Albaihaqi354/App-go-Real-time-Cha.git
    cd App-go-Real-time-Cha
    ```

2. Set up the database:

    - Create a new MySQL or PostgreSQL database.
    - Update the database configuration in `config/database.go`.

3. Install Go dependencies:

    ```bash
    go mod tidy
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

5. Open your browser and go to `http://localhost:8080` to see the application in action.

### Configuration

- **Database**: Configure your database connection in `config/database.go`.
- **Port**: Change the application port in `main.go` if needed.

## Usage

1. Register a new user.
2. Log in with your credentials.
3. Start chatting in real-time with other users.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

If you have any questions, feel free to reach out to the repository owner [Bian Albaihaqi](https://github.com/Albaihaqi354).
