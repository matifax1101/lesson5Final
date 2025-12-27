package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		dist,
		cal,
	), nil
}
