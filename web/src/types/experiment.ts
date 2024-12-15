export interface ExperimentData {
    id?: number;
    name: string;
    variables: string[];
    data: Record<string, string | number>[];
    createdAt: Date;
  }
  
  export const ChartType = {
    Scatter: 1,
    Bar: 2,
  }
  
  export interface ChartConfig {
    xAxis: string;
    yAxis: string;
    chartType: number;
  }

  export interface GetExperimentsResponse {
    data: {
      Name: string;
      Date: string;
      Location: string;
      Records: Map<string, number>;
    }[];
  }

  export interface IExperiment {
    id: number;
    name: string;
    records: Map<string, number>;
  }

  export type IExperiments = IExperiment[];

  export interface GetTrialsResponse {
    data: {
      Data: { Key: string; Value: any }[];
    }[];
  }