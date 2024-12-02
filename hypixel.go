package gopixel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

type HypixelAPIRequest struct {
	UUID        string
	SocialDepth int
}
type HypixelAPIResponse struct {
	UUID            string
	SocialDepth     int
	ProfileResponse ProfileResponse
}

type SocialMediaLinks struct {
	Discord string `json:"DISCORD"`
}
type SocialMedia struct {
	Links SocialMediaLinks `json:"links"`
}

type Player struct {
	SocialMedia SocialMedia `json:"socialMedia"`
}
type HypixelPlayerResponse struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	Player  Player `json:"player"`
}

func QueryHypixelPlayerApi(UUID string, token string) (HypixelPlayerResponse, error) {
	var hypixelPlayerResponse HypixelPlayerResponse
	url := fmt.Sprintf("https://api.hypixel.net/v2/player?uuid=%s", UUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return hypixelPlayerResponse, err
	}
	req.Header.Set("API-Key", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return hypixelPlayerResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return hypixelPlayerResponse, err
	}

	if err := json.Unmarshal(body, &hypixelPlayerResponse); err != nil {
		return hypixelPlayerResponse, err
	}
	return hypixelPlayerResponse, nil
}

func QueryHypixelApi(UUID string, token string) (ProfileResponse, error) {
	var profileResponse ProfileResponse
	url := fmt.Sprintf("https://api.hypixel.net/v2/skyblock/profiles?uuid=%s", UUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return profileResponse, err
	}
	req.Header.Set("API-Key", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return profileResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return profileResponse, err
	}

	if err := json.Unmarshal(body, &profileResponse); err != nil {
		return profileResponse, err
	}
	return profileResponse, nil
}
