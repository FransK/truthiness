import "./ExperimentViewer.css";

import { useEffect, useState } from "react";
import MyScatter from "../components/ScatterChart/ScatterChart";
import MyDropdown from "../components/SearchableDropdown/SearchableDropdown";

export default function ExperimentViewer() {
  interface GetExperimentsResponse {
    data: {
      Name: string;
      Date: string;
      Location: string;
      Records: string[];
    }[];
  }

  interface IExperiment {
    id: number;
    name: string;
    records: string[];
  }

  type IExperiments = IExperiment[];

  interface GetTrialsResponse {
    data: {
      Data: { Key: string; Value: string }[];
    }[];
  }

  const [experiments, setExperiments] = useState<IExperiments | null>(null);
  const [experiment, setExperiment] = useState<IExperiment | null>(null);
  const [xaxis, setXaxis] = useState<string>("");
  const [yaxis, setYaxis] = useState<string>("");
  const [trials, setTrials] = useState<GetTrialsResponse | null>(null);

  function handleSetExperiment(name: string) {
    setExperiment(
      experiments ? experiments.find((e) => e.name === name) || null : null
    );
  }

  useEffect(() => {
    let ignore = false;
    setExperiments(null);
    fetch("http://localhost:8080/v1/experiments")
      .then((response) => response.json())
      .then((result: GetExperimentsResponse) => {
        if (!ignore) {
          const experiments = result.data.map((e, index) => {
            return {
              id: index,
              name: e.Name,
              records: e.Records,
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

  useEffect(() => {
    if (!experiment) {
      return;
    }

    let ignore = false;
    setTrials(null);
    fetch(`http://localhost:8080/v1/experiments/${experiment.name}/trials`)
      .then((response) => response.json())
      .then((result: GetTrialsResponse) => {
        if (!ignore) {
          console.log(result);
          setTrials(result);
        }
      })
      .catch((error) => console.error("Error fetching trials:", error));
    return () => {
      ignore = true;
    };
  }, [experiment]);

  return (
    <div className="experiment-viewer">
      <div className="experiment">
        <h2>Select an experiment:</h2>
        <MyDropdown
          options={experiments ? experiments.map((e) => e.name) : []}
          id="experimentname"
          selectedVal={experiment ? experiment.name : ""}
          handleChange={(exp: string) => handleSetExperiment(exp)}
        />
      </div>

      {experiment ? (
        <>
          <div className="experiment">
            <h2>Select an X-Axis:</h2>
            <MyDropdown
              options={experiment !== null ? experiment.records : []}
              id="x-axis"
              selectedVal={xaxis}
              handleChange={(record: string) => setXaxis(record)}
            />
          </div>
          <div className="experiment">
            <h2>Select a Y-Axis:</h2>
            <MyDropdown
              options={experiment !== null ? experiment.records : []}
              id="y-axis"
              selectedVal={yaxis}
              handleChange={(record: string) => setYaxis(record)}
            />
          </div>
        </>
      ) : null}

      {experiment && xaxis && yaxis ? (
        <MyScatter
          data={
            trials
              ? trials.data.map((d) => {
                  return d.Data;
                })
              : []
          }
          xkey={xaxis}
          ykey={yaxis}
        />
      ) : null}
    </div>
  );
}
