export type Error = {
  type: "error";
  message: string;
};

export type StayCounts = {
  [key in string]: {
    name: string;
    areas: {
      name: string;
      rooms: { name: string; staycount: number }[];
    }[];
  };
};

export type StayCountsResponse = {
  type: "succeeded";
  content: {
    staycounts: StayCounts;
  };
};

export type Response = StayCounts | Error;
