
## About
Used to validate a user's state configuration.
<br/>
## Validation Error Messages
<br/> 

### TODO:
- ErrUnknownMethod
- ErrIncompatibleValue
- ErrNonObjectType

### structural:
| Error | Text | Meaning |
|---|---------|----------|
| ErrIllegalValue | value assigned to key "{KeyName}" in "{ParentObject}" is invalid | An invalid value was defined (nil, "", List, Object in Object). |
<br/> 

### syntactical:
| Error | Text | Meaning |
|---|---------|----------|
| ErrIllegalTypeName | illegal type name "{KeyName}" in "{ParentObject}" | A type was named without adhering to go's syntax limitations (e.g. "fo$o", "func", "<-+"). |
| ErrInvalidValueString | value "{ValueString}" assigned to "{KeyName}" in "{ParantObject}" is invalid | An invalid value was assigned to a key |
<br/> 

### logical:
| Error | Text | Meaning |
|---|---------|----------|
| ErrTypeNotFound | type with name "{TypeName}" in "{ParentObject}" was not found | A type was referenced as value but not defined anywhere in the YAML document. |
| ErrRecursiveTypeUsage | illegal recursive type detected for "{RecurringKeyNames}" | A recursive type was defined. |
| ErrInvalidMapKey | "{MapKey}" in "{ValueString}" is not a valid map key | An uncomparable type was chosen as map key. |
| ErrUnknownMethod | type "{TypeName}" has no method "{Literal}". | An unknown method was attempted to be referenced. |
<br/> 

### thematical:
| Error | Text | Meaning |
|---|---------|----------|
| ErrIncompatibleValue | value "{TypeName}" assigned to "{KeyName}" in "{ParentObject}" is incompatible. | The assigned value can't be used, as only golang's basic types, self defined types, and slices of them can be used. |
| ErrNonObjectType | type "{TypeName}" is not an object type. | The defined type can't be converted to a go struct as it's not an object. |
<br/>


