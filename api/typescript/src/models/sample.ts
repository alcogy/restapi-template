import { execute, executeInsert } from "./database";

export interface Sample {
  id: string;
  name: string;
  value: string;
  latest_update: string;
}

export const selectAllSamples = async (): Promise<Sample[]> => {
  const sql: string = `select * from my_table;`;
  const results = await execute(sql);

  return results as Sample[];
}

export const selectSampleByID = async (id: string): Promise<Sample> => {
  const sql: string = `select * from my_table where id = :id;`;
  const params = { id : id };
  const results = await execute(sql, params);

  return (results as Sample[])[0];
}

export const insertSample = async (value: Sample): Promise<string> => {
  const sql: string = `insert into my_table values(0, :name, :value, now());`;
  const id = await executeInsert(sql, value);
  
  return id;
}

export const updateSample = async (value: Sample) => {
  const sql: string = `update my_table set name = :name, value = :value, latest_update = now() where id = :id;`;
  await execute(sql, value);
}

export const deleteSampleData = async (id: string) => {
  const sql: string = `delete from my_table where id = :id;`;
  const params = { id : id };
  await execute(sql, params);
}