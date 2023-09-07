import { atom } from "recoil";
import { Histories, Staycount } from "@/types";

export const headerShadowState = atom({
  key: "headerShadowState",
  default: false,
});

export const staycountsState = atom<Staycount[]>({
  key: "stayCountsState",
  default: [],
});

export const historyState = atom<Histories>({
  key: "historyState",
  default: {},
});
