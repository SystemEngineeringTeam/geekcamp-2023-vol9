import { StaycountsResponse } from "@/types";

export const staycountsResponse: StaycountsResponse = {
  staycounts: [
    {
      building: "1号館",
      floors: [
        {
          floor: 2,
          rooms: [
            { id: 4, name: "202", staycount: 13 },
            { id: 3, name: "201", staycount: 10 },
          ],
        },
        {
          floor: 1,
          rooms: [
            { id: 100, name: "アメリカフェ", staycount: 13 },
            { id: 2, name: "102", staycount: 13 },
            { id: 1, name: "101", staycount: 10 },
          ],
        },
        {
          floor: 3,
          rooms: [
            { id: 5, name: "301", staycount: 10 },
            { id: 6, name: "302", staycount: 13 },
          ],
        },
        {
          floor: 4,
          rooms: [
            { id: 7, name: "401", staycount: 10 },
            { id: 8, name: "402", staycount: 13 },
          ],
        },
        {
          floor: 5,
          rooms: [
            { id: 9, name: "501", staycount: 10 },
            { id: 10, name: "502", staycount: 13 },
          ],
        },
        {
          floor: 6,
          rooms: [
            { id: 11, name: "601", staycount: 10 },
            { id: 12, name: "602", staycount: 13 },
          ],
        },
      ],
    },
    {
      building: "10号館",
      floors: [
        {
          floor: 1,
          rooms: [
            { id: 13, name: "101", staycount: 10 },
            { id: 14, name: "102", staycount: 13 },
          ],
        },
        {
          floor: 2,
          rooms: [
            { id: 15, name: "201", staycount: 10 },
            { id: 16, name: "202", staycount: 13 },
          ],
        },
        {
          floor: 3,
          rooms: [
            { id: 17, name: "301", staycount: 10 },
            { id: 18, name: "302", staycount: 13 },
          ],
        },
        {
          floor: 4,
          rooms: [
            { id: 19, name: "401", staycount: 10 },
            { id: 20, name: "402", staycount: 13 },
          ],
        },
        {
          floor: 5,
          rooms: [
            { id: 21, name: "501", staycount: 10 },
            { id: 22, name: "502", staycount: 13 },
          ],
        },
        {
          floor: 6,
          rooms: [
            { id: 23, name: "601", staycount: 10 },
            { id: 24, name: "602", staycount: 13 },
          ],
        },
      ],
    },
  ],
};
