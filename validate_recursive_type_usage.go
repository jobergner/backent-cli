package validator

func validateRecursiveTypeUsage(data map[interface{}]interface{}) (errs []error) {
	pathBuilder := newPathBuilder(data)

	pathBuilder.build(declarationPath{}, "", data, fieldLevelZero)

	for _, path := range pathBuilder.paths {
		if path.closureKind == pathClosureKindRecursiveness {
			errs = append(errs, newValidationErrorRecursiveTypeUsage(path.joinedNames()))
		}
	}

	return
}
