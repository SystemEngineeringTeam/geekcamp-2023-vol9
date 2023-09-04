import { Area, Build } from "@/types";

const build1Areas: Area[] = [
  {
    name: "1F",
    rooms: ["101", "102", "103", "104", "アメリカフェ"],
  },
  {
    name: "2F",
    rooms: ["201", "202", "203", "204"],
  },
  {
    name: "3F",
    rooms: ["301", "302", "303", "304"],
  },
  {
    name: "4F",
    rooms: ["401", "402", "403", "404"],
  },
  {
    name: "5F",
    rooms: ["501", "502", "503", "504"],
  },
  {
    name: "6F",
    rooms: ["601", "602", "603", "604"],
  },
  {
    name: "7F",
    rooms: ["701", "702", "703", "704"],
  },
];

const build10Areas: Area[] = [
  {
    name: "1F",
    rooms: ["101", "102", "103", "104"],
  },
  {
    name: "2F",
    rooms: ["201", "202", "203", "204"],
  },
  {
    name: "3F",
    rooms: ["301", "302", "303", "304"],
  },
  {
    name: "4F",
    rooms: ["401", "402", "403", "404"],
  },
  {
    name: "5F",
    rooms: ["501", "502", "503", "504"],
  },
  {
    name: "6F",
    rooms: ["601", "602", "603", "604"],
  },
];

export const builds: Build[] = [
  {
    id: "build-1",
    name: "1号館",
    image: "/images/build-1.png",
    areas: build1Areas,
    gap: 6.5,
    bottomSpace: 0,
  },
  {
    id: "build-10",
    name: "10号館",
    image: "/images/build-10.png",
    areas: build10Areas,
    gap: 17,
    bottomSpace: 15,
  },
];
