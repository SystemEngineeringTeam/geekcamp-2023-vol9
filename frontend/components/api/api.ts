import { CongestionsResponse, StaycountsResponse } from "@/types";
import axios from "axios";

// GETリクエストを送信する
export const commonGetFetch = async <T>(path: string): Promise<T | null> => {
  return axios
    .get(path)
    .then((res) => {
      return res.data as T;
    })
    .catch((err) => {
      console.log(err);
      return null;
    });
};

// 滞在者数を取得する
export const fetchStaycount = async (): Promise<StaycountsResponse | null> => {
  return commonGetFetch<StaycountsResponse>("https://campuscrowdmonitor-api.sysken.net/api/v1/staycount/get/");
};

// 混雑度を取得する
export const fetchCongestion = async (): Promise<CongestionsResponse | null> => {
  return commonGetFetch<CongestionsResponse>("https://campuscrowdmonitor-api.sysken.net/api/v1/congestion/get/");
};
