"use client";

import styles from "./Schema.module.css";
import { useState } from "react";

export default function Schema() {
  const [openTables, setOpenTables] = useState<Record<string, boolean>>({
    "cfb.venues": false,
    "cfb.conferences": false,
    "cfb.teams": false,
    "cfb.games": false,
    "cfb.stats": false,
  });

  const toggleTable = (tableId: string) => {
    setOpenTables((prev) => ({
      ...prev,
      [tableId]: !prev[tableId],
    }));
  };

  return (
    <div className={styles.schema}>
      <h2>CFB Database Schema</h2>
      <div className={styles.table} id="cfb.venues">
        <div className={styles.header} onClick={() => toggleTable("cfb.venues")}>
          <h3>cfb.venues</h3>
          <svg className={openTables["cfb.venues"] ? styles.rotated : ""}>
            <use href="#carrot-up" />
          </svg>
        </div>
        <div className={`${styles.codeBlock} ${!openTables["cfb.venues"] ? styles.hidden : ""}`}>
          <pre>
            <code>{`CREATE TABLE cfb.venues (
  id INTEGER PRIMARY KEY,
  name TEXT,
  city TEXT,
  state TEXT,
  zip TEXT,
  country_code TEXT,
  location_x DOUBLE PRECISION,
  location_y DOUBLE PRECISION,
  elevation DOUBLE PRECISION,
  year_constructed INTEGER,
  capacity INTEGER,
  dome BOOLEAN,
  grass BOOLEAN,
  timezone TEXT
);`}</code>
          </pre>
        </div>
      </div>
      <div className={styles.table} id="cfb.conferences">
        <div className={styles.header} onClick={() => toggleTable("cfb.conferences")}>
          <h3>cfb.conferences</h3>
          <svg className={openTables["cfb.conferences"] ? styles.rotated : ""}>
            <use href="#carrot-up" />
          </svg>
        </div>
        <div className={`${styles.codeBlock} ${!openTables["cfb.conferences"] ? styles.hidden : ""}`}>
          <pre>
            <code>{`CREATE TABLE cfb.conferences (
  id INTEGER PRIMARY KEY,
  name TEXT,
  short_name TEXT,
  abbreviation TEXT
);`}</code>
          </pre>
        </div>
      </div>
      <div className={styles.table} id="cfb.teams">
        <div className={styles.header} onClick={() => toggleTable("cfb.teams")}>
          <h3>cfb.teams</h3>
          <svg className={openTables["cfb.teams"] ? styles.rotated : ""}>
            <use href="#carrot-up" />
          </svg>
        </div>
        <div className={`${styles.codeBlock} ${!openTables["cfb.teams"] ? styles.hidden : ""}`}>
          <pre>
            <code>{`CREATE TABLE cfb.teams (
  id INTEGER PRIMARY KEY,
  school TEXT,
  mascot TEXT,
  abbreviation TEXT,
  alt_name1 TEXT,
  alt_name2 TEXT,
  alt_name3 TEXT,
  classification TEXT,
  color TEXT,
  alt_color TEXT,
  twitter TEXT,
  logo TEXT,
  conference_id INTEGER REFERENCES cfb.conferences(id),
  home_venue_id INTEGER REFERENCES cfb.venues(id)
);`}</code>
          </pre>
        </div>
      </div>
      <div className={styles.table} id="cfb.games">
        <div className={styles.header} onClick={() => toggleTable("cfb.games")}>
          <h3>cfb.games</h3>
          <svg className={openTables["cfb.games"] ? styles.rotated : ""}>
            <use href="#carrot-up" />
          </svg>
        </div>
        <div className={`${styles.codeBlock} ${!openTables["cfb.games"] ? styles.hidden : ""}`}>
          <pre>
            <code>{`CREATE TABLE cfb.games (
  id BIGINT PRIMARY KEY,
  season INTEGER,
  week INTEGER,
  season_type TEXT,
  start_date TIMESTAMPTZ,
  neutral_site BOOLEAN,
  conference_game BOOLEAN,
  home_qtr1_score INTEGER,
  home_qtr2_score INTEGER,
  home_qtr3_score INTEGER,
  home_qtr4_score INTEGER,
  home_final_score INTEGER,
  away_qtr1_score INTEGER,
  away_qtr2_score INTEGER,
  away_qtr3_score INTEGER,
  away_qtr4_score INTEGER,
  away_final_score INTEGER,
  excitement_index DOUBLE PRECISION,
  venue_id INTEGER REFERENCES cfb.venues(id),
  home_id INTEGER REFERENCES cfb.teams(id),
  away_id INTEGER REFERENCES cfb.teams(id),
  winner_id INTEGER REFERENCES cfb.teams(id),
  loser_id INTEGER REFERENCES cfb.teams(id)
);`}</code>
          </pre>
        </div>
      </div>
      <div className={styles.table} id="cfb.stats">
        <div className={styles.header} onClick={() => toggleTable("cfb.stats")}>
          <h3>cfb.stats</h3>
          <svg className={openTables["cfb.stats"] ? styles.rotated : ""}>
            <use href="#carrot-up" />
          </svg>
        </div>
        <div className={`${styles.codeBlock} ${!openTables["cfb.stats"] ? styles.hidden : ""}`}>
          <pre>
            <code>{`CREATE TABLE cfb.stats (
  team_id INTEGER REFERENCES cfb.teams(id),
  season INTEGER,
  interception_yards INTEGER,
  turnovers INTEGER,
  punt_return_yards INTEGER,
  fumbles_lost INTEGER,
  possession_time INTEGER,
  fumbles_recovered INTEGER,
  rushing_tds INTEGER,
  kick_return_yards INTEGER,
  pass_attempts INTEGER,
  interception_tds INTEGER,
  fourth_down_conversions INTEGER,
  punt_returns INTEGER,
  interceptions INTEGER,
  third_down_conversions INTEGER,
  passing_tds INTEGER,
  rushing_yards INTEGER,
  penalty_yards INTEGER,
  kick_return_tds INTEGER,
  fourth_downs INTEGER,
  penalties INTEGER,
  third_downs INTEGER,
  kick_returns INTEGER,
  total_yards INTEGER,
  passes_intercepted INTEGER,
  sacks INTEGER,
  punt_return_tds INTEGER,
  tackles_for_loss INTEGER,
  pass_completions INTEGER,
  first_downs INTEGER,
  rushing_attempts INTEGER,
  net_passing_yards INTEGER,
  -- Offensive stats
  o_plays INTEGER,
  o_drives INTEGER,
  o_ppa DOUBLE PRECISION,
  o_total_ppa DOUBLE PRECISION,
  o_explosiveness DOUBLE PRECISION,
  o_power_success DOUBLE PRECISION,
  o_line_yards DOUBLE PRECISION,
  o_line_yards_total INTEGER,
  o_second_level_yards DOUBLE PRECISION,
  o_second_level_yards_total INTEGER,
  o_open_field_yards DOUBLE PRECISION,
  o_open_field_yards_total INTEGER,
  o_avg_field_position_start REAL,
  o_avg_field_position_predicted_points REAL,
  o_points_per_opportunity DOUBLE PRECISION,
  o_total_opportunities INTEGER,
  -- Defensive stats
  d_plays INTEGER,
  d_drives INTEGER,
  d_ppa DOUBLE PRECISION,
  d_total_ppa DOUBLE PRECISION,
  d_explosiveness DOUBLE PRECISION,
  d_power_success DOUBLE PRECISION,
  d_line_yards DOUBLE PRECISION,
  d_line_yards_total INTEGER,
  d_second_level_yards DOUBLE PRECISION,
  d_second_level_yards_total INTEGER,
  d_open_field_yards DOUBLE PRECISION,
  d_open_field_yards_total INTEGER,
  d_avg_field_position_start REAL,
  d_avg_field_position_predicted_points REAL,
  d_points_per_opportunity DOUBLE PRECISION,
  d_total_opportunities INTEGER,
  PRIMARY KEY (team_id, season)
);`}</code>
          </pre>
        </div>
      </div>
    </div>
  );
}
