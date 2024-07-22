import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { getUser, updateUser, deleteUser} from '../api/UserEndpoints';
import axios from 'axios';

interface User {
  username: string;
  fullname: string;
  id: string;
}

const UserHomePage: React.FC = () => {
  const [user, setUser] = useState<User>({
    username: "",
    fullname: "",
    id: ""
  });

  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const id = searchParams.get('id');

  const logOut = () => {
    window.location.href = '/';
  };

  const updateUserData = async () => {
    const newUsername = prompt("Enter new username:");
    const newFullname = prompt("Enter new full name:");

    if (newUsername && newFullname) {
      const updateuser = { id: parseInt(id!), username: newUsername, fullname: newFullname};
      await updateUser(axios, id!, updateuser);

      setUser({
        ...user,
        username: newUsername,
        fullname: newFullname
      });
    }
  };

  const deleteUserData = async () => {
    await deleteUser(axios, id!);
    window.location.href = '/';
  };

  const userHome = () => {
    window.location.href = `/user-home?id=${id}`; 
  };

  const challengeHome = () => {
    window.location.href = `/challenge-home?id=${id}`; 
  };

  const submissionHome = () => {
    window.location.href = `/submission-home?id=${id}`; 
  };

  const fetchUserData = async (userId: string) => {
    const response = await getUser(axios, userId);
    const userData = response.data; 

    const currentUserData: User = {
      username: userData.username,
      fullname: userData.fullname,
      id: userData.id
    };
    setUser(currentUserData);
  };

  useEffect(() => {
    fetchUserData(id!);
  }, [id]);


  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', minHeight: '100vh' }}>
      <nav>
        <ul style={{ display: 'flex', justifyContent: 'center', listStyle: 'none', padding: 0 }}>
          <li style={{ marginRight: '10px' }}><button onClick={userHome} >User</button></li>
          <li style={{ marginRight: '10px' }}><button onClick={challengeHome} >Challenges</button></li>
          <li><button onClick={submissionHome} >Submissions</button></li>
        </ul>
      </nav>

      <div>
        <p>Username: {user.username}</p>
        <p>Full Name: {user.fullname}</p>
        <p>ID: {user.id}</p>
      </div>

      <div>
        <button onClick={logOut} style={{ marginRight: '30px' }} >Log Out</button>
        <button onClick={updateUserData} style={{ marginRight: '30px' }}>Update User</button>
        <button onClick={deleteUserData}>Delete User</button>
      </div>
    </div>
  );
}

export default UserHomePage;
