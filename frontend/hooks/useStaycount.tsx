import { staycountsState } from "@/components/recoil/state";
import { staycountsResponse } from "@/const";
import _ from "lodash";
import { useEffect } from "react";
import { useRecoilState } from "recoil";

export function useStaycount() {
  const [staycounts, setStaycounts] = useRecoilState(staycountsState);

  useEffect(() => {
    return setStaycounts((v) => {
      if (v.length !== 0) return v;

      const res = _.cloneDeep(staycountsResponse);
      res.staycounts.forEach((building) => {
        building.floors = _.sortBy(building.floors, "floor");
        building.floors.forEach((floor) => {
          floor.rooms = _.sortBy(floor.rooms, "name");
        });
      });

      return res.staycounts;
    });
  }, [setStaycounts]);

  return staycounts;
}
