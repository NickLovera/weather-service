package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/NickLovera/weather-service/contextKey"
	"github.com/NickLovera/weather-service/models"
)

type IGridPointsRepo interface {
	GetHourlyForeCast(c context.Context, wfo string, x, y int) (*models.GridPointHourly, error)
}

type gridPointsRepo struct {
}

func NewGridPointsRepo() IGridPointsRepo {
	return &gridPointsRepo{}
}

func (gpr *gridPointsRepo) GetHourlyForeCast(c context.Context, wfo string, x, y int) (*models.GridPointHourly, error) {

	hourlyUrl := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast/hourly", wfo, x, y)
	req, err := http.NewRequest("GET", hourlyUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create hourly API request. Err: %s", err)
	}

	//Pull out user agent
	req.Header.Set("User-Agent", contextKey.GetUserAgent(c))

	pointResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call points API. Err: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			//This would be an error level log
			log.Printf("failed to close hourly API response body. Err: %s\n", err)
		}
	}(pointResp.Body)

	body, err := io.ReadAll(pointResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read hourly API response body. Err: %w", err)
	}

	if pointResp.StatusCode < http.StatusOK || pointResp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("NWS hourly forecast API returned status %d: %s", pointResp.StatusCode, string(body))
	}

	var pointMetaData models.GridPointHourly
	if err := json.Unmarshal(body, &pointMetaData); err != nil {
		return nil, fmt.Errorf("failed to decode hourly forecast API response. Err: %w", err)
	}

	return &pointMetaData, nil
}
