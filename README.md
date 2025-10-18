# PCST-AI

### Process Control System Technicians' AI Platform

A modern web application featuring AI-powered troubleshooting for industrial process control systems, built with Go backend and Vue frontend.

## PREREQUISITES
- Go 1.21 or higher
- Node.js 18+ and npm
- Google Cloud account with Gemini API access

## SETUP STEPS

### 1. Environment Configuration
Copy `env.example` to `.env` and add your Gemini API key:
```
GEMINI_API_KEY=your_actual_api_key_here
```

### 2. Backend Setup
```bash
cd backend
go mod tidy
go run main.go handlers.go
```
Backend will run on http://localhost:8080

### 3. Frontend Setup
```bash
cd frontend
npm install
npm run dev
```
Frontend will run on http://localhost:5173

-----

#### TODO
- Containerize
- Set a database up
- Build an ILD to use for references