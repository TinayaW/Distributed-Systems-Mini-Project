import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';
import { getUserChallenges } from '../api/ChallengeEndpoints';

interface Challenge {
    id: string;
    title: string;
    templateFile: Array<number>;
    readmefile: Array<number>;
    difficulty: string;
    testfasesfile: Array<number>;
    authorid: number;
}

interface ChallengeMin {
    id: string;
    title: string;
    difficulty: string;
    authorid: number;
}

const UserChallengePage: React.FC = () => {
    const [Challenges, setChallenges] = useState<ChallengeMin[]>([]); 

    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    const id = searchParams.get('id');

    const userHome = () => {
        window.location.href = `/user-home?id=${id}`; 
    };

    const challengeHome = () => {
        window.location.href = `/challenge-home?id=${id}`; 
    };

    const submissionHome = () => {
        window.location.href = `/submission-home?id=${id}`; 
    };

    const getAllChallenges = () => {
        window.location.href = `/challenge-home?id=${id}`;
    };

    const getMyChallenges = () => {
        window.location.href = `/user-challenges?id=${id}`;
      };
    
      const createChallenge = () => {
        window.location.href = `/create-challenge?id=${id}`;
      };
    

    const fetchChallengeData = async () => {
        const response = await getUserChallenges(axios, id!);
        const challenges = response.data; 
        setChallenges(challenges);
    };

    const handleChallengeClick = (challengeId: string) => {
        window.location.href = `/view-challenge?id=${id}&challengeId=${challengeId}`;
    };
    
    useEffect(() => {
        fetchChallengeData();
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

        {Challenges === null ? (
            <p>No challenges available.</p>
        ) : (
            <>
                <div>
                    {Challenges.map((challenge, index) => (
                        <div key={index} style={{ marginBottom: '10px', border: '1px solid #ccc', padding: '10px', borderRadius: '5px', cursor: 'pointer' }} onClick={() => handleChallengeClick(challenge.id)}>
                            <p>Challenge Title: {challenge.title}</p>
                            <p>Challenge Difficulty: {challenge.difficulty}</p>
                            <p>Challenge Author ID: {challenge.authorid}</p>
                        </div>
                    ))}
                </div>
            </>
        )}

        <div>
            <button onClick={getAllChallenges} style={{ marginRight: '30px' }} >All Challenges</button>
            <button onClick={getMyChallenges} style={{ marginRight: '30px' }}>My Challenges</button>
            <button onClick={createChallenge} >Create Challenges</button>
        </div>
        </div>
    );
}

export default UserChallengePage;
