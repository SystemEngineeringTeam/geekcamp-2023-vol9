export function getBottomSpace(buildId: string): number {
  switch (buildId) {
    case "bldg1":
      return 0;
    case "bldg10":
      return 15;
    default:
      return 0;
  }
}

export function getGap(buildId: string): number {
  switch (buildId) {
    case "bldg1":
      return 6.5;
    case "bldg10":
      return 17;
    default:
      return 0;
  }
}
