"use client";

import styles from "./Header.module.css";
import { usePathname } from "next/navigation";

export default function Header() {
  const path = usePathname();

  return (
    <nav className={styles.nav}>
      <span className={styles.logo}>CFB</span>
      <ul>
        <a className={`${styles.link} ${path === "/" ? styles.active : ""}`} href="/">
          Home
        </a>
        <a className={`${styles.link} ${path === "/teams" ? styles.active : ""}`} href="/teams">
          Teams
        </a>
        <a className={`${styles.link} ${path === "/predict" ? styles.active : ""}`} href="/predict">
          Predict
        </a>
        <a className={`${styles.link} ${path === "/query" ? styles.active : ""}`} href="/query">
          Query
        </a>
      </ul>
      <a className={styles.login}>Login</a>
    </nav>
  );
}
