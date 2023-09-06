import { atom } from "recoil";
import { Staycount } from "@/types";

export const headerShadowState = atom({
  key: "headerShadowState",
  default: false,
});

export const staycountsState = atom<Staycount[]>({
  key: "stayCountsState",
  default: [],
});
