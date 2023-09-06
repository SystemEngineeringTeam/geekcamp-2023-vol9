import styles from "@/styles/Builds.module.scss";
import { staycountsResponse } from "@/const";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { staycountsState } from "../recoil/state";
import { getStyle } from "../util/util";
import _ from "lodash";

export default function Builds() {
  const [staycounts, setStaycounts] = useRecoilState(staycountsState);
  const router = useRouter();

  useEffect(() => {
    const res = _.cloneDeep(staycountsResponse);
    res.staycounts.forEach((building) => {
      building.floors = _.sortBy(building.floors, "floor");
      building.floors.forEach((floor) => {
        floor.rooms = _.sortBy(floor.rooms, "name");
      });
    });
    setStaycounts(res.staycounts);
  }, [setStaycounts]);

  return (
    <div className={styles.builds}>
      {staycounts.map((staycount) => (
        <div className={styles.build} key={staycount.building}>
          <div
            className={styles.image_container}
            style={{
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
