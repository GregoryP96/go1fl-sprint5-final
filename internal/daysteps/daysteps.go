package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go1fl-sprint5-final/internal/personaldata"
	"go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {

	arr := strings.Split(datastring, ",")

	if len(arr) != 2 {
		return fmt.Errorf("the length of the slice is not equal to two")
	}

	steps, err := strconv.Atoi(arr[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("incorrect number of steps")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(arr[1])

	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if duration < 0 {
		return fmt.Errorf("incorrect duration")
	}

	ds.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() string {
	dist := spentenergy.Distance(ds.Steps)
	spentCalories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	result := fmt.Sprintf(`Количество шагов: %d.
						   Дистанция составила %.2f км.
						   Вы сожгли %.2f ккал.`, ds.Steps, dist, spentCalories)
	return result
}
