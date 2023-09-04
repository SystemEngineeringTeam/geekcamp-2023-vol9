import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import styles from "@/styles/Congestion.module.scss";
import Room from "@/components/ui/room";
import { useRecoilValue } from "recoil";
import { stayCountsState } from "../recoil/state";

export default function CongestionComponent() {
  const [roomId, setRoomId] = useState<undefined | string>(undefined);
  const stayCounts = useRecoilValue(stayCountsState);
  const router = useRouter();

  useEffect(() => {
    if (router.isReady) {
      setRoomId(router.query.roomId as string);
    }
  }, [router]);

  useEffect(() => {
    if (roomId) {
      console.log(roomId);
      document.getElementById(roomId)?.scrollIntoView({
        behavior: "smooth",
      });
    }
  }, [roomId]);

  return (
    <div className={styles.congestion}>
      {Object.keys(stayCounts).map((key) =>
        stayCounts[key].areas.map((area) =>
          area.rooms.map((room) => {
            const roomIdSnap = `${key}-${area.name}-${room.name}`;
            return (
              <Room
                key={room.name}
                roomId={roomIdSnap}
                isSelect={roomId === roomIdSnap}
              />
            );
          })
        )
      )}
    </div>
  );
}
