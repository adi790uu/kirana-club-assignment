# Kirana Club Assignment

## Description

This project is a job processing system designed to handle image processing tasks. It is built using the Go programming language and leverages the Fiber web framework for handling HTTP requests. The system processes jobs submitted via an API, stores job and image data in a PostgreSQL database, and uses a worker to process images asynchronously.

## Installation and Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/kirana-club-assignment.git
   cd kirana-club-assignment
   ```

2. **Set up environment variables:**
   Create a `.env` file in the root directory with the following variables:

   ```plaintext
   DB_HOST=your_db_host
   DB_USER=your_db_user
   DB_PASS=your_db_password
   DB_NAME=your_db_name
   DB_PORT=your_db_port
   DB_SSLMODE=disable
   PORT=8080
   ```

3. **Install dependencies:**
   Ensure you have Go installed. You can download it from [the official Go website](https://golang.org/dl/).

4. **Build the project:**

   ```bash
   go build
   ```

5. **Run the project:**

   - Using the built binary:
     ```bash
     ./kirana-club-assignment
     ```
   - Or directly using Go:
     ```bash
     go run main.go
     ```

6. **Run with Docker:**
   - Build the Docker image:
     ```bash
     docker build -t kirana-club-assignment .
     ```
   - Run the Docker container:
     ```bash
     docker run --env-file .env -p 8080:8080 kirana-club-assignment
     ```

## Work Environment

- **Operating System:** Developed and tested on macOS/Linux. Windows users may need to adjust scripts.
- **Text Editor/IDE:** Visual Studio Code with Go extensions.
- **Libraries:** The project uses the Fiber web framework and GORM for ORM. Dependencies are managed via `go.mod`.

## Future Improvements

Given more time, the following improvements could be made:

- Implement a more robust error handling mechanism.
- Add more comprehensive unit and integration tests.
- Utilize a message broker such as Kafka/RabbitMQ to handle job processing.
- Creating a worker pool to handle the job processing for better scalability.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
