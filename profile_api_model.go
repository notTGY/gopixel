package gopixel

type ProfileResponse struct {
	Success  bool      `json:"success"`
	Cause    string    `json:"cause"`
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	ProfileID string            `json:"profile_id"`
	CuteName  string            `json:"cute_name"`
	Members   map[string]Member `json:"members"`
}

type Collection struct {
	Oak     int64 `json:"LOG"`
	Spruce  int64 `json:"LOG:1"`
	Birch   int64 `json:"LOG:2"`
	Jungle  int64 `json:"LOG:3"`
	Acacia  int64 `json:"LOG_2"`
	DarkOak int64 `json:"LOG_2:1"`
}

type Experience struct {
	Foraging float64 `json:"SKILL_FORAGING"`
}

type PlayerData struct {
	Experience Experience `json:"experience"`
}

type Member struct {
	Collection Collection `json:"collection"`
	PlayerData PlayerData `json:"player_data"`
}
