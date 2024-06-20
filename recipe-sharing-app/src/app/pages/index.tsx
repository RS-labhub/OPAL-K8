"use client"
import { useState } from 'react'
import axios from 'axios'

const Home = () => {
  const [username, setUsername] = useState('')
  const [message, setMessage] = useState('')

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    try {
      const permissionResponse = await axios.post('/api/permissions', { username })
      if (permissionResponse.data.allowed) {
        setMessage('You are authorized to perform this operation.')
      } else {
        setMessage('You are not authorized to perform this operation.')
      }
    } catch (error) {
      setMessage('Error checking permissions')
    }
  }

  return (
    <div>
      <h1>Recipe Sharing App</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>
            Username:
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </label>
        </div>
        <button type="submit">Check Permissions</button>
      </form>
      {message && <p>{message}</p>}
    </div>
  )
}

export default Home
