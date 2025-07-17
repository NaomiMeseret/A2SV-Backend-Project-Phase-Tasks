# Task Manger API Documentation

-> This document shows how to use the Task Manger API in clear , easy-to-follow steps
It describes:

1. Folder structure
2. How to configure MongoDB
3. Each "HTTP endpoint"
4. Which file handles the logic

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

  ``` mongodb://localhost:27017 ```

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

# API Endpoints

-> All URLS start with http://localhost:8080

1. List All Tasks

   - **Method**: GET
   - **URL**: `/tasks`
   - **Description**: Return an array of all tasks stored in MongoDB

   **Example Request**

   `GET http://localhost:8080/tasks`

   **Example Response**

   - 200 OK

```json
  [
    {
        "id": "68790b1719fc9d12d16e66d0",
        "title": "Learn MongoDB",
        "description": "Update Description",
        "date": "2025-07-20",
        "status": "In progress"
    },
    {
        "id": "68790b8319fc9d12d16e66d1",
        "title": "Watch a movie ",
        "description": "Watch We were liars series movie on Netflix",
        "date": "2025-07-17",
        "status": "In progress"
    },
    {
        "id": "68790e6bc905193546ab9122",
        "title": "Solve Leetcode Problem",
        "description": "Solve at least one leetcode probelm per day",
        "date": "2025-07-20",
        "status": "Completed"
    }
  ]
```


2. Get one task by ID

   - **Method**: GET
   - **URL**: `/tasks/:id` (Replace `:id` with the MongoDB generated ObjectID task's ID, e.g., `/tasks/68790b1719fc9d12d16e66d0`)
   - **Description**: Returns the task with the given ID

   **ExampleRequest**

   `GET http://localhost:8080/task/68790b1719fc9d12d16e66d0`
   **ExampleSuccessResponse**

   - 200 Ok

  ```json
  {
    "id": "68790b1719fc9d12d16e66d0",
    "title": "Learn MongoDB",
    "description": "Update Description",
    "date": "2025-07-20",
    "status": "In progress"
  }
  ```

   **ExampleErrorResponse**

   - 404 Not Found

   `{"message":"task not found"}`

3. Create a New Task

   - **Method**: POST
   - **URL**: `/tasks`
   - **Description**: Insert a new task in to the MongoDB
   - **Headers**:`Content-Type: application/json`
   - **Body(raw JSON)**:

  ```json
  {
      "title": "Solve Leetcode Problem",
      "description": "Solve at least one leetcode probelm per day",
      "date": "2025-07-17",
      "status": "Completed"
  }
  ```

   **Example Success Response**

   - 201 created

  ```json
  {
    "id": "68790e6bc905193546ab9122",
    "title": "Solve Leetcode Problem",
    "description": "Solve at least one leetcode probelm per day",
    "date": "2025-07-17",
    "status": "Completed"
  }
   ```

   **Example Error Response**

   • 400 Bad Request

   ` {"message":"invaild json body}`

4. Update an existing task

  - **Method**: PUT
  - **URL**: `/tasks/:id` (Replace `:id` with the ObjectID of the task you want to update)
  - **Description**: Updates the task's fields in MongoDB
  - **Headers**: `Content-Type: application/json`
  - **Body(raw json)**:

  ```json
  {
  "title": "Learn MongoDB",
  "description": "Practice MongoDB",
  "date": "2025-07-20",
  "status": "In progress"
  }
   ```

   **ExampleSucessResponse**

   - 200 Ok

  ```json
  {
  "title": "Learn MongoDB",
  "description": "Update Description",
  "date": "2025-07-20",
  "status": "In progress"
  }
   ```

   **Example ErrorResponse**

   - 404 Not Found

   `{"message":"task not found"}`

5. Delete a task

  - **Method** : DELETE
  - **URL** : `/tasks/:id` (Replace the ID to delete)
  - **Description** : Removes the task from the MongoDB

    **Example Success Response**

  - 204 No content
  - No JSON body

    **Example Error Response**

   ``` {"message": "task not found"}```
