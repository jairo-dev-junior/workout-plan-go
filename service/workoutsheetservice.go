package service

import (
	"github.com/jairo-dev-junior/workout-plan-go/client"
	"github.com/jairo-dev-junior/workout-plan-go/model"
	"github.com/jairo-dev-junior/workout-plan-go/model/client"
)

func createWorkout(bodyPart string, day int16) model.Workout {
	exercisesReq := client.GetExercise(bodyPart)
	exercises := []model.Exercise{}
	var workout = model.Workout{}

	for _, exerciseReq := range exercisesReq {
		var data = model.Exercise{
			Title:  exerciseReq.name,
			Series: 4,
			Reps:   8,
		}
		append(exercises, data)
	}

	workout.Day = day
	workout.Exercises = exercises
	workout.Title = bodyPart + " workout"

	return workout
}

func createWorkoutSheet(bodyParts []string, days []int16) model.Workoutsheet {
	workoutsheet := model.Workoutsheet{}
	for index, day := range days {
		append(workoutsheet.Workouts, createWorkout(bodyParts[index], day))
	}

	return workoutsheet
}

func Main(bodyParts []string, days []int16) model.Workoutsheet {
	return createWorkoutSheet(bodyParts, days)
}
