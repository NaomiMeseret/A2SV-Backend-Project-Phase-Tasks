# Task Manger API Documentation

-> This document shows how to use the Task Manger API in clear
It describes:

1. Folder structure
2. How to configure MongoDB
3. How to register and login as a user
4. How to use authentication (JWT tokens)
5. Each "HTTP endpoint"
6. Which file handles the logic

## Folder Structure

```plaintext
task_manager/
├── main.go                  => Starts the app by calling the router setup
├── controllers/
│   └── task_controller.go   => Handles HTTP requests and calls the service function
├── models/
│   └── task.go              => Defines the Task structure (fields and JSON tags)
├── data/
│   └── task_service.go      => Connects to MongoDB and implements CRUD operations
├── router/
│   └── router.go            => Connects each HTTP path and method to the correct controller
├── docs/
│   └── api_documentation.md => You are here
└── go.mod
```

# MongoDB Configuration

Use MongoDB for storing tasks
**Connection URI**

- **Local MongoDB(default):**

  `mongodb://localhost:27017`

**Database & Collection**

- **Database:** `taskdb`
- **Collection:** `tasks`

# Running the server

1. Open a terminal in a project folder
2. Run:

```bash
go run main.go
```

3. Base URL

`http://localhost:8080`

# Authentication & User Accounts

## 1. Register a New User

- **Endpoint:** `POST /register`
- **Description:** Create a new user account. The first user ever created will become an admin. All others are regular users.
- **Request Body Example:**

```json
{
  "username": "yourusername",
  "password": "yourpassword"
}
```

- **How to do it in Postman:**
  1. Set method to POST and URL to `http://localhost:8080/register`
  2. Go to the Body tab, select raw and JSON, and paste the example above.
  3. Click Send.
- **Success Response:**
  - 201 Created
  - Returns your user id, username, and role (admin or user)

## 2. Login to Get a Token

- **Endpoint:** `POST /login`
- **Description:** Log in with your username and password to get a JWT token. You need this token to access protected endpoints.
- **Request Body Example:**

```json
{
  "username": "yourusername",
  "password": "yourpassword"
}
```

- **How to do it in Postman:**
  1. Set method to POST and URL to `http://localhost:8080/login`
  2. Go to the Body tab, select raw and JSON, and paste the example above.
  3. Click Send.
- **Success Response:**
  - 200 OK
  - Returns `{ "token": "..." }` (copy this token for the next steps)

## 3. Using the Token (Authentication)

- For all protected endpoints (most of them), you must add an **Authorization** header:
  - **Key:** `Authorization`
  - **Value:** `Bearer <your_token_here>`
- In Postman, go to the Headers tab and add:
  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...`
- If you do not include a valid token, you will get a 401 Unauthorized error.

## 4. User Roles and Permissions

- There are two roles: `admin` and `user`.
- **Admins** can:
  - Create, update, and delete tasks
  - Promote other users to admin
  - Do everything regular users can do
- **Regular users** can:
  - View all tasks
  - View a task by ID
- The first user ever registered is automatically an admin.
- Only admins can access certain endpoints (see below).

## 5. Protected Endpoints

- The following endpoints require you to be logged in (token in Authorization header):
  - `GET /tasks` (all users)
  - `GET /tasks/:id` (all users)
  - `POST /tasks` (admin only)
  - `PUT /tasks/:id` (admin only)
  - `DELETE /tasks/:id` (admin only)
  - `PUT /promote/:id` (admin only)
- If you try to access these without a valid token, you will get a 401 error.
- If you try to do admin actions as a regular user, you will get a 403 Forbidden error.

## 6. Promote a User to Admin (Admin Only)

- **Endpoint:** `PUT /promote/:id`
- **Description:** Admins can promote another user to admin by their user ID.
- **How to do it in Postman:**
  1. Log in as an admin and copy your token.
  2. Set method to PUT and URL to `http://localhost:8080/promote/<user_id>`
  3. In the Headers tab, add your Authorization header.
  4. Click Send. No body is needed.
- **Success Response:**
  - 200 OK
  - `{ "message": "Promoted to admin" }`

# API Endpoints

-> All URLS start with http://localhost:8080

## User Registration

- **Method:** POST
- **URL:** `/register`
- **Description:** Create a new user account. The first user ever created will become an admin. All others are regular users.
- **Body:**

```json
{
  "username": "yourusername",
  "password": "yourpassword"
}
```

- **Success Response:**
  - 201 Created
  - Returns user id, username, and role
- **Example Request:**
  `POST http://localhost:8080/register`
- **Example Response:**

```json
{
  "id": "664f1a2b3c4d5e6f7a8b9c0d",
  "username": "yourusername",
  "role": "user"
}
```

## User Login

- **Method:** POST
- **URL:** `/login`
- **Description:** Log in and get a JWT token. You need this token for protected endpoints.
- **Body:**

```json
{
  "username": "yourusername",
  "password": "yourpassword"
}
```

- **Success Response:**
  - 200 OK
  - Returns `{ "token": "..." }`
- **Example Request:**
  `POST http://localhost:8080/login`
- **Example Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```

## List All Tasks

- **Method:** GET
- **URL:** `/tasks`
- **Description:** Get all tasks.
- **Requires:** JWT token (any user)
- **Example Request:**
  `GET http://localhost:8080/tasks`
- **Headers:**
  `Authorization: Bearer <your_token>`
- **Example Response:**

```json
[
  {
    "id": "68790b1719fc9d12d16e66d0",
    "title": "Learn MongoDB",
    "description": "Update Description",
    "date": "2025-07-20",
    "status": "In progress"
  }
]
```

## Get Task by ID

- **Method:** GET
- **URL:** `/tasks/:id`
- **Description:** Get a task by its ID.
- **Requires:** JWT token (any user)
- **Example Request:**
  `GET http://localhost:8080/tasks/68790b1719fc9d12d16e66d0`
- **Headers:**
  `Authorization: Bearer <your_token>`
- **Example Response:**

```json
{
  "id": "68790b1719fc9d12d16e66d0",
  "title": "Learn MongoDB",
  "description": "Update Description",
  "date": "2025-07-20",
  "status": "In progress"
}
```

## Create a New Task

- **Method:** POST
- **URL:** `/tasks`
- **Description:** Create a new task.
- **Requires:** JWT token (admin only)
- **Headers:**
  `Authorization: Bearer <your_token>`
  `Content-Type: application/json`
- **Body:**

```json
{
  "title": "Solve Leetcode Problem",
  "description": "Solve at least one leetcode problem per day",
  "date": "2025-07-17",
  "status": "Completed"
}
```

- **Success Response:**
  - 201 Created
  - Returns the created task

## Update a Task

- **Method:** PUT
- **URL:** `/tasks/:id`
- **Description:** Update a task.
- **Requires:** JWT token (admin only)
- **Headers:**
  `Authorization: Bearer <your_token>`
  `Content-Type: application/json`
- **Body:**

```json
{
  "title": "Learn MongoDB",
  "description": "Practice MongoDB",
  "date": "2025-07-20",
  "status": "In progress"
}
```

- **Success Response:**
  - 200 OK
  - Returns the updated task

## Delete a Task

- **Method:** DELETE
- **URL:** `/tasks/:id`
- **Description:** Delete a task.
- **Requires:** JWT token (admin only)
- **Headers:**
  `Authorization: Bearer <your_token>`
- **Success Response:**
  - 204 No Content

## Promote a User to Admin

- **Method:** PUT
- **URL:** `/promote/:id`
- **Description:** Promote a user to admin by their user ID.
- **Requires:** JWT token (admin only)
- **Headers:**
  `Authorization: Bearer <your_token>`
- **Success Response:**
  - 200 OK
  - `{ "message": "Promoted to admin" }`
- **Example Request:**
  `PUT http://localhost:8080/promote/664f1a2b3c4d5e6f7a8b9c0d`
- **Example Response:**

```json
{
  "message": "Promoted to admin"
}
```

## List All Users (Admin Only)

- **Method:** GET
- **URL:** `/users`
- **Description:** Get a list of all users. Only admins can use this endpoint.
- **Requires:** JWT token (admin only)
- **Headers:**
  `Authorization: Bearer <your_token>`
- **Success Response:**
  - 200 OK
  - Returns a list of users (id, username, role)
- **Example Request:**
  `GET http://localhost:8080/users`
- **Example Response:**

```json
[
  {
    "id": "664f1a2b3c4d5e6f7a8b9c0d",
    "username": "alice",
    "role": "user"
  },
  {
    "id": "664f1a2b3c4d5e6f7a8b9c0e",
    "username": "bob",
    "role": "admin"
  }
]
```

## Error Responses for Protected Endpoints

- **401 Unauthorized:** If you do not provide a valid token

```json
{
  "error": "Unauthorized"
}
```

- **403 Forbidden:** If you try to do admin actions as a regular user

```json
{
  "error": "Forbidden"
}
```
