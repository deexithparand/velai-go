# Velai-Go 🧠

Backend service for the [Velai](https://velai.onrender.com/) platform — a job tracking and career management tool.

## 🔗 Live URLs

- **Frontend**: [velai.onrender.com](https://velai.onrender.com/)
- **Backend API**: [velai-go.onrender.com](https://velai-go.onrender.com/)  
- **Frontend Repo**: [github.com/deexithparand/velai](https://github.com/deexithparand/velai)

---

## 🚀 Getting Started

### 🔧 Prerequisites

- Go 1.20+
- [Air](https://github.com/cosmtrek/air) for live reloading (optional)
- Docker (optional)

### 🛠️ Build & Run Locally

#### 1. Clone the repo
```bash
git clone https://github.com/deexithparand/velai-go.git
cd velai-go
````

#### 2. Run using Air (for development)

```bash
air
```

> Make sure you have `air` installed. Install: `go install github.com/cosmtrek/air@latest`

#### 3. Or run with Docker

```bash
docker build -t velai-go .
docker run -d -p 8000:8000 velai-go
```

---

## 📦 API

All API endpoints are hosted at:
`https://velai-go.onrender.com/`

