import { StayCounts, StayCountsResponse } from "@/types";
import { atom } from "recoil";

export const headerShadowState = atom({
  key: "headerShadowState",
  default: false,
});

export const stayCountsState = atom<StayCounts>({
  key: "stayCountsState",
  default: {},
});
