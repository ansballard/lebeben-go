import { value } from "./new.jsx";
import { red } from "../constants/colors.jsx";

const testing = 123;

console.log(`ok then, ${red}`);

const func = async () => {
  const we = await Promise.resolve({testing, value});
  console.log(we);
}

func();

console.log("whatever5")
