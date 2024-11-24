import "./ExperimentViewer.css";

import { useState } from "react";
import MyScatter from "../components/ui/ScatterChart/ScatterChart";
import MyDropdown from "../components/ui/SearchableDropdown/SearchableDropdown";
import { animals } from "../testdata/animals";
import { data01, data02 } from "../testdata/chartdata";

export default function ExperimentViewer() {
  const [experiment, setExperiment] = useState("");

  return (
    <div className="experiment-viewer">
      <div className="experiment">
        <h2>Select an experiment:</h2>
        <MyDropdown
          options={animals}
          label="name"
          id="id"
          selectedVal={experiment}
          handleChange={(exp: string) => setExperiment(exp)}
        />
      </div>
      <MyScatter data01={data01} data02={data02} />
    </div>
  );
}
