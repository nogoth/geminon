# Ticket Viewer (Vue)

A responsive 3-column grid viewer for the Ticket Tracker system.

## Features
- 3xN Grid Layout (CSS Grid)
- Floating Action Button (+) for ticket creation
- Animated creation overlay with background dimming
- Dark/Light mode support

## Usage

### 1. Start the Backend
```bash
cd ../../go/go-ticket-tracker
go run cmd/server/main.go
```

### 2. Start the Frontend
```bash
cd workspace/vue/ticket-viewer
npm install
npm run dev
```

The application will typically be available at `http://localhost:5173`.
