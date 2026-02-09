# Topic: MCPorter (Model Context Protocol Toolkit)

## Metadata
- **Date**: 2026-02-09
- **Tags**: MCP, Claude, Gemini, TypeScript, AI Tools, Automation
- **Status**: Completed
- **Related Memories**: None

## Context/Problem
Investigation into the tool "mcporter" and its role in the ecosystem of AI models like Claude and Gemini. The user wanted to understand how it is used and how it relates to these models.

## Solution/Findings
MCPorter is a specialized toolkit developed by `steipete` (Peter Steinberger) designed for the Model Context Protocol (MCP). It serves as a TypeScript runtime, CLI utility, and code generation toolkit.

### 1. The Mental Model
- **"Headless" MCP Host**: MCPorter provides a way to interact with MCP servers without needing a heavy AI UI (like Claude Desktop). It treats MCP tools as composable CLI utilities or library functions.
- **The Orchestration Layer**: It sits between the raw MCP server and the final AI integration, providing a layer for debugging, testing, and automating tool workflows in pure TypeScript or shell.
- **Server Instructions (Self-Documentation)**: Modern MCP servers (like `everything`) provide embedded instructions for LLMs. MCPorter preserves and can display these, helping agents understand cross-tool relationships and operational constraints.
- **Transport Flexibility**: It abstracts away whether a server is running via `stdio` (local process) or `http` (remote service), allowing for seamless migration from local development to cloud-hosted tools.

### 2. Core Reference & Cheat Sheet
| Command | Description | Key Options |
| :--- | :--- | :--- |
| `mcporter list` | List available tools on a server | `--schema` (show JSON inputs) |
| `mcporter call` | Execute a specific tool | `--args='{"key": "val"}'` |
| `mcporter config` | Manage server definitions | `stdio` vs `http` transports |
| `mcporter generate-cli` | Create a standalone CLI wrapper | Requires `bun` for binary build |

### 3. Agent-Specific Guidance
- **JSON Payload Safety**: When an agent calls `mcporter`, it must strictly adhere to the JSON schema found in `mcporter list --schema`.
- **Escaping Complexity**: In shell environments, use single quotes for the `--args` string to avoid shell expansion issues with nested JSON double quotes.
- **Workflow Automation**: Agents can use MCPorter to "chain" tools together by taking the stdout of one `mcporter call` and passing it as input to the next.

## Actionable Insights/Next Steps
- Use `mcporter list --schema` to verify tool signatures before writing integration code.
- Generate dedicated CLI wrappers for frequently used tools to simplify shell scripts.
- Use the TypeScript API for complex "agent-less" automations.

## References
- GitHub Repository: [steipete/mcporter](https://github.com/steipete/mcporter)
- Model Context Protocol (MCP) Official Documentation
- Sample generated CLI: `everything.ts` (in project root)
