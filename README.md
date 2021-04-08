
## About
Used to validate a user's state configuration. The goal is to validate an input that can be directly casted into golang structs with some additional limitations (thematical validation).
## Valid Example:
```
map[interface{}]interface{}{
	"name": map[interface{}]interface{}{
		"first": "string",
		"last": "string",
		"nickNames": []string,
	},
	"person": map[interface{}]interface{}{
		"name": "name",
		"age": "int",
		"fiends": "[]person",
	},
}
```
## Validation Error Messages
### structural:
| Error | Text | Meaning |
|---|---------|----------|
| ErrIllegalValue | value assigned to key "{KeyName}" in "{ParentObject}" is invalid | An invalid value was defined (nil, "", List, Object in Object). |
<br/> 

### syntactical:
| Error | Text | Meaning |
|---|---------|----------|
| ErrIllegalTypeName | illegal type name "{KeyName}" in "{ParentObject}" | A type was named without adhering to syntax limitations (e.g. "fo$o", "func", "<-+"). |
| ErrInvalidValueString | value "{ValueString}" assigned to "{KeyName}" in "{ParantObject}" is invalid | An invalid value was assigned to a key |
<br/> 

### logical:
| Error | Text | Meaning |
|---|---------|----------|
| ErrTypeNotFound | type with name "{TypeName}" in "{ParentObject}" was not found | A type was referenced as value but not defined anywhere in the data. |
| ErrRecursiveTypeUsage | illegal recursive type detected for "{RecurringKeyNames}" | A recursive type was defined. |
| ErrInvalidMapKey | "{MapKey}" in "{ValueString}" is not a valid map key | An uncomparable type was chosen as map key. |
| ErrUnknownMethod | type "{TypeName}" has no method "{Literal}". | An unknown method was attempted to be referenced. |
<br/> 

### thematical:
Despite the fact that each of these errors would find a place in one of the above mentioned categories, they are listed separately from them since they are specific to the use case, and not related to the validation of actual go declarations at all.
| Error | Text | Meaning |
|---|---------|----------|
| ErrIncompatibleValue | value "{ValueString}" assigned to "{KeyName}" in "{ParentObject}" is incompatible. | The assigned value can't be used, as only golang's basic types, self defined types, and slices of them can be used. |
| ErrNonObjectType | type "{TypeName}" is not an object type. | The defined type is not an object. |
| ErrIllegalCapitalization | {type/field name} "{literal}" starts with a capital letter. | A type or field name starts with a capital letter, which is not allowed. |
| ErrConflictingSingular | "{KeyName1}" and "{KeyName2}" share the same singular form "{Singular}". | Due to the way state will be used two field names cannot have the same singular form. |
| ErrUnavailableFieldName | "{KeyName}" not an available name. | Due to internal usage of this FieldName it is unavailable. |
<br/>


TODO:
- let IDs of types be a valid type (like playerID) and treat them as int