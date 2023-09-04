import { StayCountsResponse } from "@/types";

export const stayCountsResponse: StayCountsResponse = {
  type: "succeeded",
  content: {
    staycounts: {
      bldg1: {
        name: "1号館",
        areas: [
          {
            name: "1F",
            rooms: [
              { name: "101", staycount: 10 },
              { name: "102", staycount: 13 },
            ],
          },
          {
            name: "2F",
            rooms: [
              { name: "201", staycount: 10 },
              { name: "202", staycount: 13 },
            ],
          },
          {
            name: "3F",
            rooms: [
              { name: "301", staycount: 10 },
              { name: "302", staycount: 13 },
            ],
          },
          {
            name: "4F",
            rooms: [
              { name: "401", staycount: 10 },
              { name: "402", staycount: 13 },
            ],
          },
          {
            name: "5F",
            rooms: [
              { name: "501", staycount: 10 },
              { name: "502", staycount: 13 },
            ],
          },
          {
            name: "6F",
            rooms: [
              { name: "601", staycount: 10 },
              { name: "602", staycount: 13 },
            ],
          },
        ],
      },
      bldg10: {
        name: "10号館",
        areas: [
          {
            name: "1F",
            rooms: [
              { name: "玉森研", staycount: 10 },
              { name: "xx研", staycount: 13 },
            ],
          },
        ],
      },
    },
  },
};
