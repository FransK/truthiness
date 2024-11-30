import Dexie, { Table } from 'dexie';
import { ExperimentData } from '../types/experiment';

export class ExperimentDB extends Dexie {
  experiments!: Table<ExperimentData>;

  constructor() {
    super('ExperimentDB');
    this.version(1).stores({
      experiments: '++id, name, createdAt'
    });
  }
}

export const db = new ExperimentDB();