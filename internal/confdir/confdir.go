// Package confdir resolves the per-user and per-project configuration directory.
//
// confdir is a zero-dependency infrastructure leaf so every layer (including
// internal/log and internal/core) can import it without a layering violation.
//
// Cube is a fork of san. The canonical config directory is .cube, but for
// backward compatibility with existing san installations, Dir falls back to
// the legacy .san directory when .cube does not yet exist. New writes always
// target .cube (see DirForWrite); reads resolve to whichever exists.
package confdir

import (
	"os"
	"path/filepath"
)

// Name is the canonical configuration directory name.
const Name = ".cube"

// LegacyName is the pre-fork configuration directory name (san). Dir falls
// back to it when the canonical directory does not exist, so existing san
// users keep their settings, personas, skills, and transcripts.
const LegacyName = ".san"

// Dir returns the active configuration directory under root. It prefers
// root/.cube and falls back to root/.san when .cube is absent, so a san
// installation upgrading to Cube keeps working without migration. When
// neither exists (a fresh install), it returns root/.cube so callers create
// the canonical directory going forward.
//
// Both reads and writes use Dir: a legacy user's data stays together in
// ~/.san, a fresh user's data goes to ~/.cube, and migration is opt-in via
// `mv ~/.san ~/.cube`.
func Dir(root string) string {
	canonical := filepath.Join(root, Name)
	if _, err := os.Stat(canonical); err == nil {
		return canonical
	}
	legacy := filepath.Join(root, LegacyName)
	if _, err := os.Stat(legacy); err == nil {
		return legacy
	}
	return canonical
}
