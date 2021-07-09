import "./Actions.css";
import evalInput from "./evalInput"

function Input(props) {
  const { setFormContent, currentFormContent, action } = props;
  return (
    <>
      {Object.entries(action).map(([key, value]) => {
        return evalInput(key, value, setFormContent, currentFormContent)
      })}
    </>
  );
}

export default Input;
