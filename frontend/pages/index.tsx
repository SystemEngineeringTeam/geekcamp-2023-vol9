import { Noto_Sans_JP } from "next/font/google";
import Builds from "@/components/page/Builds";
import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { headerShadowState } from "@/components/recoil/state";

const noto = Noto_Sans_JP({
  subsets: ["latin"],
  weight: ["100", "300", "500", "700", "900"],
});

export default function Home() {
  const [_, setHeaderShadow] = useRecoilState(headerShadowState);

  useEffect(() => {
    setHeaderShadow(false);
  }, [setHeaderShadow]);

  return <Builds />;
}
