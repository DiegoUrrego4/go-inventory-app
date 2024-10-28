# Go inventory app

This project is a market inventory management application with user login and JWT authentication. Built with Go using hexagonal architecture, PostgreSQL as the database, and Docker for containerization, it supports core inventory functionalities and user authentication.

### Features
- User Authentication: JWT-based login system.
- Inventory Management: CRUD (initially) operations for managing inventory items.
- Database Persistence: PostgreSQL as the primary database.
- Dockerized Setup: Quick setup with Docker and docker-compose.

### Development
1. Download Go modules. Run the following command to download dependencies:

```
go mod tidy
```

2. Configure environment: Copy the ```.env-template``` file and rename it to ```.env``` Update it with your specific environment variables if necessary.
3. Run the database: Start the PostgreSQL database container:
```
docker-compose up -d
```
4. Run the application:
```
go run cmd/server/main.go
```
5. Verify connection: Send an HTTP request to the ```/ping``` endpoint on the assigned port to verify the setup.

## Contributing

Feel free to open issues or submit pull requests to enhance this project.

## License

This project is licensed under the MIT License.
