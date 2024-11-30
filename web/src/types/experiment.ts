export interface ExperimentData {
    id?: number;
    name: string;
    variables: string[];
    data: Record<string, string | number>[];
    createdAt: Date;
  }
  
  export type ChartType = 'scatter' | 'bar';
  
  export interface ChartConfig {
    xAxis: string;
    yAxis: string;
    chartType: ChartType;
  }