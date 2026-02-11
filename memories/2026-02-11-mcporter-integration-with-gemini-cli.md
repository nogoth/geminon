---
Topic: MCPorter Integration with Gemini CLI
Date: 2026-02-11
Tags: mcporter, gemini-cli, MCP, CLI generation, TypeScript, API, tool integration
Related Memories: None
---

## Summary

MCPorter (`mcporter`) is a TypeScript runtime, CLI, and code-generation toolkit designed to interact with Model Context Protocol (MCP) servers. Its core utility for Gemini CLI lies in its ability to extend Gemini CLI's capabilities by integrating with various MCP services, either as new CLI commands or as type-safe programmatic TypeScript clients.

## Key MCPorter Features Relevant to Gemini CLI Integration

1.  **CLI Generation (`mcporter generate-cli`)**:
    *   **Purpose**: Creates standalone CLIs from any MCP server definition. This is highly valuable for Gemini CLI as it enables exposing external MCP functionalities directly as Gemini CLI commands.
    *   **Input Flexibility**: Accepts server names (if configured), URLs, file paths to JSON definitions, inline JSON/JSON5 strings, and supports `stdio` commands.
    *   **Output Options**: Generates a TypeScript file (`.ts`) by default, with options to bundle into a single JavaScript file (`.js`) or compile into a Bun binary.
    *   **Customization**: Offers flags like `--name`, `--description`, `--include-tools`, and `--exclude-tools` for fine-grained control over the generated CLI.
    *   **Regeneration**: Generated CLIs embed metadata, facilitating `mcporter inspect-cli` for summaries and `mcporter generate-cli --from` for regenerating with updated `mcporter` versions, ensuring maintainability of integrated CLIs.

2.  **Typed Client Generation (`mcporter emit-ts`)**:
    *   **Purpose**: Produces TypeScript definition files (`.d.ts`) or complete client wrappers (`.ts`) for MCP servers. This is ideal for Gemini CLI's programmatic interactions with MCPs, ensuring strong type safety.
    *   **Modes**:
        *   `--mode types`: Generates a `.d.ts` interface with docblocks and promisified signatures.
        *   `--mode client`: Generates both the `.d.ts` and a `.ts` helper that wraps `createRuntime` and `createServerProxy`, providing a factory to create typed clients.
    *   **Benefits**: Promotes type safety, provides autocompletion, and aligns well with Gemini CLI's TypeScript codebase.

3.  **Tool Calling (CLI and TypeScript API)**:
    *   **CLI Calling**: `mcporter call <server.tool>` or `mcporter <server.tool>` allows direct invocation of MCP tools from the command line using various argument syntaxes (e.g., `key=value`, `key:value`, function-call syntax).
    *   **TypeScript API**:
        *   `callOnce`: For single, one-shot calls, handling discovery, OAuth, and transport closing automatically.
        *   `createRuntime()`: For repeated calls, offering connection pooling and advanced options while managing transports and OAuth token refreshes.
        *   `createServerProxy()`: Creates an ergonomic, type-friendly wrapper around `runtime.callTool`, facilitating argument mapping and method invocation.

4.  **Configuration Management**:
    *   `mcporter config`: Provides commands to manage `mcporter.json` files.
    *   **Resolution Order**: Prioritizes configuration from the `--config` flag, then `MCPORTER_CONFIG` environment variable, followed by project-level `config/mcporter.json`, and finally `~/.mcporter/mcporter.json[c]`, ensuring flexible configuration.
    *   **Dynamic Configuration**: Supports defining servers with `baseUrl`, `command`, `headers`, `env` variables, and `lifecycle` (e.g., `keep-alive` for daemon-managed servers).
    *   **OAuth**: Includes built-in caching and an `mcporter auth` command for simplified browser-based logins.

## How Gemini CLI Can Leverage MCPorter

*   **Extending Gemini CLI Functionality**: Gemini CLI can use `mcporter generate-cli` to dynamically create and integrate sub-CLIs for specific MCP servers, allowing users to interact with these external services directly through Gemini CLI.
*   **Internal Programmatic Tooling**: Gemini CLI's TypeScript codebase can utilize `mcporter emit-ts` to generate type-safe clients for any MCPs it needs to interact with programmatically. The `createRuntime()` and `createServerProxy()` APIs would be crucial for this.
*   **Dynamic MCP Discoverability**: Gemini CLI could integrate MCPorter's server discovery mechanisms to present users with a list of available MCPs on their system, enhancing user experience.
*   **Simplified Configuration Management**: Gemini CLI could provide a wrapper around `mcporter config` commands, simplifying the management of MCP server definitions for its users.

## Hypothetical Integration Steps for Gemini CLI

1.  **Add MCPorter as a Dependency**: Include `mcporter` in `gemini-cli`'s `package.json`.
2.  **Configuration Setup**: Define specific MCP servers via a `config/mcporter.json` within the Gemini CLI project, or leverage existing user-level `~/.mcporter/mcporter.json` configurations.
3.  **Develop Integration Points**:
    *   **For CLI Commands**: Implement a Gemini CLI command (e.g., `gemini-cli mcp generate <server-definition>`) that internally uses `mcporter generate-cli` to produce new executables, or directly integrates `mcporter.call` for direct tool invocation.
    *   **For Programmatic Interaction**: Gemini CLI's internal modules would import and utilize `mcporter`'s API (`createRuntime`, `createServerProxy`) after generating client wrappers using `mcporter emit-ts` for type-safe interactions.
