package icons

import "strings"

type Icon struct {
	Symbol string
	Color  string
}

var iconMap = map[string]Icon{
	".gitignore":          {Symbol: "\uf1d3", Color: "#f1502f"},
	".go":                 {Symbol: "\ue627", Color: "#00add8"},
	".py":                 {Symbol: "\ue73c", Color: "#3572a5"},
	"dockerfile":          {Symbol: "\uf308", Color: "#0db7ed"},
	"docker-compose.yaml": {Symbol: "\uf308", Color: "#ff5555"},
	"docker-compose.yml":  {Symbol: "\uf308", Color: "#ff5555"},
	".js":                 {Symbol: "\ue74e", Color: "#f7df1e"},
	".ts":                 {Symbol: "\ue628", Color: "#3178c6"},
	".json":               {Symbol: "\ue60b", Color: "#cbcb41"},
	".yaml":               {Symbol: "\ue73e", Color: "#ccaa00"},
	".yml":                {Symbol: "\ue73e", Color: "#ccaa00"},
	".md":                 {Symbol: "\uf48a", Color: "#83a598"},
	".rs":                 {Symbol: "\ue7a8", Color: "#dea584"},
	".cpp":                {Symbol: "\ue61d", Color: "#9f1853"},
	".sh":                 {Symbol: "\ue757", Color: "#89e051"},
	".lua":                {Symbol: "\ue620", Color: "#000080"},
}

var defaultIcon = Icon{Symbol: "\uf15b", Color: "#a6accd"}

func GetIcon(filename string) Icon {
	filename = strings.ToLower(filename)
	for ext, icon := range iconMap {
		if filename == ext || strings.HasSuffix(filename, ext) {
			return icon
		}
	}
	return defaultIcon
}
