# go-expense-tracker

This API allows users to manage their expenses with features like user authentication, JWT-based session management, and expense tracking. Users can create, read, update, and delete their expenses, with options to filter by time periods, including the past week, month, or custom date ranges. Built with Go, it uses a flexible data model to categorize expenses.

## Features

- **User Authentication**: Secure user authentication with JWT-based session management.
- **Expense Management**: Create, read, update, and delete expenses.
- **Filtering**: Filter expenses by time periods, including past week, past month, last 3 months, or custom date ranges.
- **Categorization**: Categorize expenses for better organization and tracking.

## Project Link

This project is one of the backend projects listed on the [roadmap.sh](https://roadmap.sh) website. You can find the project link [here](https://roadmap.sh/projects/expense-tracker-api).

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/alielmi98/go-expense-tracker.git
    cd go-expense-tracker
    ```

2. Build and run the Docker containers:

    ```sh
    docker-compose -f docker/docker-compose.yml up --build
    ```

3. The backend service will be available at `http://localhost:5000`.


### API Endpoints

#### Authentication

- **POST /v1/auth/register**: Register a new user
- **POST /v1/auth/login**: Login and obtain a JWT token

#### Expenses

- **POST /v1/expense/**: Create a new expense
- **GET /v1/expense/{id}**: Get an expense by ID
- **PUT /v1/expense/{id}**: Update an expense by ID
- **DELETE /v1/expense/{id}**: Delete an expense by ID
- **GET /v1/expense/**: List and filter expenses

### Example Requests

#### Create Expense

```sh
curl -X POST http://localhost:8080/v1/expense/ \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_AUTH_TOKEN" \
-d '{
    "title": "Dinner",
    "amount": 50.75,
    "category": "Food",
    "date": "2025-02-10T19:30:00Z"
}'
```

#### List Expenses

```sh
curl -X GET http://localhost:8080/v1/expense/ \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_AUTH_TOKEN" \
-d '{
    "filter": "past_week"
}'
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.