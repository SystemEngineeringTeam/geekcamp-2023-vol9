import { historyState } from "@/components/recoil/state";
import { historiesResponse } from "@/const";
import { useRecoilState } from "recoil";
import _ from "lodash";
import { faker } from "@faker-js/faker";
import { Number24 } from "@/types";
import { fetchStaycountHistory } from "@/components/api/api";

export function useHistories() {
  const [histories, setHistories] = useRecoilState(historyState);

  async function init() {
    const res = await fetchStaycountHistory();
    if (res === null) return;
    setHistories(res.histories);
  }

  function getHistory(roomId: string, date: string): Number24 | null {
    const res = _.cloneDeep(historiesResponse);
    res.histories[roomId] = Array.from({ length: 24 }, (_) =>
      faker.number.int({ min: 0, max: 100 })
    ) as Number24;
    setHistories(res.histories);

    if (res.histories[roomId] === undefined) {
      return null;
    } else {
      return res.histories[roomId];
    }
  }

  return [histories, { init, getHistory }] as const;
}
