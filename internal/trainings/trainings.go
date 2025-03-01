package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go1fl-sprint5-final/internal/personaldata"
	"go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {

	arr := strings.Split(datastring, ",")

	if len(arr) != 3 {
		return fmt.Errorf("the length of the slice is not equal to three")
	}

	steps, err := strconv.Atoi(arr[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("incorrect number of steps")
	}

	t.Steps = steps

	switch arr[1] {
	case "Бег":
		t.TrainingType = arr[1]
	case "Ходьба":
		t.TrainingType = arr[1]
	default:
		return fmt.Errorf("wrong type of training")
	}

	duration, err := time.ParseDuration(arr[2])

	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if duration < 0 {
		return fmt.Errorf("incorrect duration")
	}

	t.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() string {

	dist := spentenergy.Distance(t.Steps)
	avrSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	typeOfTraining := t.TrainingType
	var spentCalories float64

	switch typeOfTraining {
	case "Бег":
		spentCalories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	case "Ходьба":
		spentCalories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "неизвестный тип тренировки"
	}

	result := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		typeOfTraining, t.Duration.Hours(), dist, avrSpeed, spentCalories)

	return result
}
