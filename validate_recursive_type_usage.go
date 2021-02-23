package validator

func validateRecursiveTypeUsage(yamlData map[interface{}]interface{}) (errs []error) {
	pathBuilder := newPathBuilder(yamlData)

	pathBuilder.build(declarationPath{}, "", yamlData, fieldLevelZero)

	for _, path := range pathBuilder.paths {
		if path.closureKind == pathClosureKindRecursiveness {
			errs = append(errs, newValidationErrorRecursiveTypeUsage(path.joinedNames()))
		}
	}

	return
}
