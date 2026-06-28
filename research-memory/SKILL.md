---
name: research-memory
description: Saves research findings into a structured markdown format in the 'memories' directory to preserve learning for future sessions. Use this after completing research or solving a complex problem to ensure the knowledge is reusable.
---

# Research Memory Skill

This skill helps you preserve knowledge by saving research findings into a structured markdown file.

## Workflow

1. **Synthesize Findings**: Summarize the topic, findings, and actionable insights. Aim for an **"Agent-Ready"** format that allows a future agent to pick up the tool/concept instantly.
2. **Determine Metadata**: Identify relevant tags and any related memories.
3. **Format Memory**: Use the structure provided in `assets/memory-template.md`.
   - **Mental Model**: Explain *how* to think about the tool, not just *what* it is.
   - **Reference Table**: Provide a quick command/API lookup (especially translations from common tools like Git).
   - **Agent Tips**: Highlight specific technical gotchas for LLM/CLI interaction (e.g., escaping, non-interactive flags).
4. **Name and Save**: 
   - Directory: `memories/` (create if it doesn't exist).
   - Filename: `YYYY-MM-DD-kebab-case-topic.md` (e.g., `2026-02-09-docker-compose-react.md`).
5. **Write File**: Save the memory using the Write tool.

## Metadata Guidelines

- **Topic**: Clear, concise name of the subject.
- **Date**: YYYY-MM-DD.
- **Tags**: Comma-separated list of keywords.
- **Actions**: Concrete steps that can be taken based on this research.