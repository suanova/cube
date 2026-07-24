// Package env provides a single place to define environment variables that are
// exported to child processes (Bash tool, hooks, MCP servers, etc.).
//
// Every CUBE_* variable is also emitted as a CLAUDE_* alias (so Claude Code
// plugins work unmodified) and, during the fork transition, a SAN_* alias (so
// scripts keyed on the old name keep working). Reads accept any of the three,
// preferring CUBE_, then SAN_, then CLAUDE_.
package setting

import (
	"fmt"
	"os"
)

const (
	prefix       = "CUBE_"
	legacyPrefix = "SAN_"    // pre-fork name; still read and re-emitted for compat
	aliasPrefix  = "CLAUDE_" // Claude Code compatibility alias
)

// aliasPrefixes are the extra prefixes emitted alongside the canonical CUBE_
// variant. CLAUDE_ keeps Claude Code plugins working; SAN_ keeps scripts from
// the pre-fork name working.
var aliasPrefixes = []string{legacyPrefix, aliasPrefix}

// EnvPair creates env entries for a single key=value, returning the canonical
// CUBE_ variant plus the SAN_ and CLAUDE_ aliases.
//
//	EnvPair("PROJECT_DIR", "/tmp") →
//	  ["CUBE_PROJECT_DIR=/tmp", "SAN_PROJECT_DIR=/tmp", "CLAUDE_PROJECT_DIR=/tmp"]
func EnvPair(key, value string) []string {
	out := make([]string, 0, 1+len(aliasPrefixes))
	out = append(out, prefix+key+"="+value)
	for _, a := range aliasPrefixes {
		out = append(out, a+key+"="+value)
	}
	return out
}

// EnvPairs creates env entries for multiple key=value pairs.
func EnvPairs(kvs ...string) []string {
	if len(kvs)%2 != 0 {
		panic("config.EnvPairs: odd number of arguments")
	}
	out := make([]string, 0, len(kvs)/2*(1+len(aliasPrefixes)))
	for i := 0; i < len(kvs); i += 2 {
		out = append(out, EnvPair(kvs[i], kvs[i+1])...)
	}
	return out
}

// EnvPairF is like EnvPair but with a formatted suffix on the key.
//
//	EnvPairF("PLUGIN_ROOT_%s", "CODEX", "/path") →
//	  ["CUBE_PLUGIN_ROOT_CODEX=/path", "SAN_PLUGIN_ROOT_CODEX=/path", "CLAUDE_PLUGIN_ROOT_CODEX=/path"]
func EnvPairF(keyFmt, keyArg, value string) []string {
	key := fmt.Sprintf(keyFmt, keyArg)
	return EnvPair(key, value)
}

// Getenv reads the canonical CUBE_<suffix> variable, falling back to the
// legacy SAN_<suffix> and then CLAUDE_<suffix> if the canonical name is unset.
func Getenv(suffix string) string {
	if v, ok := os.LookupEnv(prefix + suffix); ok {
		return v
	}
	if v, ok := os.LookupEnv(legacyPrefix + suffix); ok {
		return v
	}
	return os.Getenv(aliasPrefix + suffix)
}
