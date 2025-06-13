# Intern Pocket - To Do App (Gin + MongoDB + Cron Jobs) 🚀📝

A professional, scalable, and extensible To Do application built with [Gin](https://gin-gonic.com/) (Go web framework) and [MongoDB](https://www.mongodb.com/), featuring scheduled background tasks using [robfig/cron](https://github.com/robfig/cron).  
Made with ❤️ for productive developers!

---

## Features ✨

- **RESTful API** for managing tasks (CRUD operations) 🔄
- **MongoDB** for persistent, flexible data storage 🍃
- **Scheduled Jobs** using cron expressions (e.g., send reminders, clean up tasks) ⏰
- **Modular Codebase** for easy maintenance and extension 🧩
- **Clear Documentation** and code comments 📚

---

## Project Structure 🗂️

<!-- Add your project structure here if needed -->

---

## Getting Started 🚦

### Prerequisites

- Go 1.18+ 🦫
- MongoDB (local or cloud) 🍃
- (Optional) Docker 🐳

### Installation

1. **Clone the repository**
    ```sh
    git clone https://github.com/yourusername/intern-pocket.git
    cd intern-pocket/app/gin-mongodb-api
    ```

2. **Install dependencies**
    ```sh
    go mod tidy
    ```

3. **Configure environment variables**
    - Set your MongoDB URI and other configs as needed.

4. **Run the application**
    ```sh
    go run main.go
    ```

---

## API Endpoints 📡

| Method | Endpoint         | Description           |
|--------|------------------|----------------------|
| GET    | `/tasks`         | List all tasks       |
| POST   | `/tasks`         | Create a new task    |
| GET    | `/tasks/:id`     | Get a task by ID     |
| PUT    | `/tasks/:id`     | Update a task        |
| DELETE | `/tasks/:id`     | Delete a task        |

---

## Cron Jobs ⏳

- **Scheduling:** Uses [robfig/cron](https://github.com/robfig/cron) for background jobs.
- **Examples:**
    - `@every 1s` — Run every second ⚡
    - `0 0 * * 0` — Run every Sunday at midnight 🌅
    - `0 12 * * *` — Run every day at noon 🕛

**How to add a new job:**
```go
c.AddFunc("CRON_EXPRESSION", YourTaskFunction)
```

---

Happy coding! 🎉👩‍💻👨‍💻