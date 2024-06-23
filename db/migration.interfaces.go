package db

import "time"

type User struct {
	ID       uint64  `json:"id" xorm:"'id' notnull pk autoincr"`
	Name     string  `json:"name" xorm:"'name' VARCHAR(100) notnull"`
	Email    string  `json:"email" xorm:"'email' VARCHAR(100) notnull"`
	Finances float64 `json:"finances" xorm:"'finances' notnull"`
	Health   float64 `json:"health" xorm:"'health' notnull"`
	Mental   float64 `json:"mental" xorm:"'mental' notnull"`
	Welfare  float64 `json:"welfare" xorm:"'welfare' notnull"`
	Score    float64 `json:"score" xorm:"'score' notnull"`
	Weight   float64 `json:"weight" xorm:"'weight'notnull"`
	Height   float64 `json:"height" xorm:"'height'notnull"`
	BMI      float64 `json:"bmi" xorm:"'bmi'notnull"`
}

type Finances struct {
	ID        uint64    `json:"id" xorm:"'id' notnull pk autoincr"`
	UserID    uint64    `json:"userid" xorm:"'userid' notnull "`
	Aliment   float64   `json:"aliment"`
	Life      float64   `json:"life"`
	Transport float64   `json:"transport"`
	Shopping  float64   `json:"shopping"`
	Service   float64   `json:"service"`
	Score     float64   `json:"score" xorm:"'score' notnull"`
	Best      float64   `json:"best" xorm:"'best' notnull"`
	Date      time.Time `json:"date" xorm:"'date' notnull"`
}

type Health struct {
	ID             uint64    `json:"id" xorm:"'id' notnull pk autoincr"`
	UserID         uint64    `json:"userid" xorm:"'userid' notnull"`
	CaloriesBurned int       `json:"calories_burned"`
	GoalCalories   int       `json:"goal_calories"`
	HealthyFood    int       `json:"healthy_food"`
	UnhealthyFood  int       `json:"unhealthy_food"`
	Score          float64   `json:"score" xorm:"'score' notnull"`
	Best           uint64    `json:"best" xorm:"'best' notnull"`
	Date           time.Time `json:"date" xorm:"'date' notnull"`
}

type Mental struct {
	ID               uint64    `json:"id" xorm:"'id' notnull pk autoincr"`
	UserID           uint64    `json:"userid" xorm:"'userid' notnull"`
	QualityOfSleep   int       `json:"quality_of_sleep"`
	StressScore      int       `json:"stress_score"`
	DepressionScore  int       `json:"depression_score"`
	MindfulnessScore int       `json:"mindfulness_score"`
	SmokingFrequency int       `json:"smoking_frequency"`
	Score            float64   `json:"score" xorm:"'score' notnull"`
	Best             uint64    `json:"best" xorm:"'best' notnull"`
	Date             time.Time `json:"date" xorm:"'date' notnull"`
}
type Welfare struct {
	ID                  uint64    `json:"id" xorm:"'id' notnull pk autoincr"`
	UserID              uint64    `json:"userid" xorm:"'userid' notnull"`
	Steps               int       `json:"steps"`
	GoalSteps           int       `json:"goal_steps"`
	HoursSlept          float64   `json:"hours_slept"`
	GoalHoursSlept      float64   `json:"goal_hours_slept"`
	MaxSmokingFrequency int       `json:"max_smoking_frequency"`
	SmokingFrequency    int       `json:"smoking_frequency"`
	Score               float64   `json:"score" xorm:"'score' notnull"`
	Best                uint64    `json:"best" xorm:"'best' notnull"`
	Date                time.Time `json:"date" xorm:"'date' notnull"`
}
