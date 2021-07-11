package main

import (
	validator "github.com/jobergner/backent-cli/validator"
)

func validateConfig(c *config) []error {
	if errs := validator.ValidateStateConfig(c.State); len(errs) != 0 {
		return errs
	}
	if errs := validator.ValidateActionsConfig(c.State, c.Actions); len(errs) != 0 {
		return errs
	}
	if errs := validator.ValidateResponsesConfig(c.State, c.Actions, c.Responses); len(errs) != 0 {
		return errs
	}
	return nil
}
