package model

import (
	"go-prompt/model/enum"

	"github.com/google/uuid"
)

type Application struct {
	Id                         uuid.UUID
	CourseIterationId          uuid.UUID
	StudyDegree                enum.StudyDegree
	StudyProgram               string
	CurrentSemester            int
	GermanLanguageProficiency  enum.LanguageProficiency
	EnglishLanguageProficiency enum.LanguageProficiency
	Devices                    []enum.Device
	CoursesTaken               []enum.Course
	Experience                 string
	Motivation                 string
	AssessmentId               uuid.UUID
	FinalGradeId               uuid.UUID
}
