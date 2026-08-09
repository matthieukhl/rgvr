package client

import (
	"encoding/json"
	"fmt"

	"github.com/matthieukhl/rgvr/internal/models"
)

// GetNumbers retrieves the list of numbers associated with the authenticated user.
func (c *Client) GetNumbers() (*models.ListResponse[models.Number], *RequestInfo, error) {
	path := "/numbers"

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, nil, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	var numbersResponse models.ListResponse[models.Number]

	err = json.NewDecoder(resp.Body).Decode(&numbersResponse)
	if err != nil {
		return nil, nil, fmt.Errorf("decoding numbers information: %w", err)
	}

	return &numbersResponse, reqInfo, nil
}

func (c *Client) GetNumber(number string) (*models.Number, *RequestInfo, error) {
	path := fmt.Sprintf("/numbers/%s", number)

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		return nil, nil, fmt.Errorf("invalid phone number format: %s. The phone number in E.164 format without the + prefix (e.g., 33140000000 for a French number).", number)
	}

	if resp.StatusCode == 401 {
		return nil, nil, fmt.Errorf("unauthorized: invalid or missing API key, or missing 'Numbers Read' permission: %s", resp.Status)
	}

	if resp.StatusCode == 404 {
		return nil, nil, fmt.Errorf("number not found or not accessible with current monitoring level: %s", resp.Status)
	}

	if resp.StatusCode != 200 {
		return nil, nil, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	var numberResponse models.Number

	err = json.NewDecoder(resp.Body).Decode(&numberResponse)
	if err != nil {
		return nil, nil, fmt.Errorf("decoding number information: %w", err)
	}

	return &numberResponse, reqInfo, nil
}
