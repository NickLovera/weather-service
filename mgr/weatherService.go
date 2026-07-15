package mgr

import (
	"context"
	"fmt"

	"github.com/NickLovera/weather-service/models"
	"github.com/NickLovera/weather-service/repo"
)

type IWeatherService interface {
	GetCurrentWeatherByLatLong(c context.Context, isCelsius bool, lat, long float64) (*models.CurrentWeather, error)
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

func (ws *weatherService) GetCurrentWeatherByLatLong(c context.Context, isCelsius bool, lat, long float64) (*models.CurrentWeather, error) {
	//Could add some validation on the lat/long so we don't make unnecessary call

	pointMetaData, err := ws.PointsRepo.GetMetaDataByLatLong(c, lat, long)
	if err != nil {
		return nil, fmt.Errorf("failed to get point metadata. Err: %s", err)
	}

	gridMetaData, err := ws.GridPointsRepo.GetHourlyForeCast(c, isCelsius, pointMetaData.Properties.GridId, pointMetaData.Properties.GridX, pointMetaData.Properties.GridY)
	if err != nil {
		return nil, fmt.Errorf("failed to get grid point metadata. Err: %s", err)
	}

	var (
		currentForecast = "No forecast data available" //Set defaults here, could error if no periods are given instead
		currentTemp     float64
	)

	if len(gridMetaData.Properties.Periods) > 0 {
		if isCelsius {
			currentTemp = gridMetaData.Properties.Periods[0].Temperature.(map[string]interface{})["value"].(float64)
		} else {
			currentTemp = gridMetaData.Properties.Periods[0].Temperature.(float64)
		}
		currentForecast = gridMetaData.Properties.Periods[0].ShortForecast
	}

	return &models.CurrentWeather{
		City:               pointMetaData.Properties.RelativeLocation.Properties.City,
		State:              pointMetaData.Properties.RelativeLocation.Properties.State,
		Forecast:           currentForecast,
		TempCharacteristic: determineTempCharacteristic(isCelsius, currentTemp),
	}, nil
}
