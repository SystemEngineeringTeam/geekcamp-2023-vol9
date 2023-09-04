export type Area = {
  name: string;
  rooms: string[];
};

export type Build = {
  id: string;
  name: string;
  image: string;
  areas: Area[];
  gap: number;
  bottomSpace: number;
};
