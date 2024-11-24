package model

type Workoutsheet struct {
	Workouts []Workout
}

type Workout struct {
	Day       int16
	Title     string
	Exercises []Exercise
}

type Exercise struct {
	Title  string
	Series int16
	Reps   int16
}
