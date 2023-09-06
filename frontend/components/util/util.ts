export function getStyle(buildId: string): {
  gap: number;
  bottomSpace: number;
} {
  switch (buildId) {
    case "1号館":
      return { gap: 6.5, bottomSpace: 0 };
    case "10号館":
      return { gap: 15, bottomSpace: 17 };
    default:
      return { gap: 6.5, bottomSpace: 0 };
  }
}
