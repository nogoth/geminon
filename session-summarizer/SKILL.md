---
name: session-summarizer
description: Summarize the current session's activities, successes, failures, pitfalls, and environmental changes. Use when the user wants a post-mortem, session summary, or wrap-up of their work.
---

# Session Summarizer

## Overview

The `session-summarizer` skill provides a structured way to reflect on the current work session. It helps identify what was achieved, what hurdles were encountered, and what environmental changes occurred.

## Workflow

To generate a comprehensive session summary, follow these steps:

1.  **Analyze Session History**: Review the conversation history to identify:
    *   Successful implementations or fixes.
    *   Errors, failed attempts, and blockers.
    *   Specific challenges and how they were navigated (pitfalls).
2.  **Inspect Environment Changes**: Check for any newly installed software, libraries, or global configuration changes.
    *   Look for commands like `npm install`, `pip install`, `apt install`, etc., in the conversation or shell history if available.
3.  **Synthesize Future Ideas**: Based on the session's progress and remaining issues, suggest logical next steps or improvements.
4.  **Format and Save the Summary**: 
    *   Use the template provided in [references/summary-template.md](references/summary-template.md) to structure the output.
    *   **Save to File**: Create a directory `sessions/summaries/` if it doesn't exist. Save the summary to `sessions/summaries/summary-<NUMBER>.md`, where `<NUMBER>` is the next available integer (starting at 1).
    *   **Provide Output**: Display the full summary in the current session as well.

## Automatic Summaries via SessionEnd Hook

Claude Code can trigger this skill automatically when a session ends, using a `SessionEnd` hook that runs [scripts/session-end-summary.sh](scripts/session-end-summary.sh). The script reads the hook's stdin JSON (`transcript_path`, `cwd`, `reason`), pipes the session transcript to a headless `claude -p` call with the summarizer instructions, and writes the result to `sessions/summaries/summary-<N>.md` itself — the headless call needs no write permissions.

To enable it, add this to `.claude/settings.json` (team-shared) or `.claude/settings.local.json` (personal):

```json
{
  "hooks": {
    "SessionEnd": [
      {
        "hooks": [
          {
            "type": "command",
            "command": "session-summarizer/scripts/session-end-summary.sh"
          }
        ]
      }
    ]
  }
}
```

Caveats:

*   Each session end costs one headless model call.
*   The transcript file is written asynchronously and may slightly lag the live session.
*   `SessionEnd` hooks cannot block and their output is never shown — check `sessions/summaries/` for the result. The script exits silently on any failure.
*   The hook does not fire for `claude -p` (non-interactive) runs, so the nested headless call cannot recurse.
*   Sessions ending with reason `resume` are skipped (the session isn't over).

## Guidelines

*   **Be Objective**: Report both successes and failures accurately.
*   **Be Specific**: Instead of "fixed bugs," say "resolved the race condition in the authentication middleware."
*   **Contextualize Pitfalls**: Explain *why* something was a pitfall (e.g., "The library documentation was outdated regarding version 3.0").
*   **Identify Future Work**: Think about what you would do next if the session continued.