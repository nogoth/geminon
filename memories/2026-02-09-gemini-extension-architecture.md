# Topic: Gemini CLI Extension & Agent Skill Architecture

## Metadata
- **Date**: 2026-02-09
- **Tags**: gemini-cli, extension-development, agent-skills, mcp-protocol, architecture
- **Status**: Completed
- **Related Memories**: 2026-02-09-mcporter-research.md

## Context/Problem
As the Gemini CLI evolves, there is a need for a clear understanding of its extension mechanisms to build portable, modular, and specialized tools. This research explores the relationship between extensions, agent skills, and the Model Context Protocol (MCP).

## Solution/Findings
The Gemini CLI uses a tiered architecture to manage AI capabilities, separating physical tool implementation from procedural expertise.

### 1. The Mental Model
The system is built on "tiered modularity":
- **Extensions** are the "Package": They contain the manifest (`gemini-extension.json`) and the "Body" (MCP servers and tools).
- **Agent Skills** are the "Expertise": They provide on-demand, step-by-step procedural logic (`SKILL.md`) that is only activated when relevant.
- **Persistent Context** (`GEMINI.md`) is the "Persona": It sets the baseline behavior and goals for all sessions where the extension is active.

This ensures the agent stays focused, only loading complex logic when a task-based trigger is met.

### 2. Core Reference & Cheat Sheet
| Component | Key File | Key Property | Purpose |
| :--- | :--- | :--- | :--- |
| **Extension Manifest** | `gemini-extension.json` | `mcpServers` | Defines name, version, and tool execution logic. |
| **Persistent Context** | `GEMINI.md` | Markdown text | Baseline instructions loaded into every session. |
| **Agent Skill** | `SKILL.md` | Frontmatter `description` | On-demand procedural expertise triggered by tasks. |
| **Custom Commands** | `commands/*.toml` | `prompt` | Shortcuts for complex prompt templates. |

### 3. Agent-Specific Guidance
- **Activation Precision**: The `description` field in a skill's YAML frontmatter is the most critical field for AI agents. It should clearly define the task types that trigger the skill.
- **Variable Substitution**: Use `${extensionPath}` in the manifest to ensure tools can find their internal scripts regardless of where the extension is installed.
- **Resource Declaration**: In `SKILL.md`, use the `available_resources` section to list the MCP tools the skill relies on, aiding the agent's planning phase.

## Actionable Insights/Next Steps
- **Modularize existing tools**: Convert loose scripts in the root into a structured extension.
- **Skill-first Documentation**: When writing internal docs for new features, use the `SKILL.md` format so the agent can immediately "learn" the new procedure.
- **Standardize Manifests**: Ensure all local extensions use the `gemini-extension.json` standard for consistent loading.

## References
- Gemini CLI Help (`cli_help` tool)
- Official Extension Documentation (geminicli.com)
- Community Examples (philschmid/gemini-cli-extension)
