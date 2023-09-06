import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import styles from "@/styles/Congestion.module.scss";
import Room from "@/components/ui/room";
import { useRecoilValue } from "recoil";
import { staycountsState } from "../recoil/state";

export default function CongestionComponent() {
  const [roomId, setRoomId] = useState<undefined | string>(undefined);
  const staycounts = useRecoilValue(staycountsState);
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
      {staycounts.map((staycount) =>
        staycount.floors.map((floor) =>
          floor.rooms.map((room) => {
            return (
              <Room
                key={room.name}
                id={room.id}
                name={room.name}
                building={staycount.building}
                staycount={room.staycount}
                isSelect={roomId === room.id.toString()}
              />
            );
          })
        )
      )}
    </div>
  );
}
