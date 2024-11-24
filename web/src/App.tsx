
import './App.css'
import MyScatter from './components/ui/ScatterChart'
import MyDropdown from './components/ui/SearchableDropdown'
import { animals } from "./testdata/animals";
import { data01, data02 } from "./testdata/chartdata";
import { useState } from "react";

function App() {
  const [value, setValue] = useState("Select option...");

  return (
    <>
      <MyScatter 
        data01={data01}
        data02={data02}
      />
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
