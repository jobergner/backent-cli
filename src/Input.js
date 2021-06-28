import "./Actions.css";
import TextInput from "./TextInput"
import BoolInput from "./BoolInput"
import NumberInput from "./NumberInput"
import SliceInput from "./SliceInput"

const numericTypes = ["int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint", "uintptr", "float32", "float64", "complex64", "complex128"]
const textTypes = ["string", "byte", "rune", "[]byte"]

function Input(props) {
    return (<>
        {Object.entries(props.action).map(([key, value]) => {
            if (value.startsWith("[]")) {
                return <div className="InputField" key={key} >{key}:<SliceInput /></div>
            }
            if (textTypes.includes(value)) {
                return <div className="InputField" key={key} >{key}:<TextInput /></div>
            }
            if (numericTypes.includes(value) || value.endsWith("ID")) {
                return <div className="InputField" key={key}>{key}:<NumberInput /></div>
            }
            if (value === "bool") {
                return <div className="InputField" key={key}>{key}:<BoolInput /></div>
            }
            console.log(value)
            return <div />
        })}
    </>);
}

export default Input;

