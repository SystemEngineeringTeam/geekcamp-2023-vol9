// エラーの型
export type Error = {
  message: string;
};

// 滞在者数の型
export type Staycount = {
  building: string;
  floors: {
    floor: number;
    rooms: {
      id: number;
      name: string;
      staycount: number | null;
    }[];
  }[];
};

// 混雑度の型
export type Congestion = {
  building: string;
  floors: {
    floor: number;
    rooms: {
      id: number;
      name: string;
      congestion: number | null;
    }[];
  }[];
};

export type Number24 = [
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number,
  number
];

// 履歴の型
export type Histories = {
  [key: string]: Number24;
};

// 滞在者数のレスポンスの型
export type StaycountsResponse = {
  staycounts: Staycount[];
};

// 混雑度のレスポンスの型
export type CongestionsResponse = {
  congestions: Congestion[];
};

// 履歴のレスポンスの型
export type HistoriesResponse = {
  histories: Histories;
};
