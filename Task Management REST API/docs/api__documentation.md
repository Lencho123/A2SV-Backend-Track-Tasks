# 📄 Task Management REST API – Documentation

This project is a simple, modular **Task Management REST API** built using the [Gin](https://github.com/gin-gonic/gin) web framework in Go. It allows you to perform basic **CRUD operations** (Create, Read, Update, Delete) on task objects stored in memory.

---

## 🗂️ Project Structure

```
task-management-rest-api/
├── main.go                 // Entry point of the application
├── controllers/            // Handles incoming HTTP requests
│   └── task_controller.go
├── models/                 // Contains the Task model definition
│   └── task.go
├── data/                   // In-memory data and business logic
│   └── task_service.go
├── router/                 // Route registration and server setup
│   └── router.go
├── docs/                   // Documentation files
│   └── documentation.md
└── go.mod                  // Go module file
```

---

## 🧠 System Overview

This API:
- Accepts HTTP requests to manage tasks.
- Operates on an in-memory slice of tasks (no database yet).
- Follows a layered architecture: **Model → Data → Controller → Router**.
- Can be easily extended to use a real database like **PostgreSQL** with **GORM**.

---

## 📌 Task Model (`models/task.go`)

The `Task` struct defines the structure of a task resource:

```go
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
```

---

## 🔀 API Routes (`router/router.go`)

The `ServerRoutes()` function initializes and binds all endpoints:

| Method | Endpoint        | Description               | Controller             |
|--------|------------------|---------------------------|------------------------|
| GET    | `/tasks`         | Fetch all tasks           | `GetTasks`             |
| GET    | `/tasks/:id`     | Fetch task by ID          | `GetTasksByID`         |
| POST   | `/tasks`         | Add a new task            | `AddNewTask`           |
| PUT    | `/tasks/:id`     | Update existing task      | `PutTask`              |
| DELETE | `/tasks/:id`     | Delete task by ID         | `DeleteTask`           |

---

## 🔧 Controller Logic (`controllers/task_controller.go`)

Each handler uses the `gin.Context` object to manage the HTTP lifecycle:

- **GetTasks**: Returns the full list of tasks.
- **GetTasksByID**: Looks up a task by ID.
- **AddNewTask**: Adds a task to the list.
- **PutTask**: Updates a task (if found).
- **DeleteTask**: Deletes a task by ID.

---

## 📦 In-Memory Data Service (`data/task_service.go`)

This acts like a temporary database using a slice of `Task`:

```go
var TaskDatas = []models.Task{
	{ID: 1, Title: "Read Clean Code", Description: "Go through Chapter 3: Functions"},
	...
}
```

### Provided Functions:

- `GetTaskWithID(id int) *Task`
- `PutTask(id int, updated Task) int`
- `DeleteTask(id int) int`

All CRUD operations modify `TaskDatas` directly.

---

## ▶️ Running the Project

### 🔧 Prerequisites
- Go installed (version 1.16+ recommended)
- `Gin` web framework (`go get -u github.com/gin-gonic/gin`)

### 🏁 Steps to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/task-management-rest-api.git
   cd task-management-rest-api
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. The API will be live at:
   ```
   http://localhost:8080
   ```

---

## 🧪 Example Requests

### Add New Task (POST `/tasks`)
```json
{
  "id": 6,
  "title": "Write API Docs",
  "description": "Summarize project and endpoints"
}
```

### Update Task (PUT `/tasks/3`)
```json
{
  "id": 3,
  "title": "Updated Title",
  "description": "Updated description"
}
```

---

## 📌 Notes
- Data is **not persisted** — all tasks reset when the server restarts.
- Add database integration in the future (e.g., GORM + PostgreSQL).
- The structure follows good separation of concerns, making future refactoring easy.