import { useEffect, useState } from "react";
import { StarFilledIcon, StarOutlineIcon } from "../util/icons";
import styles from "./room.module.scss";
import { useLocalStorage } from "@mantine/hooks";

type Props = {
  key: string;
  id: number;
  name: string;
  building: string;
  staycount: number | null;
  isSelect: boolean;
};

export default function Room(props: Props) {
  const [isStared, setIsStared] = useState(false);
  const [stars, setStars] = useLocalStorage<string[]>({
    key: "stars",
    defaultValue: [],
  });
  const roomId = props.id.toString();

  useEffect(() => {
    if (stars.includes(roomId)) setIsStared(true);
    else setIsStared(false);
  }, [roomId, stars]);

  useEffect(() => {
    if (isStared) setStars((s) => [...s, roomId]);
    else setStars((s) => s.filter((star) => star !== roomId));
  }, [isStared, roomId, setStars]);

  return (
    <section
      id={roomId}
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
        {props.staycount === null ? (
          <span className={styles.count}>{"データがありません"}</span>
        ) : (
          <span className={styles.count}>{props.staycount + "人"}</span>
        )}
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
