# taskly-Go 🧠

Backend service for the [taskly](https://taskly.onrender.com/) platform — a task tracking and management tool, built with Golang and Fiber.

## 🔗 Live URLs

- **Frontend**: [taskly.onrender.com](https://taskly.onrender.com/)
- **Backend API**: [taskly-go.onrender.com](https://taskly-go.onrender.com/)  
- **Frontend Repo**: [github.com/deexithparand/taskly](https://github.com/deexithparand/taskly)

---

## 🚀 Getting Started

### 🔧 Prerequisites

- Go 1.20+
- [Air](https://github.com/cosmtrek/air) for live reloading (optional)
- Docker (optional)

### 🛠️ Build & Run Locally

#### 1. Clone the repo
```bash
git clone https://github.com/deexithparand/taskly-go.git
cd taskly-go
````

#### 2. Run using Air (for development)

```bash
air
```

> Make sure you have `air` installed. Install: `go install github.com/cosmtrek/air@latest`

#### 3. Or run with Docker

```bash
docker build -t taskly-go .
docker run -d -p 8000:8000 taskly-go
```

---

## 📦 API

All API endpoints are hosted at:
`https://taskly-go.onrender.com/`

## 🎥 Demo Video

Watch the demo video here: [Demo Video](https://www.loom.com/share/807b0d5ff7434e56957f5be5869bbc12?sid=a80245a8-e926-40e2-9a8f-01ff9027450d)
