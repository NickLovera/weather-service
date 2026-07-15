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

	if lat < -90 || lat > 90 || long < -180 || long > 180 {
		return nil, fmt.Errorf("invalid latitude or longitude values. Latitude must be between -90 and 90, and longitude must be between -180 and 180")
	}

	pointMetaData, err := ws.PointsRepo.GetMetaDataByLatLong(c, lat, long)
	if err != nil {
		return nil, fmt.Errorf("failed to get point metadata. Err: %w", err)
	}

	gridMetaData, err := ws.GridPointsRepo.GetHourlyForeCast(c, pointMetaData.Properties.GridId, pointMetaData.Properties.GridX, pointMetaData.Properties.GridY)
	if err != nil {
		return nil, fmt.Errorf("failed to get grid point metadata. Err: %w", err)
	}

	var (
		currentForecast string
		currentTemp     float64
	)

	if len(gridMetaData.Properties.Periods) > 0 {
		currentTemp = gridMetaData.Properties.Periods[0].Temperature
		currentForecast = gridMetaData.Properties.Periods[0].ShortForecast
	} else {
		return nil, fmt.Errorf("no forecast periods available for the given location")
	}

	return &models.CurrentWeather{
		City:               pointMetaData.Properties.RelativeLocation.Properties.City,
		State:              pointMetaData.Properties.RelativeLocation.Properties.State,
		Forecast:           currentForecast,
		TempCharacteristic: determineTempCharacteristic(currentTemp),
	}, nil
}
