---
name: linkedin-writer
description: Converts raw technical logs and summaries into professional LinkedIn updates using customizable authorial voices.
---

# LinkedIn Writer

## Overview
The `linkedin-writer` skill converts raw technical logs and summaries into professional updates. It supports multiple authorial "voices" to match the user's authentic style while maintaining a professional, emoji-free standard.

## Workflow
1. **Source Extraction**: Read the target `sessions/summaries/summary-<NUMBER>.md` file.
2. **Voice & Style Selection**:
   - Check for a `--voice <NAME>` parameter or a voice notation in the input.
   - If a voice is specified, look for its definition in `linkedin-writer/voice_styles.md`.
   - If no voice is specified or found, fallback to the **Default Professional** voice.
3. **Thematic Analysis**: Identify the central technical challenge or achievement to form the "Lead."
4. **Drafting the Post**:
   - Apply the selected Voice Guidelines (Tone, Constraints, Goal).
   - **The Lead**: Start with a strong opening adapted to the voice.
   - **The Narrative**: Translate technical milestones through the lens of the voice.
   - **The Insight/Educator**: Frame takeaways as strategic insights or educational points.
5. **Formatting**:
   - Use double line breaks for readability.
   - Use dashes (-) or asterisks (*) for lists.
   - **Strictly No Emojis**.
6. **Persistence**: Save the final text to `sessions/summaries/linkedin/post-<NUMBER>.txt`.

## Usage Examples
The `--voice` parameter can be used to override the default style. Here are a few examples:

- **Hacker Voice**: Focused on raw technical details.
  `linkedin-writer sessions/summaries/summary-5.md --voice Hacker`
- **Minimalist Voice**: For high-impact, concise updates.
  `linkedin-writer sessions/summaries/summary-5.md --voice Minimalist`
- **Mentor Voice**: To share lessons learned with the community.
  `linkedin-writer sessions/summaries/summary-5.md --voice Mentor`

## Style Guidelines (Default Voice: Professional)
- **Tone**: Professional, analytical, and authoritative.
- **Vocabulary**: Use industry-standard terms (e.g., "bottleneck," "dependency management").
- **Visuals**: Use capitalization and whitespace to create hierarchy instead of icons.

## Custom Voice Configuration
Users can define custom voices in a `voice_styles.md` file using the following Markdown notation:

```markdown
## [Voice Name]
- **Tone**: [Description of the tone, e.g., Minimalist, Storyteller]
- **Goal**: [Primary objective of the post]
- **Constraints**: [Specific formatting or narrative rules]
```

## Voice Management
The `linkedin-writer` skill includes a Voice Manager tool to register and manage custom voices in `voice_styles.md`.

### Registration Workflow
1. **Guided Process**: The agent will ask for:
   - **Voice Name**: A unique identifier.
   - **Tone**: The desired personality (e.g., "Hacker-style," "Executive," "Mentor").
   - **Goal**: What the voice aims to achieve (e.g., "Recruitment," "Thought Leadership").
   - **Constraints**: Specific rules (e.g., "Short sentences," "Focus on code snippets").
2. **Quick Prompt**: Users can provide a single-sentence description of the voice they want, and the agent will derive the attributes.

### Storage
All voices are stored in `linkedin-writer/voice_styles.md` using the standard Markdown notation.
