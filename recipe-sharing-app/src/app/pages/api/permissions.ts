// pages/api/permissions.ts
import type { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';

export async function handler(
  req: NextApiRequest,
  res: NextApiResponse<any>
) {
  if (req.method === 'POST') {
    const { username } = req.body;

    try {
      // Fetch user data from MySQL or another data source
      const userData = await fetchUserDataFromMySQL(username);

      // Construct the data payload for OPAL policy evaluation
      const data = {
        input: {
          user: {
            karma: userData.karma,
            location: userData.location,
          },
        },
      };

      // Send the policy evaluation request to the OPAL server
      const response = await axios.post('http://localhost:7002/v1/data/example/allow', data);

      // Return the result of the policy evaluation to the client
      res.status(200).json({ allowed: response.data.result });
    } catch (error) {
      console.error('Error checking permissions:', error);
      res.status(500).json({ error: 'Error checking permissions' });
    }
  } else {
    res.setHeader('Allow', ['POST']);
    res.status(405).end(`Method ${req.method} Not Allowed`);
  }
}

async function fetchUserDataFromMySQL(username: string) {
  // Implement logic to fetch user data from MySQL
  // Example implementation using a mock database
  const mockDatabase = {
    user1: { karma: 150, location: 'allowed_location' },
    // Add more users as needed
  };

  // Simulate fetching user data (replace with actual database query)
  return mockDatabase[username] || { karma: 0, location: '' };
}
