import {
  ResponsiveContainer,
  ScatterChart,
  Scatter,
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
} from "recharts";
import {
  ChartConfig,
  GetTrialsResponse,
  IExperiment,
} from "../types/experiment";
import { useEffect, useState } from "react";

interface Props {
  experiment: IExperiment;
  config: ChartConfig;
}

export function ExperimentChart({ experiment, config }: Props) {
  const [trials, setTrials] = useState<GetTrialsResponse | null>(null);
  const { xAxis, yAxis, chartType } = config;

  useEffect(() => {
    if (!experiment) {
      return;
    }

    let ignore = false;
    setTrials(null);
    fetch(
      `${import.meta.env.VITE_REST_ADDR}/v1/experiments/${
        experiment.name
      }/trials`
    )
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

  if (!xAxis || !yAxis) {
    return (
      <div className="h-[400px] flex items-center justify-center text-gray-500">
        Select variables to visualize
      </div>
    );
  }

  return (
    <div className="h-[400px] w-full">
      <ResponsiveContainer width="100%" height="100%">
        {chartType === "scatter" ? (
          <ScatterChart margin={{ top: 20, right: 20, bottom: 20, left: 20 }}>
            <CartesianGrid />
            <XAxis type="number" dataKey={xAxis} name={xAxis} />
            <YAxis type="number" dataKey={yAxis} name={yAxis} />
            <Tooltip cursor={{ strokeDasharray: "3 3" }} />
            <Scatter
              data={
                trials
                  ? trials.data.map((d) => {
                      return d.Data;
                    })
                  : []
              }
              fill="#8884d8"
            />
          </ScatterChart>
        ) : (
          <BarChart
            data={
              trials
                ? trials.data.map((d) => {
                    return d.Data;
                  })
                : []
            }
            margin={{ top: 20, right: 20, bottom: 20, left: 20 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey={xAxis} />
            <YAxis />
            <Tooltip />
            <Bar dataKey={yAxis} fill="#8884d8" />
          </BarChart>
        )}
      </ResponsiveContainer>
    </div>
  );
}
