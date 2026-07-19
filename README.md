# Gemini CLI Project Overview

This repository is for use in the Gemini CLI, helping me(and maybe others) get their work done. 

## Core Skills

We've got a handful of skills to lend a hand:

### Research Memory
This skill helps tuck away findings from your research. Think of it as keeping good notes, ensuring what you learn today can be found and used again tomorrow. It keeps things tidy in a structured markdown format in the `memories/` directory.

### Context Rehydration
When you need to recall past lessons, this skill helps bring those research memories back into play. It's about remembering what we've already figured out, so you don't have to start fresh every time. Finds these memories in the `memories/` directory.

### Session Summarizer
At the end of a long stretch of work, this skill helps sum things up. It lays out what got done, what tripped us up, and anything new that came along. Good for looking back and seeing where we stand. Summaries are kept in the `sessions/summaries/` directory.

## Workspaces

We've expanded into building other software in the `workspace/` directory. for example:

### Go
*   **Ticket Tracker** (`workspace/go/go-ticket-tracker`): A robust backend API for tracking tickets. It's built with Go, uses SQLite for storage, and provides an OpenAPI specification.

### Python
*   **Ticket Client** (`workspace/python/ticket-client`): A Python SDK designed to talk to the Ticket Tracker API. It handles the heavy lifting of communicating with the backend.

### Vue
*   **Ticket Viewer** (`workspace/vue/ticket-viewer`): A clean Vue.js frontend for viewing and managing tickets. (WIP)
*   **Geminon Spaceman** (`workspace/vue/geminon-spaceman`): A futuristic dashboard styled after a SPACE interface, demonstrating our UI capabilities.

## Extensions

Gemini-cli validated extensions, these are for making my life more fun, or yours:

*   **Shades of Pink** (`extensions/shades-of-pink-theme-extension`): A theme extension that brings a splash of color to the environment.

## Directories of Note

### `memories/`
This is where all the good research findings get stored. Each file is a structured note on something learned, ready to be pulled up when needed.

### `sessions/summaries/`
After a work session, a summary of what happened here gets filed away. It's a record of the day's progress, challenges, and what's next.

### `workspace/`
The home for our active coding projects. Whether it's Go, Python, or Vue, this is where the code lives.

### `extensions/`
A place for add-ons and themes that enhance the Gemini CLI experience.
