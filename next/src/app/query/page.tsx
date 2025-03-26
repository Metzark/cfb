import styles from "./page.module.css";
import { Metadata } from "next";
import Query from "@/components/Query/Query";
import Schema from "@/components/Schema/Schema";

export const metadata: Metadata = {
  title: "CFB - Query",
};

export default function Page() {
  return (
    <main className={styles.main}>
      <Query />
      <Schema />
    </main>
  );
}
