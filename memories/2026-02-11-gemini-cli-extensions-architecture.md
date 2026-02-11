---
topic: Gemini CLI Extensions Architecture and Best Practices
date: 2026-02-11
tags: gemini-cli, extensions, architecture, skills, custom-commands, hooks, mcp-servers, GEMINI.md, gemini-extension.json, SKILL.md
---

# Gemini CLI Extensions: Architecture and Best Practices

This memory synthesizes findings from analyzing example Gemini CLI extensions to understand their structure, purpose, and best practices for design and implementation.

## Core Components and Structure

A Gemini CLI extension is a modular package that extends the CLI's capabilities. Key components include:

-   **`gemini-extension.json`**:
    *   **Purpose**: The central manifest file for an extension, providing essential metadata (e.g., `name`, `version`, `description`).
    *   **Declarative Features**: Can declaratively define core functionalities such as:
        *   `excludeTools`: An array of strings specifying tools or specific tool calls to be blocked (e.g., `"run_shell_command(rm -rf)"`).
        *   `mcpServers`: Definitions for Managed Code Protocol (MCP) servers that expose new programmatic tools to the agent.
        *   `themes`: Definitions for custom UI themes to personalize the CLI experience.
    *   **Skills-Focused Extensions**: For extensions primarily offering skills, this file can be minimal, relying on directory structure for skill discovery.

-   **`GEMINI.md`**:
    *   **Purpose**: Provides broader, overarching contextual instructions to the LLM when operating within the extension's scope.
    *   **Content**: Typically contains the extension's personality, general tooling knowledge, and descriptive notes about restrictions or guiding principles.
    *   **Discovery**: The CLI hierarchically searches for `GEMINI.md` files (current directory, parent directories, global context) to load relevant context.

-   **`skills/` Directory**:
    *   **Purpose**: Contains individual skill definitions that instruct the LLM on specific behaviors or responses.
    *   **Structure**: Skills are often organized within nested subdirectories (e.g., `skills/my-skill/SKILL.md`).

-   **`SKILL.md` (within `skills/`)**:
    *   **Purpose**: The dedicated markdown file for defining a specific skill's behavior and instructions for the LLM.
    *   **Content**: Includes a YAML front matter (`name`, `description`) and a markdown body detailing the skill's workflow, expected inputs, desired outputs, and how it aligns with the agent's persona.

-   **`commands/` Directory**:
    *   **Purpose**: Defines custom commands that can augment the LLM's prompt with dynamic information.
    *   **Structure**: Can organize commands into subcategories (e.g., `commands/fs/`).
    *   **Definition**: Custom commands are often defined using `.toml` files (e.g., `grep-code.toml`).
    *   **Content (TOML)**: A TOML file typically defines a `prompt` field that can embed shell commands using `!{}` syntax and arguments using `{{args}}`, feeding the command's output into the LLM's context.

-   **`hooks/` Directory**:
    *   **Purpose**: Provides a mechanism for defining and executing custom scripts in response to specific CLI lifecycle events (hooks).
    *   **Structure**: Contains `hooks.json` (for defining hooks) and a `scripts/` directory (for associated executable scripts).
    *   **`hooks.json` Content**: Specifies hook points (e.g., `SessionStart`) and the commands/scripts to execute, often using `${extensionPath}` for relative pathing.

## Key Learnings for Extension Design

1.  **Declarative First**: Whenever possible, define extension functionalities declaratively in `gemini-extension.json` (for tool exclusion, themes, MCP servers) or `.toml` files (for custom commands). This enhances readability and maintainability.
2.  **Modularity**: Organize skills, commands, and scripts into logical subdirectories for clarity and scalability.
3.  **Context vs. Skill**: Understand the distinction between `GEMINI.md` (broader contextual guidance for the LLM) and `SKILL.md` (specific instructions for LLM behavior in response to a task).
4.  **Active Enforcement**: For restrictions, prefer programmatic enforcement via `excludeTools` in `gemini-extension.json` over just descriptive notes in `GEMINI.md`.
5.  **Dynamic Augmentation**: Use custom commands to dynamically inject information from shell commands or other sources directly into the LLM's prompt.
6.  **Event-Driven Behavior**: Leverage hooks to automate tasks or react to CLI lifecycle events.

## Actions

-   When creating a Gemini CLI extension, prioritize `gemini-extension.json` for metadata and declarative features.
-   Define specific LLM-facing skills using `SKILL.md` files within `skills/` subdirectories.
-   Utilize `GEMINI.md` to establish the overall context, personality, and general guidelines for the extension.
-   Implement tool restrictions using the `excludeTools` array in `gemini-extension.json` for active enforcement.
-   Consider custom commands (using `.toml` files) to integrate shell utilities or other prompt-augmenting logic.
-   Explore hooks for event-driven automation and MCP servers for exposing new programmatic tools when advanced functionality is required.
