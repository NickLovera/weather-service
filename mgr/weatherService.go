package mgr

import (
	"context"
	"fmt"

	"github.com/NickLovera/weather-service/models"
	"github.com/NickLovera/weather-service/repo"
)

type IWeatherService interface {
	GetCurrentWeatherByLatLong(c context.Context, lat, long float64) (*models.CurrentWeather, error)
}

type weatherService struct {
	PointsRepo     repo.IPointsRepo
	GridPointsRepo repo.IGridPointsRepo
}

func NewWeatherService(pointsRepo repo.IPointsRepo, gridPointsRepo repo.IGridPointsRepo) IWeatherService {
	return &weatherService{
		PointsRepo:     pointsRepo,
		GridPointsRepo: gridPointsRepo,
	}
}

func (ws *weatherService) GetCurrentWeatherByLatLong(c context.Context, lat, long float64) (*models.CurrentWeather, error) {
	//Could add some validation on the lat/long so we don't make unnecessary call

	pointMetaData, err := ws.PointsRepo.GetMetaDataByLatLong(c, lat, long)
	if err != nil {
		return nil, fmt.Errorf("failed to get point metadata. Err: %s", err)
	}

	gridMetaData, err := ws.GridPointsRepo.GetHourlyForeCast(c, pointMetaData.Properties.GridId, pointMetaData.Properties.GridX, pointMetaData.Properties.GridY)
	if err != nil {
		return nil, fmt.Errorf("failed to get grid point metadata. Err: %s", err)
	}

	var (
		currentForecast = "No forecast data available" //Set defaults here, could error if no periods are given instead
		currentTemp     float64
	)

	if len(gridMetaData.Properties.Periods) > 0 {
		currentTemp = extractTemperature(gridMetaData.Properties.Periods[0].Temperature)
		currentForecast = gridMetaData.Properties.Periods[0].ShortForecast
	}

	return &models.CurrentWeather{
		City:               pointMetaData.Properties.RelativeLocation.Properties.City,
		State:              pointMetaData.Properties.RelativeLocation.Properties.State,
		Forecast:           currentForecast,
		Temperature:        currentTemp,
		TempCharacteristic: determineTempCharacteristic(currentTemp),
	}, nil
}

func extractTemperature(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case map[string]interface{}:
		if val, ok := v["value"].(float64); ok {
			return val
		}
		if val, ok := v["value"].(float32); ok {
			return float64(val)
		}
		if val, ok := v["value"].(int); ok {
			return float64(val)
		}
	}
	return 0
}
