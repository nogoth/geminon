---
name: research-memory
description: Saves research findings into a structured markdown format in the 'memories' directory to preserve learning for future sessions. Use when the user says "remember this", "save this for later", "write this down", or "take a note of this", and after completing research or solving a complex problem worth reusing.
---

# Research Memory Skill

This skill helps you preserve knowledge by saving research findings into a structured markdown file.

## Workflow

1. **Synthesize Findings**: Summarize the topic, findings, and actionable insights. Aim for an **"Agent-Ready"** format that allows a future agent to pick up the tool/concept instantly.
2. **Determine Metadata**: Identify relevant tags and any related memories. Metadata lives in YAML frontmatter at the top of the file so retrieval tooling (`rehydrate.sh`) can search it.
3. **Format Memory**: Use the structure provided in `assets/memory-template.md`.
   - **Mental Model**: Explain *how* to think about the tool, not just *what* it is.
   - **Reference Table**: Provide a quick command/API lookup (especially translations from common tools like Git).
   - **Agent Tips**: Highlight specific technical gotchas for LLM/CLI interaction (e.g., escaping, non-interactive flags).
4. **Name and Save**: 
   - Directory: `memories/` (create if it doesn't exist).
   - Filename: `YYYY-MM-DD-kebab-case-topic.md` (e.g., `2026-02-09-docker-compose-react.md`).
5. **Write File**: Save the memory using the Write tool.

## Metadata Guidelines

All metadata goes in the YAML frontmatter block:

- **topic**: Clear, concise name of the subject.
- **date**: YYYY-MM-DD.
- **tags**: List of lowercase keywords, e.g. `[vcs, jujutsu, conflicts]`.
- **status**: `Completed`, `In-Progress`, or `Stale`.
- **related**: List of related memory filenames that exist in `memories/`; `[]` if none.