// pages/api/permissions.ts
import type { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';
import { getUserData } from '../../../../lib/db';

type Data = {
  allowed: boolean;
};

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method === 'POST') {
    const { username } = req.body;

    try {
      const user = await getUserData(username);
      if (!user) {
        return res.status(404).json({ allowed: false });
      }

      const data = {
        input: {
          user: { karma: user.karma, location: user.location }
        }
      };

      const response = await axios.post('http://localhost:8181/v1/data/example/allow', data);
      res.status(200).json({ allowed: response.data.result });
    } catch (error) {
      res.status(500).json({ allowed: false });
    }
  } else {
    res.setHeader('Allow', ['POST']);
    res.status(405).end(`Method ${req.method} Not Allowed`);
  }
}
