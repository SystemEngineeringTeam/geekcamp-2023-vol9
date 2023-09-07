import { StarFilledIcon, StarOutlineIcon } from "../../util/icons";
import styles from "./room.module.scss";
import Linechart from "../linechart/linechart";

type Props = {
  key: string;
  id: number;
  name: string;
  building: string;
  staycount: number | null;
  isSelect: boolean;
  isStar: boolean;
  addStar: (roomId: string) => void;
  removeStar: (roomId: string) => void;
};

export default function Room(props: Props) {
  const roomId = props.id.toString();

  return (
    <section
      id={roomId}
      className={styles.room}
      data-select={props.isSelect}
      data-star={props.isStar}
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

      {props.isStar ? (
        <StarFilledIcon
          className={styles.star}
          onClick={() => props.removeStar(roomId)}
        />
      ) : (
        <StarOutlineIcon
          className={styles.star}
          onClick={() => props.addStar(roomId)}
        />
      )}

      <Linechart roomId={roomId} />
    </section>
  );
}
