---
name: context-rehydration
description: Searches previously saved research memories in the 'memories' directory by keyword (via rehydrate.sh) and loads them into the current context. Use when the user says "recall", "remember", "do you remember", "have we done/researched this before", or refers to a previous session or past research.
---

# Context Rehydration Skill

This skill allows you to "remember" past learnings by loading structured memory files into the current session.

## Workflow

1. **Search and Load**: Use the `rehydrate.sh` helper script (at the repo root) to find and display memories by keyword. It matches against filenames and file contents, including the YAML frontmatter tags.
   - Example: `./rehydrate.sh jujutsu`
   - On multiple matches it lists the candidate files with a matching-line preview; re-run with a narrower keyword or `cat` the specific file.
2. **Synthesize**: Read the frontmatter metadata and the "Actionable Insights/Next Steps" section from the output to inform the current task.
3. **Apply Learning**: Use the rehydrated context to proceed with the user's request.

## Automation & Validation

- **Helper Script**: `./rehydrate.sh` provides a fast-track for searching and reading memories.

## Search Strategy

- Prefer `rehydrate.sh`; if it is unavailable, fall back to `grep -ril <keyword> memories/` — content search, not just filenames, since tags live in frontmatter.
- If the `memories/` directory doesn't exist, inform the user that no memories have been recorded yet.