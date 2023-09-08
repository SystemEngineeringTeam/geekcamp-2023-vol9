import "@/styles/globals.css";
import styles from "@/styles/Home.module.scss";
import Head from "next/head";
import type { AppProps } from "next/app";
import Header from "@/components/base/header/header";
import { Noto_Sans_JP } from "next/font/google";
import { RecoilRoot } from "recoil";

const noto = Noto_Sans_JP({
  subsets: ["latin"],
  weight: ["100", "300", "500", "700", "900"],
});

const APP_NAME = "CampusCrowdMonitor";
const DESCRIPTION =
  "CampusCrowdMonitor は大学の各施設の混雑状況を知れるアプリです.";
const SITE_URL = "https://campus-crowd-monitor.vercel.app/";
const IMG_URL = "https://campus-crowd-monitor.vercel.app/ogp.png";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <Head>
        <title>{APP_NAME}</title>
        <meta name="description" content={DESCRIPTION} />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.png" />

        <meta property="og:title" content={APP_NAME} />
        <meta property="og:description" content={DESCRIPTION} />
        <meta property="og:type" content="website" />
        <meta property="og:url" content={SITE_URL} />
        <meta property="og:image" content={IMG_URL} />
        <meta property="og:site_name" content={APP_NAME} />
      </Head>

      <div className={noto.className}>
        <Header />
        <main className={styles.main}>
          <Component {...pageProps} />
        </main>
      </div>
    </RecoilRoot>
  );
}
