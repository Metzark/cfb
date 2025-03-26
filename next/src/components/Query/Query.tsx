"use client";

import styles from "./Query.module.css";
import { useState } from "react";

interface QueryResult {
  rows: any[];
  rowCount: number;
}

export default function Query() {
  const [query, setQuery] = useState<string>("SELECT school, mascot FROM cfb.teams LIMIT 5;");
  const [results, setResults] = useState<QueryResult | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const executeQuery = async () => {
    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch("/api/query", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ query }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || "Failed to execute query");
      }

      setResults(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred");
    } finally {
      setIsLoading(false);
    }
  };

  const renderTable = () => {
    if (!results?.rows.length) return null;

    const columns = Object.keys(results.rows[0]);

    return (
      <table className={styles.queryTable}>
        <thead>
          <tr>
            {columns.map((column) => (
              <th key={column}>{column}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {results.rows.map((row, rowIndex) => (
            <tr key={rowIndex}>
              {columns.map((column) => (
                <td key={`${rowIndex}-${column}`}>{typeof row[column] === "object" ? JSON.stringify(row[column]) : String(row[column])}</td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    );
  };

  return (
    <div className={styles.query}>
      <ul className={styles.tabList}>
        <button className={styles.active}>SQL</button>
        <button>ChatGPT</button>
      </ul>
      <div className={styles.queryOptionDescription}>
        <p>Run PostgresSQL queries on the CFB database</p>
      </div>
      <div className={styles.textarea}>
        <textarea value={query} onChange={(e) => setQuery(e.target.value)} placeholder="Enter your SQL query here..." />
        <button onClick={executeQuery} disabled={isLoading} className={isLoading ? styles.loading : ""}>
          {isLoading ? "Running..." : "Go"}
        </button>
      </div>
      {error && <div className={styles.queryError}>{error}</div>}
      <div className={styles.queryTableWrap}>{renderTable()}</div>
    </div>
  );
}
