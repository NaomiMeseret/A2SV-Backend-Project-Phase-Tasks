# Task Manager API Documentation (Beginner Version)

This is a simple guide for using the Task Manager API. It explains how to register, log in, use tokens, and what each endpoint does. Everything is organized using Clean Architecture.

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
    "username": "Naomi",
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
    "username": "Naomi",
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
