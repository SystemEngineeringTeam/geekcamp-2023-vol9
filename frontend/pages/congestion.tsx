import CongestionComponent from "@/components/page/Congestion";
import { headerShadowState } from "@/components/recoil/state";
import { useEffect } from "react";
import { useRecoilState } from "recoil";

export default function Congestion() {
  const [_, setHeaderShadow] = useRecoilState(headerShadowState);

  useEffect(() => {
    setHeaderShadow(true);
  }, [setHeaderShadow]);

  return <CongestionComponent />;
}
