package config

type Config struct {
	Length        int
	IncludeLower  bool
	IncludeUpper  bool
	IncludeNumbers bool
	IncludeSymbols bool
}

func DefaultConfig() Config {
	return Config{
		Length:        8,
		IncludeLower:  true,
		IncludeUpper:  true,
		IncludeNumbers: true,
		IncludeSymbols: false,
	}
}
