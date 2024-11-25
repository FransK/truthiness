import "./ExperimentViewer.css";

import { useEffect, useState } from "react";
import MyScatter from "../components/ui/ScatterChart/ScatterChart";
import MyDropdown from "../components/ui/SearchableDropdown/SearchableDropdown";
import { data01, data02 } from "../testdata/chartdata";

export default function ExperimentViewer() {
  interface ApiResponse {
    data: {
      Name: string;
      Date: string;
      Location: string;
    }[];
  }

  interface IExperiment {
    id: number;
    name: string;
  }

  type IExperiments = IExperiment[];

  const [experiment, setExperiment] = useState("");
  const [experiments, setExperiments] = useState<IExperiments | null>(null);

  useEffect(() => {
    let ignore = false;
    setExperiments(null);
    fetch("http://localhost:8080/v1/experiments")
      .then((response) => response.json())
      .then((result: ApiResponse) => {
        if (!ignore) {
          console.log(result);
          const experiments = result.data.map((e, index) => {
            return {
              id: index,
              name: e.Name,
            };
          });
          setExperiments(experiments);
        }
      })
      .catch((error) => console.error("Error fetching experiments:", error));
    return () => {
      ignore = true;
    };
  }, []);

  return (
    <div className="experiment-viewer">
      <div className="experiment">
        <h2>Select an experiment:</h2>
        <MyDropdown
          options={experiments ? experiments.map((e) => e.name) : []}
          id="id"
          selectedVal={experiment}
          handleChange={(exp: string) => setExperiment(exp)}
        />
      </div>
      <MyScatter data01={data01} data02={data02} />
    </div>
  );
}
