// lib/db.ts
import mysql from 'mysql2/promise';

export const connectToDatabase = async () => {
  const connection = await mysql.createConnection({
    host: 'localhost',
    user: 'your-username',
    password: 'your-password',
    database: 'your-database',
  });
  return connection;
};

export const getUserData = async (username: string) => {
  const connection = await connectToDatabase();
  const [rows] = await connection.execute('SELECT karma, location FROM users WHERE username = ?', [username]);
  await connection.end();
  return rows;
};
