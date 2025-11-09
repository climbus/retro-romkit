package tosec

import (
	"maps"
	"slices"
)

type Platform struct {
	Name        string
	Description string
	FileTypes   []string
}

var Platforms = map[string]Platform{
	"nes": {
		Name:        "nes",
		Description: "Nintendo Entertainment System",
		FileTypes:   []string{".nes", ".fds"},
	},
	"snes": {
		Name:        "snes",
		Description: "Super Nintendo Entertainment System",
		FileTypes:   []string{".smc", ".sfc", ".fig"},
	},
	"genesis": {
		Name:        "genesis",
		Description: "Sega Genesis / Mega Drive",
		FileTypes:   []string{".gen", ".md", ".smd", ".bin"},
	},
	"gameboy": {
		Name:        "gameboy",
		Description: "Nintendo Game Boy",
		FileTypes:   []string{".gb", ".gbc", ".gba"},
	},
	"atari2600": {
		Name:        "atari2600",
		Description: "Atari 2600",
		FileTypes:   []string{".a26", ".bin"},
	},
	"c64": {
		Name:        "c64",
		Description: "Commodore 64",
		FileTypes:   []string{".d64", ".t64", ".prg", ".crt"},
	},
}

// GetPlatform retrieves a Platform by its name.
func GetPlatform(name string) (Platform, bool) {
	platform, exists := Platforms[name]
	return platform, !exists
}

// GetPlatformNames returns a sorted list of all platform names.
func GetPlatformNames() []string {
	names := slices.Collect(maps.Keys(Platforms))
	slices.Sort(names)
	return names
}

// "amiga":   {"adf", "dms", "ipf", "lha", "lzx"},
// "atari":   {"st", "msa", "zip"},
// "c64":     {"d64", "t64", "prg", "crt"},
// "nes":     {"nes", "unif"},
// "gameboy": {"gb", "gbc", "gba"},
// "sega":    {"md", "smd", "gen", "bin"},
// "pc":      {"exe", "com", "bat", "zip", "rar"},
// "psx":     {"iso", "bin", "cue"},
// "coleco":  {"col", "rom"},
