import { useRouter } from "next/router";
import { useEffect, useRef, useState } from "react";
import styles from "@/styles/Congestion.module.scss";
import Room from "@/components/ui/room/room";
import { useStaycount } from "@/hooks/useStaycount";
import { useLocalStorage } from "@mantine/hooks";
import { useHistories } from "@/hooks/useHistories";

export default function CongestionComponent() {
  const [_, { init }] = useHistories();
  const [roomId, setRoomId] = useState<undefined | string>(undefined);
  const staycounts = useStaycount();
  const [stars, setStars] = useLocalStorage<string[]>({
    key: "stars",
    defaultValue: [],
  });
  const router = useRouter();
  const mounted = useRef(false);

  function addStar(roomId: string) {
    setStars((s) => [...s, roomId]);
  }

  function removeStar(roomId: string) {
    setStars((s) => s.filter((star) => star !== roomId));
  }

  useEffect(() => {
    if (!mounted.current) {
      void init();
      mounted.current = true;
    }
  }, [init]);

  useEffect(() => {
    if (router.isReady) {
      setRoomId(router.query.roomId as string);
    }
  }, [router]);

  useEffect(() => {
    if (roomId) {
      document.getElementById(roomId)?.scrollIntoView({
        behavior: "smooth",
        block: "center",
      });
    }
  }, [roomId]);

  return (
    <div className={styles.congestion}>
      {staycounts.map((staycount) => (
        <div key={staycount.building}>
          <h2 className={styles.building_name}>{staycount.building}</h2>

          <div className={styles.building} key={staycount.building}>
            {staycount.floors.map((floor) =>
              floor.rooms.map((room) => {
                return (
                  <Room
                    key={room.name}
                    id={room.id}
                    name={room.name}
                    building={staycount.building}
                    staycount={room.staycount}
                    isSelect={roomId === room.id.toString()}
                    isStar={stars.includes(room.id.toString())}
                    addStar={addStar}
                    removeStar={removeStar}
                  />
                );
              })
            )}
          </div>
        </div>
      ))}
    </div>
  );
}
