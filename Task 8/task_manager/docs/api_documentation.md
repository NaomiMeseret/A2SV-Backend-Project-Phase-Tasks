# Task Manager API Documentation 

This is a simple guide for using the Task Manager API. It explains how to register, log in, use tokens, and what each endpoint does. Everything is organized using Clean Architecture and includes  testing.

## Folder Structure (Clean Architecture)

```
task-manager/
├── Delivery/        # Handles web requests (main.go, controllers, routers)
├── Domain/          # Main business objects (Task, User, interfaces)
├── Infrastructure/  # Helpers (passwords, JWT, auth middleware)
├── Repositories/    # Database code
├── Usecases/        # App logic (register, login, tasks)
└── docs/            # Documentation
```


### How Clean Architecture Works Here

- **Delivery**: Handles HTTP requests and responses. It does not contain business logic.
- **Usecases**: Contains all the main app logic (like registering, logging in, creating tasks). This is where business rules live.
- **Domain**: Defines the main objects (like User and Task) and interfaces. These are just plain Go structs, no JSON tags or framework stuff.
- **Repositories**: Handles all database operations. The usecase talks to the repository using interfaces, not direct DB code.
- **Infrastructure**: Has helpers for things like password hashing and JWT tokens. These are used by the usecase or delivery layers.
- **Why?** This keeps your code organized, easy to test, and easy to change later. Each part has one job.

## MongoDB Setup

- Make sure MongoDB is running on your computer.
- Default connection: `mongodb://localhost:27017`
- Database: `task_manager`
- Collections: `users`, `tasks`

## How to Run the Server

1. Open a terminal in your project folder.
2. Run:
   ```bash
   go run Delivery/main.go
   ```
3. The server will start at: `http://localhost:8080`

## User Registration

- **POST /register**
- Anyone can register. The first user will be an admin. Others are regular users.
- **Body (JSON):**
  ```json
  {
    "email": "naomi@gmail.com",
    "password": "yourpassword"
  }
  ```
- **In Postman:**
  1. Set method to POST and URL to `http://localhost:8080/register`
  2. Go to Body > raw > JSON, paste the example above.
  3. Click Send.
- **Response:**
  - Success: `{ "message": "User registered successfully" }`
  - Error: `{ "error": "..." }`

## User Login

- **POST /login**
- Log in with your username and password to get a JWT token.
- **Body (JSON):**
  ```json
  {
    "email": "naomi@gmail.com",
    "password": "yourpassword"
  }
  ```
- **In Postman:**
  1. Set method to POST and URL to `http://localhost:8080/login`
  2. Go to Body > raw > JSON, paste the example above.
  3. Click Send.
- **Response:**
  - Success: `{ "token": "..." }` (copy this token for protected endpoints)
  - Error: `{ "error": "..." }`

## Using the Token (Authentication)

- For all protected endpoints, add a header:
  - **Key:** `Authorization`
  - **Value:** `Bearer <your_token_here>`
- If you don’t include a valid token, you’ll get a 401 Unauthorized error.

## User Roles

- There are two roles: `admin` and `user`.
- The first user is always an admin.
- Only admins can create, update, or delete tasks, promote users, or list all users.
- All users (including admins) can view tasks.

## Endpoints

### Register

- **POST /register**
- Anyone can use this to create an account.

### Login

- **POST /login**
- Anyone can use this to log in and get a token.

### Create Task (Admin Only)

- **POST /tasks**
- Only admins can create tasks.
- **Body (JSON):**
  ```json
  {
    "title": "Task title",
    "description": "Task details",
    "date": "2025-07-20",
    "status": "pending"
  }
  ```

### Update Task (Admin Only)

- **PUT /tasks/:id**
- Only admins can update tasks.
- **Body (JSON):**
  ```json
  {
    "title": "New title",
    "description": "New details",
    "date": "2025-07-21",
    "status": "completed"
  }
  ```

### Delete Task (Admin Only)

- **DELETE /tasks/:id**
- Only admins can delete tasks.

### Get All Tasks

- **GET /tasks**
- Any logged-in user can see all tasks.

### Get Task by ID

- **GET /tasks/:id**
- Any logged-in user can see a task by its ID.

### Promote User to Admin (Admin Only)

- **PUT /promote/:id**
- Only admins can promote another user to admin.
- No body needed.

### List All Users (Admin Only)

- **GET /users**
- Only admins can see all users (no passwords shown).

## Error Responses

- **401 Unauthorized:** If you don’t provide a valid token.

```json
{ "error": "Unauthorized" }
```

- **403 Forbidden:** If you try to do admin actions as a regular user.
  ```json
  { "error": "Forbidden" }
  ```
- **400/404:** For bad input or not found.

##  Unit Testing for Task Managemnet API

This project uses the testify library to test the code and make sure everything works.

- We wrote tests for all the important parts: domain, usecases, and controllers.
- We use "mocks" (fake versions) so tests don't need a real database.
- Each test starts fresh, so tests don't mess each other up.
- We check that the results are what we expect using "assertions" (test checks).

This project uses the [testify](https://github.com/stretchr/testify) library for comprehensive unit testing.

### Running Tests

```bash
go test ./... -v
```

This will run all the tests and show you what passed or failed.

### How to See Test Coverage

To see how much of your code is tested, type:

```bash
go test ./... -cover
```

### Continuous Integration (CI)

To automate testing on every commit, add a workflow file (e.g., `.github/workflows/go.yml`) with:

```yaml
name: Go Tests
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: "1.20"
      - name: Run tests
        run: go test ./... -v -cover
```

### Notes

- All critical components (domain, usecases, controllers) are covered with unit tests.
- Mocks are used to isolate dependencies and ensure test independence.
- Each test suite uses setup/teardown for a clean state.
- Assertions verify expected outcomes and edge cases.

