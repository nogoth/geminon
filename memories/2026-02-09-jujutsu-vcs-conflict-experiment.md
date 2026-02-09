# Topic: Jujutsu (jj) VCS Conflict Handling

## Metadata
- **Date**: 2026-02-09
- **Tags**: vcs, jujutsu, git, conflicts, developer-tools
- **Status**: Completed
- **Related Memories**: 2026-02-09-jujutsu-vcs-trial.md

## Context/Problem
In traditional VCS like Git, conflicts are a "stop-the-world" event. You cannot commit while in a conflicted state, and the repository is locked into a merge/rebase mode. This disrupts the developer's flow and makes it difficult to share or postpone conflict resolution.

## Solution/Findings
Jujutsu (jj) treats conflicts as first-class citizens. A commit (change) can exist in a conflicted state indefinitely. This allows developers to continue working, share conflicted branches, and resolve them at their own pace.

### 1. The Mental Model
Conflicts in Jujutsu are not markers in a file that break the build; they are data structures within the commit. The working copy automatically renders these data structures as conflict markers. Resolving a conflict is simply a matter of updating the content of the commit (or a subsequent one) to no longer contain the conflict data.

### 2. Core Reference & Cheat Sheet
| Command | Description |
|---------|-------------|
| `jj rebase -s <rev> -d <dest>` | Rebases a revision and its descendants. Can create conflicts without stopping. |
| `jj log` | Shows conflicted commits with a `(conflict)` label. |
| `jj resolve` | Opens a tool to resolve conflicts (if configured). |
| `jj squash` | Can be used to merge a resolution commit into the conflicted parent. |
| `jj evolog` | Shows the evolution of a change, including how conflicts were introduced or resolved. |

### 3. Agent-Specific Guidance
- **Shell Quoting**: When using functional selectors like `description("...")`, always use single quotes around the entire argument to prevent shell expansion or syntax errors.
- **Working Copy**: Remember that the working copy is always a commit (`@`). Editing files immediately updates this commit.
- **Conflict Markers**: Jujutsu uses a descriptive "Snapshot" style by default, showing the diff from the common base to both sides.

## Actionable Insights/Next Steps
- Use `jj` for complex refactors where multiple divergent branches might need to be merged incrementally.
- Leverage the ability to "park" a conflict and switch tasks without losing state.
- Explore `jj git push` to see how conflicted commits are handled when pushed to a Git remote (they are usually blocked unless resolved).

## References
- Jujutsu Documentation: https://github.com/martinvonz/jj
- Conflict Resolution Guide: https://martinvonz.github.io/jj/latest/conflicts/
