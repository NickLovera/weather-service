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

type IPointsRepo interface {
	GetMetaDataByLatLong(c context.Context, lat, long float64) (*models.PointMetaData, error)
}

type pointsRepo struct{}

func NewPointsRepo() IPointsRepo {
	return &pointsRepo{}
}

func (ws *pointsRepo) GetMetaDataByLatLong(c context.Context, lat, long float64) (*models.PointMetaData, error) {

	pointUrl := fmt.Sprintf("https://api.weather.gov/points/%f,%f", lat, long)
	req, err := http.NewRequest("GET", pointUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create points API request. Err: %s", err)
	}

	//Pull out user agent
	req.Header.Set("User-Agent", contextKey.GetUserAgent(c))

	pointResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call points API. Err: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			//This would be an error level log
			log.Printf("failed to close points API response body. Err: %s\n", err)
		}
	}(pointResp.Body)

	var pointMetaData models.PointMetaData
	if err := json.NewDecoder(pointResp.Body).Decode(&pointMetaData); err != nil {
		return nil, fmt.Errorf("failed to decode points API response. Err: %s", err)
	}

	return &pointMetaData, nil
}
