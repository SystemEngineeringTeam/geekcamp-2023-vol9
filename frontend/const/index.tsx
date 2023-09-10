import { HistoriesResponse, StaycountsResponse } from "@/types";

export const labels = Array.from({ length: 24 }, (_, index) => `${index}時`);

class ID {
  private index: number = 0;

  public next(): number {
    return this.index++;
  }
}

const id = new ID();

export const staycountsResponse: StaycountsResponse = {
  staycounts: [
    {
      building: "1号館",
      floors: [
        {
          floor: 2,
          rooms: [{ id: id.next(), name: "アメリカフェ", staycount: 3 }],
        },
        {
          floor: 1,
          rooms: [{ id: id.next(), name: "キャリアセンター", staycount: 0 }],
        },
        {
          floor: 3,
          rooms: [
            { id: id.next(), name: "301", staycount: 5 },
            { id: id.next(), name: "302", staycount: 2 },
          ],
        },
        {
          floor: 4,
          rooms: [
            { id: id.next(), name: "401", staycount: 0 },
            { id: id.next(), name: "402", staycount: 0 },
          ],
        },
        {
          floor: 5,
          rooms: [
            { id: id.next(), name: "501", staycount: 0 },
            { id: id.next(), name: "502", staycount: 0 },
          ],
        },
        {
          floor: 6,
          rooms: [
            { id: id.next(), name: "601", staycount: 0 },
            { id: id.next(), name: "602", staycount: 0 },
          ],
        },
        {
          floor: 7,
          rooms: [
            { id: id.next(), name: "701", staycount: 0 },
            { id: id.next(), name: "702", staycount: 0 },
            { id: id.next(), name: "鳥居教授", staycount: 0 },
            { id: id.next(), name: "鳥居研", staycount: 0 },
          ],
        },
      ],
    },
    {
      building: "10号館",
      floors: [
        {
          floor: 1,
          rooms: [],
        },
        {
          floor: 2,
          rooms: [
            { id: id.next(), name: "G2201", staycount: 0 },
            { id: id.next(), name: "G2202", staycount: 0 },
            { id: id.next(), name: "G2203", staycount: 0 },
            { id: id.next(), name: "G2204", staycount: 0 },
          ],
        },
        {
          floor: 3,
          rooms: [
            { id: id.next(), name: "G2305", staycount: 0 },
            { id: id.next(), name: "G2306", staycount: 0 },
            { id: id.next(), name: "G2307", staycount: 0 },
            { id: id.next(), name: "G2308", staycount: 0 },
          ],
        },
        {
          floor: 4,
          rooms: [
            { id: id.next(), name: "G2409", staycount: 0 },
            { id: id.next(), name: "G2410", staycount: 0 },
            { id: id.next(), name: "G2411", staycount: 0 },
            { id: id.next(), name: "G2412", staycount: 0 },
          ],
        },
        {
          floor: 5,
          rooms: [
            { id: id.next(), name: "G2513", staycount: 0 },
            { id: id.next(), name: "G2514", staycount: 0 },
            { id: id.next(), name: "G2515", staycount: 0 },
            { id: id.next(), name: "G2516", staycount: 0 },
          ],
        },
        {
          floor: 6,
          rooms: [
            { id: id.next(), name: "G2617", staycount: 0 },
            { id: id.next(), name: "G2618", staycount: 0 },
            { id: id.next(), name: "G2619", staycount: 0 },
            { id: id.next(), name: "G2620", staycount: 0 },
          ],
        },
        {
          floor: 7,
          rooms: [
            { id: id.next(), name: "G2621", staycount: 0 },
            { id: id.next(), name: "G2622", staycount: 0 },
            { id: id.next(), name: "G2623", staycount: 0 },
            { id: id.next(), name: "G2624", staycount: 0 },
          ],
        },
      ],
    },
    {
      building: "4号館",
      floors: [
        {
          floor: 1,
          rooms: [
            { id: id.next(), name: "内種教授", staycount: 0 },
            { id: id.next(), name: "内種研", staycount: 1 },
            { id: id.next(), name: "玉森准教授", staycount: 0 },
            { id: id.next(), name: "玉森研", staycount: 3 },
          ],
        },
        {
          floor: 2,
          rooms: [
            { id: id.next(), name: "小野木教授", staycount: 0 },
            { id: id.next(), name: "小野木研", staycount: 0 },
            { id: id.next(), name: "小栗講師", staycount: 0 },
            { id: id.next(), name: "塚田教授", staycount: 0 },
            { id: id.next(), name: "塚田研", staycount: 0 },
            { id: id.next(), name: "森本教授", staycount: 0 },
            { id: id.next(), name: "森本研", staycount: 0 },
          ],
        },
        {
          floor: 3,
          rooms: [
            { id: id.next(), name: "伊藤雅教授", staycount: 0 },
            { id: id.next(), name: "伊藤雅研", staycount: 0 },
            { id: id.next(), name: "中村教授", staycount: 0 },
            { id: id.next(), name: "中村研", staycount: 1 },
            { id: id.next(), name: "菱田教授", staycount: 0 },
            { id: id.next(), name: "菱田研", staycount: 0 },
          ],
        },
        {
          floor: 4,
          rooms: [
            { id: id.next(), name: "401", staycount: 0 },
            { id: id.next(), name: "402", staycount: 0 },
            { id: id.next(), name: "403", staycount: 0 },
            { id: id.next(), name: "404", staycount: 0 },
          ],
        },
      ],
    },
    {
      building: "4号館別館",
      floors: [
        {
          floor: 1,
          rooms: [
            { id: id.next(), name: "事務室", staycount: 0 },
            { id: id.next(), name: "梶准教授", staycount: 0 },
            { id: id.next(), name: "梶研", staycount: 5 },
          ],
        },
        {
          floor: 2,
          rooms: [
            { id: id.next(), name: "伊藤暢浩教授", staycount: 0 },
            { id: id.next(), name: "伊藤暢浩研", staycount: 3 },
            { id: id.next(), name: "河辺教授", staycount: 0 },
            { id: id.next(), name: "河辺研", staycount: 0 },
          ],
        },
        {
          floor: 3,
          rooms: [
            { id: id.next(), name: "内藤教授", staycount: 0 },
            { id: id.next(), name: "内藤研", staycount: 0 },
            { id: id.next(), name: "松河教授", staycount: 0 },
            { id: id.next(), name: "松河研", staycount: 0 },
            { id: id.next(), name: "シス研", staycount: 6 },
          ],
        },
        {
          floor: 4,
          rooms: [
            { id: id.next(), name: "澤野准教授", staycount: 0 },
            { id: id.next(), name: "澤野研", staycount: 0 },
            { id: id.next(), name: "水野教授", staycount: 0 },
            { id: id.next(), name: "水野研", staycount: 0 },
            { id: id.next(), name: "山本教授", staycount: 0 },
            { id: id.next(), name: "山本研", staycount: 0 },
          ],
        },
      ],
    },
  ],
};

export const historiesResponse: HistoriesResponse = {
  histories: {
    "0": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 3, 2, 0, 3, 3, 4, 2, 0, 0, 0, 0, 0, 0,
    ],
    "1": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
    ],
    "2": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 3, 2, 0, 3, 4, 4, 4, 0, 0, 0, 0, 0, 0,
    ],
    "3": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 3, 2, 2, 0, 0, 0, 0, 0, 0, 0,
    ],
    "13": [
      0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    ],
    "39": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0,
    ],
    "41": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 3, 3, 3, 4, 0, 0, 0, 0, 0, 0,
    ],
    "44": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    ],
    "52": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
    ],
    "61": [
      2, 4, 4, 2, 2, 2, 2, 3, 4, 4, 4, 3, 4, 8, 8, 8, 5, 5, 0, 0, 0, 0, 0, 0,
    ],
    "63": [
      0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 2, 3, 3, 3, 0, 0, 0, 0, 0, 0,
    ],
    "70": [
      3, 0, 0, 3, 3, 3, 3, 3, 5, 5, 5, 3, 6, 6, 6, 6, 5, 5, 0, 0, 0, 0, 0, 0,
    ],
    "71": [
      0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    ],
  },
};
