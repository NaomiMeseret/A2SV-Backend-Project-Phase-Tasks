# Task Manger API Documentation

-> This document shows how to use the Task Manger API in clear , easy-to-follow steps
It describes:

1. Folder structure
2. Each "HTTP endpoint"
3. Which file handles the logic

## Folder Structure

```plaintext
task_manager/
├── main.go                  => Starts the app by calling the router setup
├── controllers/
│   └── task_controller.go   => Handles HTTP requests and calls the service function
├── models/
│   └── task.go              => Defines the Task structure (fields and JSON tags)
├── data/
│   └── task_service.go      => Contains logic for storing, creating, updating, and deleting tasks
├── router/
│   └── router.go            => Connects each HTTP path and method to the correct controller
├── docs/
│   └── api_documentation.md => You are here
└── go.mod
```

# Running the server

1. Open a terminal in a project folder
2. Run:

```bash
go run main.go
```

3. Base URL
   ```http://localhost:8080```

# API Endpoints

-> All URLS start with http://localhost:8080

1. List All Tasks

   - **Method**: GET  
   - **URL**: `/tasks`  
   - **Description**: This endpoint shows you every task  

   **Example Request**

    `GET http://localhost:8080/tasks`

    **Example Response**

    - 200 OK 
    
    ```json
        [
        {
            "id": "1",
            "title": "Note Taking",
            "description": "Take a note for a daily standup meeting",
            "date": "2025-07-16",
            "status": "Completed"
        },
        {
            "id": "2",
            "title": "Do Task 3",
            "description": "Read documenation and complete the task",
            "date": "2025-07-16",
            "status": "in progress"
        },
        {
            "id": "3",
            "title": "Solve leetcode problem",
            "description": "Solve at least 2 problem using Go",
            "date": "2025-07-16",
            "status": "In progress"
        }
    ]``` 
2. Get one task by ID

    - **Method**: GET  
    - **URL**: `/tasks/:id` (Replace `:id` with the task's ID, e.g., `/tasks/2`)  
    - **Description**: Returns the task with the given ID  

    **ExampleRequest**

    ```GET http://localhost:8080/task/2```
    **ExampleSuccessResponse**

    - 200 Ok

    ```json 
        {
        "id": "2",
        "title": "Do Task 3",
        "description": "Read documenation and complete the task",
        "date": "2025-07-16",
        "status": "in progress" }
    ```
    **ExampleErrorResponse**

    - 404 Not Found

    ```{"message":"task not found"}```

3. Create a New Task

    - **Method**: POST  
    - **URL**: `/tasks`  
    - **Description**: Adds a new task to the list  
    - **Headers**:`Content-Type: application/json`
    - **Body(raw JSON)**:
    ```json
        {
            "id": "3",
            "title": "Solve leetcode problem",
            "description": "Solve at least 2 problem using Go",
            "date": "2025-07-16",
            "status": "In progress"
        }
    ```
    **Example Success Response**

    - 201 created

    ```json
        {
            "id": "3",
            "title": "Solve leetcode problem",
            "description": "Solve at least 2 problem using Go",
            "date": "2025-07-16",
            "status": "In progress"
        }
    ```
    **Example Error Response**

    • 400 Bad Request

    ``` {"message":"invaild json body}```

4. Update an existing task
    - **Method**: PUT  
    - **URL**: `/tasks/:id` (Replace `:id` with the ID you want to update)  
    - **Description**: This endpoint replaces the taks's data with what you send
    - **Headers**: `Content-Type: application/json`  
    - **Body(raw json)**:

    ```json
        {
            "id": "3",
            "title": "Solve Two leetcode problem",
            "description": "Aim for 2 problems per day using Go",
            "date": "2025-07-16",
            "status": "In progress"
        }
    ```
    **ExampleSucessResponse**

    - 200 Ok

    ```json
        {
            "id": "3",
            "title": "Solve Two leetcode problem",
            "description": "Aim for 2 problems per day using Go",
            "date": "2025-07-16",
            "status": "In progress"
        }
    ```

    **Example ErrorResponse**

    - 404 Not Found

    `{"message":"task not found"}` 

5. Delete a task

    - **Method** : DELETE
    - **URL** : `/tasks/:id` (Replace the ID to delet)
    • **Description** :This endpoint remove the task from the list.
    **Example Success Response**

    - 204 No content
    -  No JSON body

    **Example Error Response**

    ``` {"message": "task not found"}```

