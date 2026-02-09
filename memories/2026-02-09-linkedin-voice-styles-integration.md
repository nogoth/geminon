# Topic: LinkedIn Voice Styles Integration

## Metadata
- **Date**: 2026-02-09
- **Tags**: linkedin, skill-development, automation, writing-styles
- **Status**: Completed
- **Related Memories**: 2026-02-09-jujutsu-vcs-conflict-experiment.md

## Context/Problem
The `linkedin-writer` skill needed a more integrated way to manage diverse authorial voices. Previously, `voice_styles.md` was a separate file in the root, making the skill less portable.

## Solution/Findings
- Integrated `voice_styles.md` into the `linkedin-writer` skill directory.
- Updated `linkedin-writer/SKILL.md` to reference the local `voice_styles.md`.
- Expanded the available voices to include:
  - **Professional**: Analytical and authoritative.
  - **Hacker**: Raw, technical, and fast-paced.
  - **Minimalist**: Extremely concise and impactful.
  - **Mentor**: Encouraging and educational.
  - **Visionary**: High-level and future-oriented.

## Actionable Insights/Next Steps
- Use the `--voice` parameter to toggle between different writing styles.
- When generating posts for technical experiments, the "Hacker" voice provides the best signal-to-noise ratio for developer audiences.
- The "Visionary" voice is best for high-level project milestones.

## References
- `linkedin-writer/SKILL.md`
- `linkedin-writer/voice_styles.md`
