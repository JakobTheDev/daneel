package program

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
)

func ListPrograms() (programs []models.Program, err error) {
	programs, err = GetProgramListFromDb()
	if err != nil {
		return nil, fmt.Errorf("failed to get program list: %w", err)
	}
	return programs, nil
}

func AddProgram(program models.Program) (err error) {
	err = AddProgramToDb(program)
	if err != nil {
		return fmt.Errorf("failed to add program: %w", err)
	}
	return nil
}

func RemoveProgram(program models.Program) (err error) {
	err = RemoveProgramFromDb(program)
	if err != nil {
		return fmt.Errorf("failed to remove program: %w", err)
	}
	return nil
}
