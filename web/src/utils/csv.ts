import Papa from 'papaparse';
import { ExperimentData } from '../types/experiment';

export const parseCSV = (file: File): Promise<ExperimentData> => {
  return new Promise((resolve, reject) => {
    Papa.parse(file, {
      header: true,
      dynamicTyping: true,
      complete: (results) => {
        const variables = Object.keys(results.data[0]);
        resolve({
          name: file.name.replace('.csv', ''),
          variables,
          data: results.data as Record<string, string | number>[],
          createdAt: new Date()
        });
      },
      error: (error) => {
        reject(error);
      }
    });
  });
};