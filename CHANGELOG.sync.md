# Upstream Sync Changelog

## 2026-08-05 — 24 commits from san#main

| SHA | Intent | Type | Risk |
|-----|--------|------|------|
| `c1578bf085a7` | Fix persona selector metadata truncation width to prevent rows from overflowing the panel | bugfix | low |
| `d896ee854097` | Restructure README documentation around three core value pillars (small, fast, open) with measured benchmarks, replacing flat feature lists with quantified performance claims | internal | low |
| `ca1fa7140778` | Fix composer cursor misalignment on wrapped or multi-line input by indenting all rows after the first to match the prompt width | bugfix | low |
| `bf96ef64c792` | Expand @ autocomplete to show all file types and respect .gitignore patterns while increasing scan limits for deeper project trees | feature | medium |
| `2ec655daa7cf` | Extend ProcessImageRefs to also detect bare image file paths (e.g., from drag-drop into the terminal) in addition to @-prefixed references, with more lenient semantics for bare paths (text preserved, failed loads silently skipped). | feature | medium |
| `0c8a6ea72f8c` | Rebuild the documentation site homepage and design system to match the README's updated positioning ('Minimal overhead. Maximum agent.'), including a scripted terminal demo, trigram-based 'Why San' section, and intro.html scaling fixes. | refactor | low |
| `d57c2b9ff4a7` | Update GitHub Actions checkout action from a pinned SHA (v7.0.0) to the v4 tag in the release-bot workflow | internal | low |
| `4132b6ec606a` | Strip image attachments from conversation history when seeding agent messages for text-only models, preventing providers from rejecting replayed image_url parts after a model switch | bugfix | medium |
| `f07c90601a8a` | Fix cancel (Esc) to actually stop the agent loop: re-check ctx between sequential tool calls, hold turn-queue drains on StopCancelled, and clear interruptPending on new user messages. | bugfix | medium |
| `708461a0aa49` | Release version 1.22.1 with changelog updates documenting new features, changes, and fixes since v1.22.0 | internal | low |
| `ac820629dd11` | Keep the assistant's streamed text visible above docked permission modals and freeze animated indicators (tool spinners, assistant bullets, agent rows) while the modal awaits user input | bugfix | medium |
| `19bcffb8fdad` | Deliver background task results into a running conversation as soon as it is safe to append, instead of parking them until the turn boundary, and rename conv.Runtime members by origin (On*/Handle*/verb) for clarity. | feature | medium |
| `e45ec0ef7e48` | Update documentation to match the actual implementation: correct the prompt slot table, permission mode table, and dead source file references | internal | low |
| `ff91421a0af2` | Bump actions/stale GitHub Action from v10.4.0 to v11.0.0 in the stale workflow | internal | low |
| `17c9402f9e9b` | Bump actions/checkout from v4 to v7 in GitHub Actions workflows, updating pinned SHAs to v7.0.1 in ci.yml and the major version tag in release-bot.yml. | internal | low |
| `9d9b4d7d1a6d` | Fix agent behavior prompt to prevent premature stopping by adding explicit persistence instruction and narrowing the scope of the 'recommend don't implement' guidance to only open-ended exploratory questions | bugfix | low |
| `ab6833ed0dec` | Fix TUI rendering bug where full-width separator rows caused orphaned rule rows above the composer after terminal resize by making separators one column short to avoid deferred-wrap state ambiguity | bugfix | low |
| `85ec4b86e12b` | Report tool calls waiting on permission prompts as waiting rather than running, by carrying the tool call ID through the permission gate and using it to distinguish gated calls from executing ones in the render path. | bugfix | medium |
| `01ce7256b641` | Fix a security bypass where a PreToolUse hook's 'allow' decision skipped the permission gate entirely, by wiring in Settings.ResolveHookAllow so hook allows are vetted against deny rules, circuit breaker, confirmation tiers, and ask rules. | bugfix | medium |
| `fe74e72b77ae` | Allow text-only models to receive image paths inline instead of blocking image-bearing messages, while fixing slash-command misdetection of leading image paths and correcting DeepSeek thinking-effort options | feature | medium |
| `c6bd60a3fbd0` | Fix clipboard image temp files being deleted at turn end (breaking follow-up turns that reference the inlined path), honor ProcessImageRefs' error contract by returning surviving text/images on failure, and make image extension selection deterministic. | bugfix | medium |
| `861f28833c53` | Fix resumed sessions demoting permission mode to normal and enable shift+tab mode cycling during active turns, with thread-safe posture reads | bugfix | medium |
| `d7647f94cfad` | Fix TUI renderer to prevent orphan rows after terminal resize by correcting row count calculations for rewrapped content and restoring one-column separator slack | bugfix | medium |
| `58fafbb63cfe` | Remove the GitHub Actions stale workflow in favor of using san-ci to manage stale issues and PRs | internal | low |
## 2026-07-31 — 10 commits from san#main

| SHA | Intent | Type | Risk |
|-----|--------|------|------|
| `c0d5fb295c74` | Add a configurable OpenAI/Claude-compatible custom LLM provider with inline UI setup, persisted base URL/API-key auth, model listing, and related default-model/URL-normalization fixes. | feature | medium |
| `fda00fe50635` | Simplify the agent/subagent runtime model by making the general subagent the implicit default and selecting behavior via modes such as default/explore instead of explicit runtime/persona names. | feature | medium |
| `187a89cc2966` | Clarifies agent runtime selection terminology by renaming internal request/selection resolver APIs and related references without changing behavior. | refactor | low |
| `37a0c1ed9d2d` | Remove the separate agent/subagent type concept so agents are selected and represented without a type discriminator. | breaking | high |
| `158df2712b1d` | Align Autopilot suggestion behavior with its settings/UI so suggestions are enabled by default and apply outside Autopilot mode, while preserving full Bash commands in tool rendering. | bugfix | medium |
| `2e045b6625f6` | Unknown requested agent names are now accepted as display-only labels while falling back to the base agent configuration, with disabled known agents still rejected. | bugfix | medium |
| `b46072f5c9b2` | Fix MCP client adoption/connection lifecycle handling so retained or transferred clients notify the correct registry and are not dropped on reconfiguration or cwd changes. | bugfix | medium |
| `ca043bd831fe` | Ensure subagents respect disabled tool filters and reduce unnecessary Agent-tool delegation, while removing a duplicate simplify skill. | bugfix | medium |
| `f82097e05ee5` | Add first-class Ollama provider support in the provider selector UI, including a dedicated base-URL form instead of API-key editing. | feature | medium |
| `8e9e115379c7` | Rename the model-selection slash command from /model to /models, remove the stale /glob slash command, and update documentation/version references accordingly. | breaking | high |
## 2026-07-31 — 10 commits from san#main

| SHA | Intent | Type | Risk |
|-----|--------|------|------|
| `4703695a8090` | Simplify TUI tool-call rendering by using generic Read labels and removing extra Bash connector markers in collapsed/long output. | bugfix | low |
| `d2f28399428f` | Fix TUI rendering of multiline Bash tool calls so physical command lines are visually connected while soft-wrapped lines still align under the prompt. | bugfix | low |
| `d6a5cbe66bb5` | Simplify internal TUI rendering for nested tool results, including cleaner bash command/result alignment and avoiding connector-only rows for blank command lines. | refactor | low |
| `87f916a10030` | Fix Bash tool rendering so wrapped command/output lines keep TUI connector prefixes continuous in live and transcript views. | bugfix | medium |
| `2ddb8ff0647b` | Fix TUI rendering so single-line Bash tool calls stay compact, using truncated previews and integrating row detail without duplicate rendering. | bugfix | medium |
| `fb7be3d55d98` | Refine TUI Bash tool-call preview rendering so the command retains tool-call styling while only the description is dimmed, with tests updated for the new styling behavior. | bugfix | low |
| `787de9f95af8` | Fix TUI scrollback handoff ordering by pinning a patched Bubble Tea renderer that flushes queued frames before immediate scrollback insertion, preventing stale live-frame rows from being embedded in native scrollback. | bugfix | medium |
| `74ac1647da05` | Fix TUI native scrollback handoff races by serializing queued scrollback print commits in FIFO order instead of relying on a timed delay. | bugfix | medium |
| `ce55709e39dd` | Fix TUI scrollback rendering so commits taller than the viewport do not corrupt or overwrite the terminal's native scrollback/history, including resize/update handling and regression coverage. | bugfix | medium |
| `0ecb2718fe6a` | Retry LLM stream/completion requests when provider streams fail with transient HTTP/2 stream errors instead of surfacing them as terminal failures. | bugfix | medium |
This file is auto-generated by [upstream-semantic-sync](https://github.com/upstream-semantic-sync). It records every upstream commit that was adopted into this downstream fork.

## 2026-07-31 — 10 commits from san#main

| SHA | Intent | Type | Risk |
|-----|--------|------|------|
| `b9c36d4cebab` | Change permission bypass semantics so bypass skips confirmation tiers but never skips a root/home recursive-removal circuit breaker, with matching main/subagent gate behavior. | bugfix | high |
| `2fe0240eedc8` | Ship /simplify as an embedded builtin prompt command that runs through the custom-command pipeline while allowing user/project commands to shadow it. | feature | medium |
| `1b181c33a9cd` | Split agent stopping into a dedicated AgentStop tool and harden background task termination via structural SIGTERM/process-group handling. | breaking | high |
| `98816b406098` | Fix TUI rendering so Bash results nested beneath their tool call appear as command output/terminal state instead of repeating the Bash label, while standalone Bash results still show the tool label. | bugfix | medium |
| `15a8438a9c44` | Fix TUI rendering so file/edit tool results are nested under their corresponding tool calls/paths and do not repeat the tool name. | bugfix | medium |
| `83cb4e1e0ee8` | Refines TUI tool-call/result rendering so Bash commands consistently use command blocks and Read tool labels are formatted more clearly. | bugfix | medium |
| `ca655f6beebf` | Fix TUI transcript rendering so nested Edit/Write tool result diffs use the caller-provided indentation and align with their result trailers. | bugfix | medium |
| `63f4dfe99a23` | Fix TUI rendering so visible tool result details remain visually connected/nested under their tool call and summary trailer. | bugfix | low |
| `77c42e2d22d8` | Shorten and tighten TUI tool/file-change rendering output, including removing extra spacing after Bash prompts. | bugfix | low |
| `320ebe4bc93b` | Fix TUI Bash tool result rendering so collapsed output still shows the connector line between the command and summarized result. | bugfix | medium |
