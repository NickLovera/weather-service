package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/NickLovera/weather-service/models"
)

type IPointsRepo interface {
	GetMetaDataByLatLong(c context.Context, lat, long float64) (*models.PointMetaData, error)
}

type pointsRepo struct {
	client *http.Client
}

func NewPointsRepo(client *http.Client) IPointsRepo {
	return &pointsRepo{
		client: client,
	}
}

func (ws *pointsRepo) GetMetaDataByLatLong(c context.Context, lat, long float64) (*models.PointMetaData, error) {

	pointUrl := fmt.Sprintf("https://api.weather.gov/points/%f,%f", lat, long)
	req, err := http.NewRequestWithContext(c, "GET", pointUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create points API request. Err: %w", err)
	}

	req.Header.Set(
		"User-Agent",
		"weather-service/1.0 (github.com/NickLovera/weather-service)",
	)

	pointResp, err := ws.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call points API. Err: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			//This would be an error level log
			log.Printf("failed to close points API response body. Err: %s\n", err)
		}
	}(pointResp.Body)

	body, err := io.ReadAll(pointResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read points API response body. Err: %w", err)
	}

	if pointResp.StatusCode < http.StatusOK || pointResp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("NWS points API returned status %d: %s", pointResp.StatusCode, string(body))
	}

	var pointMetaData models.PointMetaData
	if err := json.Unmarshal(body, &pointMetaData); err != nil {
		return nil, fmt.Errorf("failed to decode points API response. Err: %w", err)
	}

	return &pointMetaData, nil
}
