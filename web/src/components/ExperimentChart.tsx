import {
  ResponsiveContainer,
  ComposedChart,
  Scatter,
  BarChart,
  Bar,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
} from "recharts";
import {
  ChartConfig,
  ChartType,
  GetTrialsResponse,
  IExperiment,
} from "../types/experiment";
import { useEffect, useState } from "react";

interface Props {
  experiment: IExperiment;
  config: ChartConfig;
}

export function ExperimentChart({ experiment, config }: Props) {
  const [trialData, setTrialData] = useState<{Key: string, Value: any}[][]>([])
  const { xAxis, yAxis, chartType } = config;

  useEffect(() => {
    setTrialData([]);

    if (!experiment) {
      return;
    }

    if (!xAxis || !yAxis) {
      return;
    }

    let ignore = false;
    let fetchString = chartType == ChartType.Scatter ? 
      `${import.meta.env.VITE_REST_ADDR}/v1/experiments/${
        experiment.name
      }/trials?x_axis=${xAxis}&y_axis=${yAxis}&model=linear` :
      `${import.meta.env.VITE_REST_ADDR}/v1/experiments/${
        experiment.name
      }/trials?x_axis=${xAxis}&y_axis=${yAxis}`
    fetch(fetchString)
      .then((response) => {
        if (!response.ok) {
          throw Error(response.statusText)
        }
        return response.json();
      })
      .then((result: GetTrialsResponse) => {
        if (!ignore) {          
          if (result) {            
            // After modifying the `Data` arrays, use the updated `trials.data`
            setTrialData(result.data.map((t) => t.Data));
          }
        }
      })
      .catch((error) => console.error("Error fetching trials:", error));
    return () => {
      ignore = true;
    };
  }, [experiment, xAxis, yAxis, chartType]);

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
        {chartType == ChartType.Scatter ? (
          <ComposedChart data={trialData} margin={{ top: 20, right: 20, bottom: 20, left: 20 }}>
            <CartesianGrid />
            <XAxis type="number" dataKey={xAxis} name={xAxis} />
            <YAxis type="number" name={yAxis} />
            <Tooltip cursor={{ strokeDasharray: "3 3" }} />
            <Scatter dataKey={yAxis} fill="#8884d8" isAnimationActive={false} />
            <Line dataKey="LineY" stroke="#8884d8" dot={false} activeDot={false} legendType="none" isAnimationActive={false} />
          </ComposedChart>
        ) : (
          <BarChart
            data={ trialData }
            margin={{ top: 20, right: 20, bottom: 20, left: 20 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey={xAxis} />
            <YAxis />
            <Tooltip />
            <Bar dataKey={yAxis} fill="#8884d8" isAnimationActive={false} />
          </BarChart>
        )}
      </ResponsiveContainer>
    </div>
  );
}
