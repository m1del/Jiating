import React from 'react';
import { styles } from '../../../styles';

type ButtonProps = {
  buttonText: string;
  onClick?: () => void;
  type?: "button" | "submit";
  additionalClasses?: string;
};

const Button: React.FC<ButtonProps> = ({
  buttonText,
  onClick = () => {},
  type = "button",
  additionalClasses = ""
}) => {
  return (
    <button type={type} className={`${styles.button} ${additionalClasses}`} onClick={onClick}>
      {buttonText}
    </button>
  );
};

export default Button;
