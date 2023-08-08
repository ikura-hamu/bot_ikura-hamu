package conf

import "os"

type Mode int

const (
	ProdMode Mode = iota
	DevMode
)

func GetMode() Mode {
	if len(os.Args) < 2 {
		return ProdMode
	}

	switch os.Args[1] {
	case "dev":
		return DevMode
	default:
		return ProdMode
	}
}

func GetModeName(mode Mode) string {
	switch mode {
	case ProdMode:
		return "Production"
	case DevMode:
		return "Development"
	default:
		return "Unknown"
	}
}
