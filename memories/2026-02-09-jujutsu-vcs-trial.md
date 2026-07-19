---
topic: Jujutsu (jj) Version Control
date: 2026-02-09
tags: [vcs, jujutsu, git, workflow, dev-tools]
status: Completed
related: []
---

# Topic: Jujutsu (jj) Version Control

## Context/Problem
Exploring modern alternatives to Git that offer a simplified mental model, better history management, and first-class conflict handling while maintaining compatibility with existing Git ecosystems.

## Solution/Findings
Jujutsu (jj) successfully replaces the staging area with a commit-based working copy. It treats the working copy as a "commit in progress" that is automatically snapshotted.

### 1. The Jujutsu Mental Model
*   **No Staging Area**: Forget `git add`. Changes in your workspace are automatically tracked.
*   **The Working Copy Commit (`@`)**: Your current workspace is always a commit. When you finish a task, you use `jj new` to move to a new "child" commit.
*   **Anonymous Commits**: Commits don't need names or branches. They are identified by a short change ID (stable) and a commit ID (changes when rebased/amended).
*   **Immutable History is a Myth**: `jj` makes it trivial to edit any commit in the graph at any time.

### 2. Core Commands Reference
| Command | Git Equivalent | Description |
| :--- | :--- | :--- |
| `jj log` | `git log --graph` | View the revision graph. |
| `jj st` | `git status` | Show changes in `@`. |
| `jj describe` | `git commit --amend -m` | Update the description of a commit. |
| `jj new` | `git checkout -b` / `git commit` | Create a new commit on top of the current one. |
| `jj squash` | `git commit --amend` | Move changes from a commit into its parent. |
| `jj diff` | `git diff` | View changes in the working copy. |
| `jj rebase` | `git rebase` | Move a commit and its descendants to a new parent. |
| `jj undo` | `git reset --hard HEAD@{1}` | Undo the last `jj` operation (stored in op log). |
| `jj git fetch` | `git fetch` | Update from Git remotes. |
| `jj git push` | `git push` | Push branches to Git remotes. |

### 3. Revset Cheat Sheet (Selection Syntax)
Revsets are used to identify commits. **Always wrap in single quotes** to avoid shell issues.
*   `@`: The working copy commit.
*   `@-`: The parent of the working copy.
*   `description("feat:")`: Commits with "feat:" in the message.
*   `mine()`: Commits authored by you.
*   `all()`: All commits (including hidden/archived ones).
*   `x::y`: Commits reachable from `x` up to `y` (descendants of `x` that are ancestors of `y`).
*   `~::@`: All ancestors of the working copy.

### 4. Conflict Management
*   **Conflicts are committed**: Unlike Git, `jj` allows you to commit files with conflict markers. This means you can "save" a conflicted state and come back to it later or rebase it.
*   **Resolution**: Simply edit the file to resolve the markers, then `jj squash` or just move to a new commit. The conflict is resolved when the file no longer contains markers.

### 5. Git Interoperability
*   **Colocated Repos**: Created via `jj git init --colocate`. Both `.git` and `.jj` exist. `jj` tracks the Git branches as "bookmarks".
*   **Bookmarks**: Equivalent to Git branches. Use `jj bookmark create <name>` and `jj bookmark set <name> -r <rev>`.

## Actionable Insights/Next Steps
- **Workflow Strategy**: Use `jj new` frequently to partition work. Use `jj squash -i` (interactive) to fine-tune what goes into each commit.
- **Undo is a Superpower**: If a rebase goes wrong, `jj op log` followed by `jj undo` will restore the *entire repository state* to before the command.
- **Configuration**: Set `ui.default-command = "log"` to see the graph by default.
- **Safety**: `jj` is extremely safe because every operation is recorded in the `op log`.

## References
- Official Docs: https://jj-vcs.dev/
- Git Comparison: https://jj-vcs.dev/docs/git-comparison/
- Revset Documentation: https://jj-vcs.dev/docs/revsets/