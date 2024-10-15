import mysql, { QueryResult } from "mysql2/promise";

export const getConnection = () => {
  const pool = mysql.createPool({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE,
    port: Number(process.env.DB_PORT),
    namedPlaceholders: true,
  });

  return pool;
};

/**
 * SQL実行
 * @param sql
 * @param data
 * @returns
 */
export async function execute(sql: string, data?: any): Promise<QueryResult> {
  const pool = await getConnection();
  const [results] = await pool.query(sql, data);
  await pool.end();

  return results;
};

export async function executeInsert(sql: string, data?: any): Promise<string> {
  const pool = await getConnection();
  const [results] = await pool.query(sql, data);
  const id = (results as any).insertId;
  await pool.end();

  return id as string;
};