# Fork Conventions: cube vs. upstream (san)

This document explains how the **cube** fork diverges from its upstream
([genai-io/san](https://github.com/genai-io/san)) and the conventions that
govern that divergence.

## Summary

The only substantive divergence is a **branding rename**: everything
user-facing was changed from **san** (三) to **cube**. No features have been
added, removed, or restructured — the codebase is functionally identical to
upstream.

The rename was applied selectively: user-visible identity changed, but the
Go module path and import paths were deliberately kept as-is to avoid a
mass, break-every-import migration.

## Rename mapping

| Layer | san (upstream) | cube (this fork) | Notes |
|---|---|---|---|
| Binary / CLI | `san` | `cube` | `cmd/cube/` |
| Go module path | `github.com/genai-io/san` | **unchanged** | Kept to preserve all imports |
| Go import paths | `github.com/genai-io/san/internal/…` | **unchanged** | Follows module path |
| User config dir | `~/.san/` | `~/.cube/` | `~/.san/` still read as fallback |
| Project config dir | `.san/` | `.cube/` | `.san/` still read as fallback |
| Plugin metadata dir | `.san-plugin/` | `.cube-plugin/` | `.san-plugin/` still read as fallback |
| Environment variables | `SAN_*` | `CUBE_*` | `SAN_*` still read as fallback; `CLAUDE_*` also accepted |
| Instruction files | `SAN.md`, `SAN.local.md` | `CUBE.md`, `CUBE.local.md` | Both loaded; `CUBE.md` takes priority |
| Git remote | — | origin = `suanova/cube`, upstream = `genai-io/san` | Upstream is fetch-only |
| CI / install scripts | `genai-io/san` | `suanova/cube` | Release bot, install.sh, install.ps1 |
| Branding | `< SAN ✦ />` | `< CUBE ✦ />` | README, site |

## Backward-compatibility fallbacks

The fork intentionally retains fallback support for the old `san` names so
that existing user configurations keep working without migration:

- **Config dirs**: `confdir.Dir()` resolves `~/.cube/` first, falls back to
  `~/.san/`.
- **Env vars**: `setting.Getenv()` checks `CUBE_` → `SAN_` → `CLAUDE_`.
- **Env emission**: `setting.EnvPair()` emits all three prefixes so child
  processes see whichever they expect.
- **Instruction files**: The memory loader reads both `CUBE.md` and `SAN.md`;
  cube-prefixed files take priority when both exist.
- **Plugin dirs**: The plugin loader checks `.cube-plugin/` first, then
  `.san-plugin/`.

## Upstream sync

An automated GitHub Action (`.github/workflows/upstream-sync.yml`) runs weekly
(Monday 06:00 UTC) and creates a PR to merge upstream changes. The base commit
the fork was created from is anchored at `f83363c6`.

Because the Go module path is unchanged, most upstream merges apply cleanly.
The primary conflict surface is any file that was renamed in the
`10cfaec5` ("rename san to cube") commit — mostly `cmd/cube/` and the
user-facing constants in `confdir`, `setting`, and `plugin` packages.

## Known stale references

The rename was not exhaustive. The following still reference "san" and should
be updated over time (mostly comments and the MCP client name):

- `internal/mcp/client.go` — `ClientName = "san"` (sent in MCP `initialize`)
- Several doc comments still say `~/.san` or `.san/` instead of `~/.cube` /
  `.cube/` (in `setting/loader.go`, `secret/store.go`, `skill/registry.go`,
  `selflearn/skill.go`, `app/kit/save_level.go`, `command/registry.go`)
- `site/` — the entire static site is still branded as SAN
- `assets/san-intro.gif`, `assets/san.png` — asset filenames not renamed

These are cosmetic or protocol-level issues and do not affect runtime
behaviour due to the fallback mechanisms described above.

## Conventions for future changes

1. **Use `cube` for all new user-facing names.** New env vars, config keys,
   file names, and CLI flags should use the `CUBE_` / `.cube` prefix.
2. **Preserve fallback reading of `san` names** when adding new config paths
   or env vars, following the pattern established in `confdir` and `setting`.
3. **Do not rename the Go module path.** It stays as
   `github.com/genai-io/san` to keep imports stable and simplify upstream
   merges.
4. **When merging upstream,** watch for changes in files touched by commit
   `10cfaec5` — these are the most likely conflict sites. Resolve conflicts
   in favour of `cube` naming.
5. **Fix stale `san` references** when modifying nearby code, but do not
   open rename-only PRs for comments unless they could genuinely confuse
   a reader.
