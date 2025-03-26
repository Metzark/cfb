import { NextResponse } from "next/server";
import { Pool } from "pg";

// Create a connection pool using the read-only user
const pool = new Pool({
  user: "readonly",
  password: process.env.CFB_DB_READONLY_PASSWORD,
  host: process.env.CFB_DB_HOST,
  port: parseInt(process.env.CFB_DB_PORT || "5432"),
  database: process.env.CFB_DB_NAME,
  ssl: process.env.CFB_DB_SSLMODE === "require" ? { rejectUnauthorized: false } : false,
});

export async function POST(req: Request) {
  try {
    const { query } = await req.json();

    // Basic SQL injection prevention - only allow SELECT statements
    if (!query.trim().toLowerCase().startsWith("select")) {
      return NextResponse.json({ error: "Only SELECT queries are allowed" }, { status: 400 });
    }

    // Execute the query
    const result = await pool.query(query);

    return NextResponse.json({
      rows: result.rows,
      rowCount: result.rowCount,
    });
  } catch (error) {
    console.error("Database query error:", error);
    return NextResponse.json({ error: "Failed to execute query" }, { status: 500 });
  }
}
