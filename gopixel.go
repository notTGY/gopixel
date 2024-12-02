package gopixel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func QueryPlayerApi(UUID string, token string) (PlayerResponse, error) {
	var playerResponse PlayerResponse
	url := fmt.Sprintf("https://api.hypixel.net/v2/player?uuid=%s", UUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return playerResponse, err
	}
	req.Header.Set("API-Key", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return playerResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return playerResponse, err
	}

	err = json.Unmarshal(body, &playerResponse)
	if err != nil {
		return playerResponse, err
	}
	return playerResponse, nil
}

func QuerySkyblockProfileApi(UUID string, token string) (ProfileResponse, error) {
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

	err = json.Unmarshal(body, &profileResponse)
	if err != nil {
		return profileResponse, err
	}
	return profileResponse, nil
}
