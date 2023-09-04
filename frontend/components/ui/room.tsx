import styles from './room.module.scss';

export default function Room({ roomId }: { roomId: string }) {
  return <section id={roomId} className={styles.room}>room: {roomId}</section>;
}
