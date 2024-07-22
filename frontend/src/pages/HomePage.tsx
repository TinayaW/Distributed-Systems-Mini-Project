import React, { useState } from 'react';
import { getUserByUserName } from '../api/UserEndpoints';
import axios from 'axios';

const HomePage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const handleLogin = async () => {
    if (!username) {
      setErrorMessage("Username is required");
      return;
    } else if (!password) {
      setErrorMessage("Password is required");
      return;
    }
    try {
      const response = await getUserByUserName(axios, username);
      const userData = response.data; 

      if (userData.userpassword === password) {
        window.location.href = `/user-home?id=${userData.id}`;
      } else {
        setErrorMessage("Incorrect password");
      }
    } catch (error) {
      setErrorMessage("User not found");
    }
  };

  const handleCreateAccount = () => {
    window.location.href = '/create-account'; 
  };

  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', minHeight: '100vh' }}>
      <h1 style={{ textAlign: 'center' }}>Home Page</h1>
      <div style={{ textAlign: 'center' }}>
        <input 
          type="text" 
          placeholder="Username" 
          style={{ marginBottom: '10px' }} 
          value={username} 
          onChange={(e) => setUsername(e.target.value)} 
        />
        <br />
        <input 
          type="password" 
          placeholder="Password" 
          style={{ marginBottom: '10px' }} 
          value={password} 
          onChange={(e) => setPassword(e.target.value)} 
        />
        <br />
        <button onClick={handleLogin} style={{ marginBottom: '10px' }}>Login</button>
        <br />
        <div style={{ color: 'red', marginBottom: '20px' }}>{errorMessage}</div> 
        <button onClick={handleCreateAccount}>Create Account</button>
      </div>
    </div>
  );
}

export default HomePage;
