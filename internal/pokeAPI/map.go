package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (Locations, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {

		var location Locations
		if err := json.Unmarshal(data, &location); err != nil {
			return Locations{}, fmt.Errorf("error unmarshalling response body: %w", err)
		}

		return location, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Locations{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Locations{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	var location Locations
	if err = json.Unmarshal(body, &location); err != nil {
		return Locations{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	c.cache.Add(fullURL, body)

	return location, nil

}
