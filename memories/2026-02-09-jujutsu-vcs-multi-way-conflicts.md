# Topic: Jujutsu (jj) Multi-way Conflict Handling and External Merge Tools

## Metadata
- **Date**: 2026-02-09
- **Tags**: vcs, jujutsu, conflicts, merge-tools, dev-tools
- **Status**: Completed
- **Related Memories**: 2026-02-09-jujutsu-vcs-conflict-experiment.md, 2026-02-09-jujutsu-vcs-trial.md

## Context/Problem
Investigating how Jujutsu (jj) handles conflicts involving more than two branches (multi-way) and how it integrates with external merge tools for resolution.

## Solution/Findings

### 1. Multi-way Conflicts
Jujutsu natively supports N-way (or N-sided) conflicts. A conflict is created whenever multiple divergent operations affect the same part of a file.
- **Algebraic Model**: Conflicts are represented as a set of "adds" and "removes" (e.g., `A + B + C - 2*Base`).
- **Conflict Markers**: For N-sided conflicts, `jj` renders markers that show the diff from the common base for each side.
  ```
  <<<<<<< conflict 1 of 1
  %%%%%%% diff from: <base_rev> "Base"
  \        to: <side_a_rev> "Side A"
  -Base Line
  +Side A Content
  %%%%%%% diff from: <base_rev> "Base"
  \        to: <side_b_rev> "Side B"
  -Base Line
  +Side B Content
  +++++++ <side_c_rev> "Side C"
  Side C Content
  >>>>>>> conflict 1 of 1 ends
  ```
- **Commitment**: Conflicted states are first-class citizens and can be committed, shared, and rebased just like any other commit.

### 2. External Merge Tool Integration
Jujutsu can delegate conflict resolution to external tools via the `jj resolve` command.
- **Configuration**: Merge tools are defined in the configuration (usually `~/.config/jj/config.toml`).
  ```toml
  [merge-tools.kdiff3]
  program = "kdiff3"
  merge-args = ["$base", "$left", "$right", "$output"]
  ```
- **Limitations**: Currently, `jj resolve` **only supports 3-way merges** (2-sided conflicts). If a conflict has more than 2 sides, `jj resolve` will fail with an error: `"The conflict at 'path' has N sides. At most 2 sides are supported."`
- **Manual Resolution**: For N-sided conflicts (N > 2), developers must currently resolve them manually by editing the conflict markers in their editor.

## Actionable Insights/Next Steps
- When dealing with complex merges of 3+ branches, be prepared for manual conflict resolution as automated tools might only handle pairs.
- Leverage `jj resolve --tool <name>` for standard 3-way merges.
- Use `jj config set --user merge-tools.<name>.<key> <value>` to quickly configure new tools.

## References
- Jujutsu Documentation: https://jj-vcs.dev/docs/conflicts/
- Conflict Resolution Command: `jj help resolve`
