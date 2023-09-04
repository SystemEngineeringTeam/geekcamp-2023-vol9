import Builds from "@/components/page/Builds";
import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { headerShadowState } from "@/components/recoil/state";

export default function Home() {
  const [_, setHeaderShadow] = useRecoilState(headerShadowState);

  useEffect(() => {
    setHeaderShadow(false);
  }, [setHeaderShadow]);

  return <Builds />;
}
