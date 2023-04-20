package typescript

func Const(name string) *Code {
	return NewCode().Const(name)
}

func Let(name string) *Code {
	return NewCode().Let(name)
}

func Function(name string) *Code {
	return NewCode().Function(name)
}

func If(code *Code) *Code {
	return NewCode().If(code)
}

func ForIn(decl, iterable *Code) *Code {
	return NewCode().ForIn(decl, iterable)
}

func Interface(name string, fields ...InterfaceField) *Code {
	return NewCode().Interface(name, fields...)
}

func Id(s string) *Code {
	return NewCode().Id(s)
}

func Return() *Code {
	return NewCode().Return()
}

func Object(fields ...ObjectField) *Code {
	return NewCode().Object(fields...)
}

func ObjectSpaced(fields ...ObjectField) *Code {
	return NewCode().ObjectSpaced(fields...)
}

func Index(code *Code) *Code {
	return NewCode().Index(code)
}

func Export() *Code {
	return NewCode().Export()
}

func Lit(s string) string {
	return NewCode().Lit(s)
}

func Delete() *Code {
	return NewCode().Delete()
}

func CodeSet(code ...*Code) *Code {
	return NewCode().CodeSet(code...)
}

func Empty() *Code {
	return NewCode()
}

func Private() *Code {
	return NewCode().Private()
}

func Public() *Code {
	return NewCode().Public()
}

func ArrowFunc(code ...*Code) *Code {
	return NewCode().ArrowFunc(code...)
}

func Switch() *Code {
	return NewCode().Switch()
}

func Case(matcher *Code, code ...*Code) *Code {
	return NewCode().Case(matcher, code...)
}

func Default(code ...*Code) *Code {
	return NewCode().Default(code...)
}

func Promise(typeName string) *Code {
	return NewCode().Promise(typeName)
}
