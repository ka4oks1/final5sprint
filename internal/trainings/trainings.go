package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parsedInfo := strings.Split(datastring, ",")
	if len(parsedInfo) != 3 {

		return errors.New("wrong length")

	}
	steps, err := strconv.Atoi(parsedInfo[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("incorrect steps count")
	}
	t.Steps = steps

	trainingType := parsedInfo[1]
	t.TrainingType = trainingType

	timeDur, err := time.ParseDuration(parsedInfo[2])
	if err != nil {
		return err
	}
	if timeDur <= 0 {
		return errors.New("incorrect duration time")
	}

	t.Duration = timeDur

	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	meanS := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var spentCal float64
	var err error

	switch {
	case t.TrainingType == "Бег":
		spentCal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case t.TrainingType == "Ходьба":
		spentCal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		spentCal, err = 0, errors.New("unexpected training type")
	}
	if err != nil {
		return "", err
	}

	resultStr := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), dist, meanS, spentCal)

	return resultStr, err
}
