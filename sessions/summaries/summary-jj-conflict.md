# Session Summary: Jujutsu Conflict Experiment

## Technical Challenge
Demonstrating Jujutsu's (jj) unique approach to version control conflicts compared to traditional systems like Git.

## What Worked
- Successfully initialized a co-located Git/Jujutsu repository.
- Created an intentional two-sided conflict by rebasing divergent branches.
- Observed that Jujutsu allows the repository to stay in a "normal" state even with active conflicts.
- Verified the descriptive conflict marker format.
- Resolved the conflict using the `jj squash` workflow.

## Pitfalls
- Initial syntax errors in shell commands due to improper quoting of Jujutsu revision selectors.
- Misunderstanding of `jj git init --colocate` vs `--colocated`.

## Lessons Learned
- Jujutsu's "conflicts as first-class citizens" model significantly reduces the friction of rebasing and merging.
- Proper shell quoting is critical when using `jj`'s powerful functional query language.

## Future Work
- Investigate how `jj` handles multi-way conflicts and its integration with external merge tools.
