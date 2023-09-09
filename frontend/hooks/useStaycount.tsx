import { fetchStaycount } from "@/components/api/api";
import { staycountsState } from "@/components/recoil/state";
import { useEffect, useRef } from "react";
import { useRecoilState } from "recoil";
import _ from "lodash";

export function useStaycount() {
  const [staycounts, setStaycounts] = useRecoilState(staycountsState);
  const mounted = useRef(false);

  useEffect(() => {
    (async () => {
      if (!mounted.current) {
        const tmpRes = await fetchStaycount();
        const res = _.cloneDeep(tmpRes);

        if (res === null) return;
        res.staycounts.forEach((building) => {
          building.floors = _.sortBy(building.floors, "floor");
          building.floors.forEach((floor) => {
            floor.rooms = _.sortBy(floor.rooms, "name");
          });
        });

        setStaycounts(res.staycounts);
        mounted.current = true;
      }
    })();
  }, [setStaycounts]);

  return staycounts;
}
