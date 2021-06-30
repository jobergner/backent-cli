export const numericTypes = [
  "int8",
  "uint8",
  "int16",
  "uint16",
  "int32",
  "uint32",
  "int64",
  "uint64",
  "int",
  "uint",
  "uintptr",
  "float32",
  "float64",
  "complex64",
  "complex128",
];

export const textTypes = ["string", "byte", "rune", "[]byte"];

export const defualtValuePerValue = (value) => {
  if (value.startsWith("[]")) {
    return []
  }
  if (numericTypes.includes(value) || value.endsWith("ID")) {
    return 0;
  }
  if (textTypes.includes(value)) {
    return "";
  }
  return false;
};

export const defaultValueAction = (action) => {
  const defaultAciton = {};
  for (const [key, value] of Object.entries(action)) {
    defaultAciton[key] = defualtValuePerValue(value);
  }
  return defaultAciton;
};
