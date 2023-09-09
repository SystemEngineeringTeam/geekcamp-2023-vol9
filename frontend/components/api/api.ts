import { CongestionsResponse, HistoriesResponse, StaycountsResponse } from "@/types";
import axios from "axios";

const BASE_URL = "https://campuscrowdmonitor-api.sysken.net";

// GETリクエストを送信する
export const commonGetFetch = async <T>(path: string): Promise<T | null> => {
  const url = new URL(path, BASE_URL).href;

  return axios
    .get(url)
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
  return commonGetFetch<StaycountsResponse>("/api/v1/staycount/get/");
};

// 混雑度を取得する
export const fetchCongestion =
  async (): Promise<CongestionsResponse | null> => {
    return commonGetFetch<CongestionsResponse>("/api/v1/congestion/get/");
  };

// 今日の滞在者数の履歴を取得する
export const fetchStaycountHistory =
  async (): Promise<HistoriesResponse | null> => {
    return commonGetFetch<HistoriesResponse>("/api/v1/staycount/histories/");
  };
