import { useEffect, useState } from "react";
import { FileUpload } from "../components/FileUpload";
import { ExperimentList } from "../components/ExperimentList";
import { VisualizationControls } from "../components/VisualizationControls";
import { ExperimentChart } from "../components/ExperimentChart";
import {
  ChartConfig,
  GetExperimentsResponse,
  IExperiment,
  IExperiments,
} from "../types/experiment";
import { LoginForm } from "../components/LoginForm";

interface Props {
  onLogin: (username: string) => void;
  isLoggedIn: boolean;
}

export function ExperimentViewer({ onLogin, isLoggedIn }: Props) {
  const [selectedExperiment, setSelectedExperiment] =
    useState<IExperiment | null>(null);
  const [chartConfig, setChartConfig] = useState<ChartConfig>({
    xAxis: "",
    yAxis: "",
    chartType: "scatter",
  });
  const [experiments, setExperiments] = useState<IExperiments | null>(null);

  function fetchExperiments() {
    let ignore = false;
    fetch(`${import.meta.env.VITE_REST_ADDR}/v1/experiments`)
      .then((response) => response.json())
      .then((result: GetExperimentsResponse) => {
        if (!ignore) {
          const experiments = result.data.map((e, index) => {
            return {
              id: index,
              name: e.Name,
              records: new Map(Object.entries(e.Records)),
            };
          });
          setExperiments(experiments);
        }
      })
      .catch((error) => console.error("Error fetching experiments:", error));
    return () => {
      ignore = true;
    };
  }

  useEffect(() => {
    fetchExperiments();
  }, []);

  const handleUploadSuccess = () => {
    fetchExperiments();
  };

  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="max-w-7xl mx-auto px-4">
        <h1 className="text-3xl font-bold text-gray-900 mb-8">
          Experiment Data Visualizer
        </h1>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div className="space-y-8">
            {isLoggedIn ? (
              <div>
                <h2 className="text-lg font-semibold mb-4">
                  Upload New Experiment
                </h2>
                <FileUpload onUploadSuccess={handleUploadSuccess} />
              </div>
            ) : (
              <div>
                <h2 className="text-lg font-semibold mb-4">
                  Login to upload a new experiment
                </h2>
                <LoginForm onLogin={onLogin} />
              </div>
            )}

            <ExperimentList
              experiments={experiments}
              onSelect={setSelectedExperiment}
              selectedId={selectedExperiment?.id}
            />
          </div>

          <div className="md:col-span-2 space-y-6">
            {selectedExperiment ? (
              <>
                <div className="bg-white p-6 rounded-lg shadow-sm">
                  <h2 className="text-lg font-semibold mb-4">
                    {selectedExperiment.name}
                  </h2>
                  <VisualizationControls
                    experiment={selectedExperiment}
                    config={chartConfig}
                    onConfigChange={setChartConfig}
                  />
                </div>

                <div className="bg-white p-6 rounded-lg shadow-sm">
                  <ExperimentChart
                    experiment={selectedExperiment}
                    config={chartConfig}
                  />
                </div>
              </>
            ) : (
              <div className="bg-white p-6 rounded-lg shadow-sm text-center text-gray-500">
                Select an experiment to visualize
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
