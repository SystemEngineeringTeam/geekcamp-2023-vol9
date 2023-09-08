import { useRouter } from "next/router";
import styles from "./header.module.scss";
import { useRecoilValue } from "recoil";
import { headerShadowState } from "@/components/recoil/state";

export default function Header() {
  const shadow = useRecoilValue(headerShadowState);
  const router = useRouter();

  return (
    <header
      className={styles.header}
      data-shadow={shadow}
      onClick={() => router.push("/")}
    >
      <h1 className={styles.title}>CampusCrowdMonitor</h1>
    </header>
  );
}
