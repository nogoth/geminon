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

## Guidelines

*   **Be Objective**: Report both successes and failures accurately.
*   **Be Specific**: Instead of "fixed bugs," say "resolved the race condition in the authentication middleware."
*   **Contextualize Pitfalls**: Explain *why* something was a pitfall (e.g., "The library documentation was outdated regarding version 3.0").
*   **Identify Future Work**: Think about what you would do next if the session continued.