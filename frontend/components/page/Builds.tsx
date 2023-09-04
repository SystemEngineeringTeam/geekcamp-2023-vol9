import styles from "@/styles/Builds.module.scss";
import { stayCountsResponse } from "@/const";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { stayCountsState } from "../recoil/state";
import { getBottomSpace, getGap } from "../util/util";

export default function Builds() {
  const [stayCounts, setStayCounts] = useRecoilState(stayCountsState);
  const router = useRouter();

  useEffect(() => {
    const res = stayCountsResponse;
    if (res.type === "succeeded") {
      setStayCounts(res.content.staycounts);
    }
  }, [setStayCounts]);

  return (
    <div className={styles.builds}>
      {Object.keys(stayCounts).map((key) => (
        <div className={styles.build} key={key}>
          <div
            className={styles.image_container}
            style={{
              backgroundImage: `url(/images/${key}.png)`,
            }}
          ></div>

          <div
            className={styles.selecter}
            style={{ paddingBottom: `${getBottomSpace(key)}px` }}
          >
            {stayCounts[key].areas.map((area) => (
              <div
                className={styles.area}
                style={{ marginTop: `${getGap(key)}px` }}
                key={area.name}
              >
                <h2 className={styles.area_name}>{area.name}</h2>

                {area.rooms.map((room) => (
                  <span
                    className={styles.room}
                    onClick={() =>
                      router.push({
                        pathname: "/congestion",
                        query: { roomId: `${key}-${area.name}-${room.name}` },
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
