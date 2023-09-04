import styles from "@/styles/Builds.module.scss";
import { builds } from "@/const";
import { useRouter } from "next/router";

export default function Builds() {
  const router = useRouter();

  return (
    <div className={styles.builds}>
      {builds.map((build) => (
        <div className={styles.build} key={build.id}>
          <div
            className={styles.image_container}
            style={{
              backgroundImage: `url(${build.image})`,
            }}
          ></div>

          <div
            className={styles.selecter}
            style={{ paddingBottom: `${build.bottomSpace}px` }}
          >
            {build.areas.map((area) => (
              <div
                className={styles.area}
                style={{ marginTop: `${build.gap}px` }}
                key={area.name}
              >
                <h2 className={styles.area_name}>{area.name}</h2>

                {area.rooms.map((room) => (
                  <span
                    className={styles.room}
                    onClick={() =>
                      router.push({
                        pathname: "/congestion",
                        query: { roomId: `${build.id}-${room}` },
                      })
                    }
                    key={room}
                  >
                    {room}
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
