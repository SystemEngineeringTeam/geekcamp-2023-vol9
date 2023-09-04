import { useEffect, useState } from "react";
import { StarFilledIcon, StarOutlineIcon } from "../util/icons";
import styles from "./room.module.scss";
import { useLocalStorage } from "@mantine/hooks";

type Props = {
  id: string;
  name: string;
  staycount: number;
  building: string;
  area: string;
  isSelect: boolean;
};

export default function Room(props: Props) {
  const [isStared, setIsStared] = useState(false);
  const [stars, setStars] = useLocalStorage<string[]>({
    key: "stars",
    defaultValue: [],
  });

  useEffect(() => {
    if (stars.includes(props.id)) setIsStared(true);
    else setIsStared(false);
  }, [props.id, stars]);

  useEffect(() => {
    if (isStared) setStars((s) => [...s, props.id]);
    else setStars((s) => s.filter((star) => star !== props.id));
  }, [isStared, props.id, setStars]);

  return (
    <section
      id={props.id}
      className={styles.room}
      data-select={props.isSelect}
      data-star={isStared}
    >
      <h1 className={styles.title}>
        <span className={styles.room_name}>{props.name}</span>
        <span className={styles.building_name}>({props.building})</span>
      </h1>

      <p className={styles.staycount}>
        <span>{"滞在者数: "}</span>
        <span className={styles.count}>{props.staycount}</span>
        <span>{"人"}</span>
      </p>

      {isStared ? (
        <StarFilledIcon
          className={styles.star}
          onClick={() => setIsStared(false)}
        />
      ) : (
        <StarOutlineIcon
          className={styles.star}
          onClick={() => setIsStared(true)}
        />
      )}
    </section>
  );
}
