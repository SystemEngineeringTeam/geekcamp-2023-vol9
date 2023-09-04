import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import styles from "@/styles/Congestion.module.scss";
import { builds } from "@/const";
import Room from "../ui/room";

export default function CongestionComponent() {
  const [roomId, setRoomId] = useState<undefined | string>(undefined);
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
      {builds.map((build) =>
        build.areas.map((area) =>
          area.rooms.map((room) => {
            return <Room key={room} roomId={`${build.id}-${room}`} />;
          })
        )
      )}
    </div>
  );
}
