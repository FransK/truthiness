import { IExperiment, ChartConfig } from "../types/experiment";

interface Props {
  experiment: IExperiment;
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

export function VisualizationControls({
  experiment,
  config,
  onConfigChange,
}: Props) {
  const handleAxisChange = (axis: "xAxis" | "yAxis", variable: string) => {
    if (axis === "xAxis" && experiment.records.get(variable) == 2) {
      const chartType = "bar";
      onConfigChange({ ...config, [axis]: variable, chartType });
    } else {
      const chartType = "scatter";
      onConfigChange({ ...config, [axis]: variable, chartType });
    }
  };

  return (
    <div className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700">
          X Axis
        </label>
        <select
          value={config.xAxis}
          onChange={(e) => handleAxisChange("xAxis", e.target.value)}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200"
        >
          <option value="">Select variable</option>
          {[...experiment.records.entries()].map(([key, _], index) => (
            <option key={index} value={key}>
              {key}
            </option>
          ))}
        </select>
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700">
          Y Axis
        </label>
        <select
          value={config.yAxis}
          onChange={(e) => handleAxisChange("yAxis", e.target.value)}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200"
        >
          <option value="">Select variable</option>
          {[...experiment.records.entries()].filter(([_, value]) => value != 2).map(([key, _], index) => (
            <option key={index} value={key}>
              {key}
            </option>
          ))}
        </select>
      </div>
    </div>
  );
}
