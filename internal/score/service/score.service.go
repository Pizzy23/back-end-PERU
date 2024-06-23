package score

import "peru/db"

func CalculateWelfareScore(data db.Welfare) float64 {
	activityScore := (float64(data.Steps) / float64(data.GoalSteps) * 100)
	sleepScore := (data.HoursSlept / data.GoalHoursSlept * 100)
	riskScore := (1 - (float64(data.SmokingFrequency) / float64(data.MaxSmokingFrequency))) * 100

	return (activityScore + sleepScore + riskScore) / 3
}

func CalculateMentalScore(data db.Mental) float64 {
	return (float64(data.StressScore) + float64(data.DepressionScore) + float64(data.MindfulnessScore)) / 3
}

func CalculateHealthScore(data db.Health) float64 {
	nutritionScore := (float64(data.HealthyFood) / float64(data.UnhealthyFood)) * 100
	return (float64(data.CaloriesBurned)/float64(data.GoalCalories)*100 + nutritionScore) / 2
}

func CalculateFinanceScore(data db.Finances) float64 {
	w1, w2, w3, w4, w5 := 1.0, 1.0, 1.0, 1.0, 1.0
	totalScore := (w1*data.Aliment + w2*data.Life + w3*data.Transport + w4*data.Shopping + w5*data.Service) / (w1 + w2 + w3 + w4 + w5)
	return totalScore
}

func CalculateUserScore(finances float64, health float64, mental float64, welfare float64) float64 {
	return (finances + health + mental + welfare) / 4
}

func CalculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func BMICategory(bmi float64) string {
	if bmi < 18.5 {
		return "Underweight"
	} else if bmi >= 18.5 && bmi < 24.9 {
		return "Normal weight"
	} else if bmi >= 25 && bmi < 29.9 {
		return "Overweight"
	}
	return "Obesity"
}
