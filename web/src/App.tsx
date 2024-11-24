
import './App.css'
import MyScatter from './components/ui/ScatterChart'
import MyDropdown from './components/ui/SearchableDropdown'
import { animals } from "./testdata/animals";
import { useState } from "react";

function App() {
  const [value, setValue] = useState("Select option...");

  return (
    <>
      <MyScatter />
      <MyDropdown
        options={animals}
        label="name"
        id="id"
        selectedVal={value}
        handleChange={(val: string) => setValue(val)}
      />
    </>
  )
}

export default App
