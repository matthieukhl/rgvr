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
