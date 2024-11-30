import { useLiveQuery } from "dexie-react-hooks";
import { db } from "../db/db";
import { ExperimentData } from "../types/experiment";

interface Props {
  onSelect: (experiment: ExperimentData) => void;
  selectedId?: number;
}

export function ExperimentList({ onSelect, selectedId }: Props) {
  const experiments = useLiveQuery(() =>
    db.experiments.orderBy("createdAt").reverse().toArray()
  );

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
              {experiment.variables.length} variables • {experiment.data.length}{" "}
              trials
            </p>
          </button>
        ))}
      </div>
    </div>
  );
}
