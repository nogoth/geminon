# Topic: LinkedIn Voice Documentation and Usage Examples

## Metadata
- **Date**: 2026-02-09
- **Tags**: linkedin-writer, documentation, technical-writing, personas
- **Status**: Completed
- **Related Memories**: 2026-02-09-linkedin-voice-library-expansion.md, 2026-02-09-linkedin-voice-styles-integration.md

## Context/Problem
While the `linkedin-writer` skill had been expanded with new technical personas, there was a lack of clear documentation on how to actually invoke these voices via the CLI, specifically using the `--voice` parameter.

## Solution/Findings
Updated `linkedin-writer/SKILL.md` to include a "Usage Examples" section. This section provides concrete CLI commands for three distinct personas (Hacker, Minimalist, and Mentor) to serve as a template for users.

Key updates:
- Added `## Usage Examples` section.
- Provided example commands:
  - `linkedin-writer sessions/summaries/summary-5.md --voice Hacker`
  - `linkedin-writer sessions/summaries/summary-5.md --voice Minimalist`
  - `linkedin-writer sessions/summaries/summary-5.md --voice Mentor`

## Actionable Insights/Next Steps
- Users can now quickly reference `SKILL.md` to understand the syntax for persona-driven post generation.
- Future voice additions should follow this pattern of including usage examples in the skill documentation to ensure discoverability.

## References
- `linkedin-writer/SKILL.md`
- `linkedin-writer/voice_styles.md`
