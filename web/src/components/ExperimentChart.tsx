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
import { ExperimentData, ChartConfig } from "../types/experiment";

interface Props {
  experiment: ExperimentData;
  config: ChartConfig;
}

export function ExperimentChart({ experiment, config }: Props) {
  const { xAxis, yAxis, chartType } = config;

  if (!xAxis || !yAxis) {
    return (
      <div className="h-[400px] flex items-center justify-center text-gray-500">
        Select variables to visualize
      </div>
    );
  }

  const data = experiment.data.map((row) => ({
    x: row[xAxis],
    y: row[yAxis],
  }));

  return (
    <div className="h-[400px] w-full">
      <ResponsiveContainer width="100%" height="100%">
        {chartType === "scatter" ? (
          <ScatterChart margin={{ top: 20, right: 20, bottom: 20, left: 20 }}>
            <CartesianGrid />
            <XAxis type="number" dataKey="x" name={xAxis} />
            <YAxis type="number" dataKey="y" name={yAxis} />
            <Tooltip cursor={{ strokeDasharray: "3 3" }} />
            <Scatter data={data} fill="#8884d8" />
          </ScatterChart>
        ) : (
          <BarChart
            data={data}
            margin={{ top: 20, right: 20, bottom: 20, left: 20 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="x" />
            <YAxis />
            <Tooltip />
            <Bar dataKey="y" fill="#8884d8" />
          </BarChart>
        )}
      </ResponsiveContainer>
    </div>
  );
}
