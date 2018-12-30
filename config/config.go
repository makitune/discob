package config

type Config struct {
	Discord struct {
		UserName string `json:"username"`
		Token    string `json:"token"`
	} `json:"discord"`
	Search  Search `json:"cse"`
	Command struct {
		ErrorMessage       string
		FoodPorn           Command `json:"foodporn"`
		HeadsUp            Command `json:"headsup"`
		Welcome            Command `json:"welcome"`
		JoinVoiceChannel   Command `json:"joinVoiceChannel"`
		DefectVoiceChannel Command `json:"defectVoiceChannel"`
	} `json:"command"`
}

type Search struct {
	Key       string `json:"key"`
	Cx        string `json:"id"`
	OutputDir string `json:"outputDir"`
}

type Command struct {
	Keywords []string `json:"keywords"`
	Messages []string `json:"messages"`
}
