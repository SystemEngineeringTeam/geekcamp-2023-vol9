import styles from "./room.module.scss";

export default function Room({
  roomId,
  isSelect,
}: {
  roomId: string;
  isSelect: boolean;
}) {
  return (
    <section id={roomId} className={styles.room} data-select={isSelect}>
      <h2 className={styles.room_name}>{roomId}</h2>
    </section>
  );
}
