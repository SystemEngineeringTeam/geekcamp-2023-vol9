export function getStyle(buildId: string): {
  gap: number;
  bottomSpace: number;
  widthRatio: number;
} {
  switch (buildId) {
    case "1号館":
      return { widthRatio: 1, gap: 6.5, bottomSpace: 0 };
    case "10号館":
      return { widthRatio: 1, gap: 15, bottomSpace: 17 };
    case "4号館":
      return { widthRatio: 0.7, gap: 15, bottomSpace: 8 };
    case "4号館別館":
      return { widthRatio: 0.87, gap: 15, bottomSpace: 8 };
    default:
      return { widthRatio: 1, gap: 6.5, bottomSpace: 0 };
  }
}
