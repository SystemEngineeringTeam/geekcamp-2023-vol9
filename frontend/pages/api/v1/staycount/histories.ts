import { historiesResponse } from "@/const";
import type { NextApiRequest, NextApiResponse } from "next";

const handler = (req: NextApiRequest, res: NextApiResponse) => {
  res.status(200).json(historiesResponse);
};

export default handler;
