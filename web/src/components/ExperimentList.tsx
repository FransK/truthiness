import {
  GetExperimentsResponse,
  IExperiment,
  IExperiments,
} from "../types/experiment";
import { useEffect, useState } from "react";

interface Props {
  onSelect: (experiment: IExperiment) => void;
  selectedId?: number;
}

export function ExperimentList({ onSelect, selectedId }: Props) {
  const [experiments, setExperiments] = useState<IExperiments | null>(null);

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

  if (!experiments) return <div>Loading...</div>;
  if (experiments.length === 0) return <div>No experiments uploaded yet</div>;

  return (
    <div className="space-y-2">
      <h2 className="text-lg font-semibold">Select Experiment</h2>
      <div className="space-y-2">
        {experiments.map((experiment) => (
          <button
            key={experiment.id}
            onClick={() => onSelect(experiment)}
            className={`w-full p-4 text-left rounded-lg border transition-colors
              ${
                selectedId === experiment.id
                  ? "border-blue-500 bg-blue-50"
                  : "border-gray-200 hover:border-blue-300"
              }`}
          >
            <h3 className="font-medium">{experiment.name}</h3>
            <p className="text-sm text-gray-500">
              {experiment.records.length} variables • trials
            </p>
          </button>
        ))}
      </div>
    </div>
  );
}
