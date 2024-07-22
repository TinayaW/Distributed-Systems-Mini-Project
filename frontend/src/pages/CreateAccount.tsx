import React, { useState } from 'react';
import { createUser } from '../api/UserEndpoints';
import axios from 'axios';

const CreateAccountPage: React.FC = () => {
  const [username, setUsername] = useState('');
  const [fullname, setFullName] = useState('');
  const [userpassword, setUserPassword] = useState('');
  const [id, setId] = useState('');
  const [error, setError] = useState('');

  const generateId = () => {
    const randomNumber = Math.floor(Math.random() * 1000000);
    setId(randomNumber.toString()); 
  };

  const handleSubmit = async () => {
    if (!username) {
      setError('Username is required');
      return;
    } else if (!fullname) {
      setError('Full name is required');
      return;
    } else if (!userpassword) {
      setError('Password is required');
      return;
    } else if (!id) {
      setError('ID is required. Click "Generate ID" to generate one.');
      return;
    }

    const user = { id: parseInt(id), username, fullname, userpassword};
    try {
      await createUser(axios, user);
      window.location.href = '/';
    } catch (error) {
      console.error('Error creating account:', error);
    }
  };

  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', minHeight: '100vh' }}>
      <h1 style={{ textAlign: 'center' }}>Create Account</h1>
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
          type="text" 
          placeholder="Full Name" 
          style={{ marginBottom: '10px' }} 
          value={fullname} 
          onChange={(e) => setFullName(e.target.value)}
        />
        <br />
        <input 
          type="password" 
          placeholder="Password" 
          style={{ marginBottom: '10px' }} 
          value={userpassword} 
          onChange={(e) => setUserPassword(e.target.value)}
        />
        <br />
        <button onClick={generateId} style={{ marginBottom: '10px' }}>Generate ID</button>
        <p>Auto-generated ID: {id}</p>
        <br />
        <button onClick={handleSubmit}>Create Account</button>
        {error && <p style={{ color: 'red' }}>{error}</p>}
      </div>
    </div>
  );
}

export default CreateAccountPage;
