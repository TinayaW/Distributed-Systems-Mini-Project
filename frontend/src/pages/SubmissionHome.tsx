import React, { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { getSubmissionByUserId, getSubmission } from '../api/SubmissionEndpoints';
import axios from 'axios';

interface Submission {
  id: string;
  score: number;  
  challengeId: string;
  userId: string;
  fileName: string;
  fileExtension: string;
  file: string;
}

const SubmissionHomePage: React.FC = () => {
  const [Submission, setSubmission] = useState<Submission[]>([]);

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

  const fetchSubmission = async () => {
    const submission = await getSubmissionByUserId(axios, id!);
    const fetchSubmissions = submission.data;
    setSubmission(fetchSubmissions);
  };

  const handleSubmissionClick = async (submissionId: string) => {
    const submissionData = await getSubmission(axios, submissionId);
    if (submissionData.data.file) {
      const decodedData = atob(submissionData.data.file);
      const arrayBuffer = new Uint8Array(decodedData.length);
      for (let i = 0; i < decodedData.length; i++) {
      arrayBuffer[i] = decodedData.charCodeAt(i);
      }
      const blob = new Blob([arrayBuffer], { type: 'application/zip' });
      const url = URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.download = `submission-${submissionId}-file.zip`;
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      URL.revokeObjectURL(url);
    }
  };

  useEffect(() => {
    fetchSubmission();
  } , [id]);

  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', minHeight: '100vh' }}>
        <nav>
            <ul style={{ display: 'flex', justifyContent: 'center', listStyle: 'none', padding: 0 }}>
            <li style={{ marginRight: '10px' }}><button onClick={userHome} >User</button></li>
            <li style={{ marginRight: '10px' }}><button onClick={challengeHome} >Challenges</button></li>
            <li><button onClick={submissionHome} >Submissions</button></li>
            </ul>
        </nav>

        <h1>My Submission</h1>

        { Submission === null ? (
          <p>No Submission Available.</p>
        ) : (
        <>
          <div>
            {Submission.map(( submission, index) => (
              <div key={index} style={{ marginBottom: '10px', border: '1px solid #ccc', padding: '10px', borderRadius: '5px', cursor: 'pointer' }} onClick={() => handleSubmissionClick(submission.id)}>
                <p>Submission ID: {submission.id}</p>
                <p>Score: {submission.score}</p>
                <p>Challenge ID: {submission.challengeId}</p>
              </div>
            ))}
          </div>
        </>
      )}
    </div>
  );
}

export default SubmissionHomePage;
