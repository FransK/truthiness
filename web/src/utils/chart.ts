import { ExperimentData, ChartType } from '../types/experiment';

export const determineChartType = (data: ExperimentData, variable: string): ChartType => {
  const values = data.data.map(row => row[variable]);
  return values.every(value => typeof value === 'number') ? 'scatter' : 'bar';
};