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

func (ds *DaySteps) Parse(datastring string) (err error) {
	parsedInfo := strings.Split(datastring, ",")
	if len(parsedInfo) != 2 {

		return errors.New("wrong length")

	}
	steps, err := strconv.Atoi(parsedInfo[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("incorrect steps count")
	}
	ds.Steps = steps

	timeDur, err := time.ParseDuration(parsedInfo[1])
	if err != nil {
		return err
	}
	if timeDur <= 0 {
		return errors.New("incorrect duration time")
	}

	ds.Duration = timeDur

	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {

	dist := spentenergy.Distance(ds.Steps, ds.Height)

	var spentCal float64
	var err error

	spentCal, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	resultStr := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, dist, spentCal)
	return resultStr, err
}
