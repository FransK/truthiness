import {
  ResponsiveContainer,
  ComposedChart,
  Scatter,
  BarChart,
  Bar,
  LabelList,
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

    if (!experiment || !xAxis) {
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
            setTrialData(result.data.map((t) => t.data));
          }
        }
      })
      .catch((error) => console.error("Error fetching trials:", error));
    return () => {
      ignore = true;
    };
  }, [experiment, xAxis, yAxis, chartType]);
  // End useEffect

  // Create the chart based on selected variables
  let chartContent = <>Select variables to visualize</>;
  if(xAxis && !yAxis) {
    // If no yAxis show a frequency chart for each xAxis (TODO: add bins for continuous variables)
    const buckets = new Map();
    trialData.forEach((t) => {
      for (const [key, value] of Object.entries(t)) {
        if (key === xAxis) {
          if (!buckets.has(value)) {
              buckets.set(value, 1);
            } else {
              buckets.set(value, buckets.get(value) + 1);
            }
          }
        }
    });
    const bucketArray = Array.from(buckets, ([x, count]) => ({x, count}));
    bucketArray.sort((a, b) => {
      if (typeof a.x === "number" && typeof b.x === "number") {
        return a.x - b.x; // Numerical sorting
      } else {
        return String(a.x).localeCompare(String(b.x)); // Lexicographical sorting
      }
    });
    chartContent = (
      <ResponsiveContainer width="100%" height="100%">
        <BarChart data={bucketArray} margin={{ top: 20, right: 20, bottom: 20, left: 20 }}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="x" label={{value: xAxis, position: "bottom"}}/>
          <YAxis label={{value: "Frequency", angle: -90, position: "left"}}/>
          <Tooltip />
          <Bar dataKey="count" fill="#8884d8" isAnimationActive={false} />
        </BarChart>
      </ResponsiveContainer>
    );
  } else if(xAxis && yAxis) {
    if (chartType == ChartType.Scatter) {
      chartContent = (
        <ResponsiveContainer width="100%" height="100%">        
          <ComposedChart data={trialData} margin={{ top: 20, right: 20, bottom: 20, left: 20 }}>
            <CartesianGrid />
            <XAxis type="number" dataKey={xAxis} name={xAxis} label={{value: xAxis, position: "bottom"}}/>
            <YAxis type="number" name={yAxis} label={{value: yAxis, angle: -90, position: "left"}}/>
            <Tooltip cursor={{ strokeDasharray: "3 3" }} />
            <Scatter dataKey={yAxis} fill="#8884d8" isAnimationActive={false} />
            <Line dataKey="LineY" stroke="#8884d8" dot={false} activeDot={false} legendType="none" isAnimationActive={false} />
          </ComposedChart>
        </ResponsiveContainer>
      );
    } else if (chartType == ChartType.Bar) {
      // If showing a bar chart, show the mean of each group with respect to yAxis
      const means = calculateAverages(trialData, xAxis, yAxis);
      chartContent = (
        <ResponsiveContainer width="100%" height="100%">    
          <BarChart data={means} margin={{ top: 20, right: 20, bottom: 20, left: 20 }} >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="key" label={{value: xAxis, position: "bottom"}}/>
            <YAxis label={{value: "Mean " + yAxis, angle: -90, position: "left"}}/>
            <Tooltip />
            <Bar dataKey="mean" fill="#8884d8" isAnimationActive={false}>
              <LabelList dataKey="count" position="top" />
            </Bar>
          </BarChart>
        </ResponsiveContainer>
      );
    }
  }

  return (
    <div className="h-[400px] flex items-center justify-center text-gray-500">
      {chartContent}
    </div>
  );
}

// calculateAverages: takes an array of trials and the names of an x and y axis
// returns an array of the average y value by grouping each x value
function calculateAverages(trialData: Array<Record<string, any>>, xAxis: string, yAxis: string): Array<Record<string, any>> {
  const means = new Map<string | number, { sum: number; count: number }>();

  // Iterate through the trialData array
  trialData.forEach((t) => {
    if (xAxis in t && yAxis in t) {
      const xValue = t[xAxis];
      const yValue = t[yAxis];
      if (typeof yValue === "number") {
        if (!means.has(xValue)) {
          means.set(xValue, { sum: 0, count: 0 });
        }
        const current = means.get(xValue)!;
        current.sum += yValue;
        current.count += 1;
      }
    }
  });

  // Calculate averages and return as an array of tuples
  const result = Array.from(means, ([key, value]) => {
    return {
      key: key,
      mean: value.sum / value.count,
      count: value.count,
    }
  });

  result.sort((a, b) => {
    if (typeof a.key === "number" && typeof b.key === "number") {
      return a.key - b.key; // Numerical sorting
    } else {
      return String(a.key).localeCompare(String(b.key)); // Lexicographical sorting
    }
  });

  return result;
}
