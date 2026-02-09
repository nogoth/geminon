---
name: context-rehydration
description: Searches and loads previously saved research memories from the 'memories' directory into the current context. Use this when the user asks about past research or when a relevant memory might help solve the current task.
---

# Context Rehydration Skill

This skill allows you to "remember" past learnings by loading structured memory files into the current session.

## Workflow

1. **Search and Load**: Use the `rehydrate.sh` helper script to find and display memories by keyword.
   - Example: `./rehydrate.sh jujutsu`
2. **Synthesize**: Read the "Metadata" and "Actionable Insights" from the output to inform the current task.
3. **Apply Learning**: Use the rehydrated context to proceed with the user's request.

## Automation & Validation

- **Helper Script**: `/opt/rehydrate.sh` provides a fast-track for searching and reading memories.
- **Verification**: Use `/opt/tests/test_context_rehydration.sh` to ensure the memory library remains healthy and follows the expected template.

## Search Strategy

- Use `list_directory` on the `memories/` folder.
- If the directory doesn't exist, inform the user that no memories have been recorded yet.
- Match filenames against keywords from the current task.