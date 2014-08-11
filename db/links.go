package db

func LinkOne(kind Kind, model Model) error {
	return PopulateById(kind, model)
}

// This doesn't work because go does not automatically
// convert a slice of type to its interface
func LinkMany(kind Kind, models *[]Model) []error {
	var errors []error

	for _, model := range *models {
		if err := PopulateById(kind, model); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
