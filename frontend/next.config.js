/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: "https://campuscrowdmonitor-api.sysken.net/api/:path*",
      },
    ];
  },
};

module.exports = nextConfig;
