package gopixel

type SocialMediaLinks struct {
	Discord string `json:"DISCORD"`
}
type SocialMedia struct {
	Links SocialMediaLinks `json:"links"`
}

type Player struct {
	SocialMedia SocialMedia `json:"socialMedia"`
}
type PlayerResponse struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	Player  Player `json:"player"`
}
