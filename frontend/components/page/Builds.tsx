import styles from "@/styles/Builds.module.scss";
import { useRouter } from "next/router";
import { getStyle } from "../util/util";
import { useStaycount } from "@/hooks/useStaycount";
import _ from "lodash";

export default function Builds() {
  const staycounts = useStaycount();
  const router = useRouter();

  return (
    <div className={styles.builds}>
      {staycounts.map((staycount) => (
        <div className={styles.build} key={staycount.building}>
          <div
            className={styles.image_container}
            style={{
              ["--width-ratio" as string]: getStyle(staycount.building)
                .widthRatio,
              backgroundImage: `url(/images/${staycount.building}.png)`,
            }}
          ></div>

          <div
            className={styles.selecter}
            style={{
              paddingBottom: `${getStyle(staycount.building).bottomSpace}px`,
            }}
          >
            {staycount.floors.map((floor) => (
              <div
                className={styles.area}
                style={{ marginTop: `${getStyle(staycount.building).gap}px` }}
                key={floor.floor}
              >
                <h2 className={styles.area_name}>{floor.floor}F</h2>

                {floor.rooms.map((room) => (
                  <span
                    className={styles.room}
                    onClick={() =>
                      router.push({
                        pathname: "/congestion",
                        query: { roomId: `${room.id}` },
                      })
                    }
                    key={room.name}
                  >
                    {room.name}
                  </span>
                ))}
              </div>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
}
