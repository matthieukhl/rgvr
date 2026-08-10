package client

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (c *Client) AssignNumber(number string, target *models.NumberAssignment) (*RequestInfo, error) {
	path := fmt.Sprintf("/numbers/%s", number)

	body, err := json.Marshal(target)
	if err != nil {
		return nil, fmt.Errorf("encoding request body: %w", err)
	}

	resp, reqInfo, err := c.Patch(path, body)
	if err != nil {
		return reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return reqInfo, fmt.Errorf("%s: invalid request - check number format and assignment target", resp.Status)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return reqInfo, fmt.Errorf("%s: invalid API key or missing 'Numbers Read' permission", resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return reqInfo, fmt.Errorf("%s: number not found", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return reqInfo, fmt.Errorf("unexpected response from API: %w", err)
	}

	return reqInfo, nil
}
