package config

type Config struct {
	Discord struct {
		UserName string `json:"username"`
		Token    string `json:"token"`
	} `json:"discord"`
	Search  Search `json:"cse"`
	Command struct {
		ErrorMessage string
		FoodPorn     Command `json:"foodPorn"`
		Welcome      Command `json:"welcome"`
	} `json:"command"`
}

type Search struct {
	Key string `json:"key"`
	Cx  string `json:"id"`
}

type Command struct {
	Keywords []string `json:"keywords"`
	Messages []string `json:"messages"`
}