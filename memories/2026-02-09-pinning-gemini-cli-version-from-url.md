# Pinning gemini-cli version with npx from GitHub URL

**Date:** 2026-02-09

**Topic:** How to pin a specific version of gemini-cli when running it via `npx` from a GitHub URL.

**Tags:** gemini-cli, npx, npm, version pinning, github url

## Research Findings

When using `npx` to execute commands from a GitHub URL, version pinning is not as straightforward as with direct npm package names. However, `npx` often utilizes `npm exec` under the hood, which supports version pinning.

The general syntax for pinning a version with `npm exec` is:
- For an npm package: `npm exec <package_name>@<version>`
- For a package from a URL: `npm exec --package=<package_url>@<version> -- <command>`

Therefore, to run a specific version of `gemini-cli` from its GitHub repository, one could use a command similar to:
`npx npm exec --package=https://github.com/google-gemini/gemini-cli@<specific_version> -- gemini-cli`

This command tells `npm exec` to fetch and use the specified version of the `gemini-cli` repository.

## Actions

- To run a specific version of a package from a GitHub URL using `npx`, append the desired version to the URL within the `--package` argument of `npm exec`: `npx npm exec --package=<github_url>@<version> -- <command>`.
- For published npm packages, the syntax is simpler: `npm exec <package_name>@<version>`.
