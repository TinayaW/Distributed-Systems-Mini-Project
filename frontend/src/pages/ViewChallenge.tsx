import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';
import { getChallenge, deleteChallenge } from '../api/ChallengeEndpoints';
import { createSubmission } from '../api/SubmissionEndpoints';
import ReactMarkdown from 'react-markdown'
import { InputWrapper } from "../components/InputWrapper";
import ZipFileIcon from '@mui/icons-material/FolderZip';
import InfoIcon from '@mui/icons-material/Info';

interface Challenge {
    id: string;
    title: string;
    templatefile: string;
    readmefile: string;
    difficulty: string;
    testfasesfile: string;
    authorid: number;
}

const ViewChallengePage: React.FC = () => {
    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    const id = searchParams.get('id');
    const challengeId = searchParams.get('challengeId');
    const [challenge, setChallenge] = useState<Challenge | null>(null);
    const [submissionFontColor, setSubmissionFontColor] = useState<string>('red');
    const [submissionFile, setsubmissionFile] = useState({} as FileList);

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

    const createChallengePage = () => {
        window.location.href = `/create-challenge?id=${id}`;
    };
    const downloadTemplate = () => {
        if (challenge!.templatefile) {
            const decodedData = atob(challenge!.templatefile);
            const arrayBuffer = new Uint8Array(decodedData.length);
            for (let i = 0; i < decodedData.length; i++) {
            arrayBuffer[i] = decodedData.charCodeAt(i);
            }
            const blob = new Blob([arrayBuffer], { type: 'application/zip' });
            const url = URL.createObjectURL(blob);
            const link = document.createElement('a');
            link.href = url;
            link.download = `challenge-${challengeId}-template.zip`;
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
            URL.revokeObjectURL(url);
        }
    };

    const editThisChallenges = async () => {
        window.location.href = `/edit-challenge?id=${id}&challengeId=${challengeId}`;
    };
    
    const deleteThisChallenges = async () => {
        await deleteChallenge(axios, challengeId!);
        window.location.href = `/user-challenges?id=${id}`;
    };

    const clearAllInputs = () => {
        setsubmissionFile({} as FileList);
    }

    const addSubmission = async () => {
        try {
            const randomNumber = Math.floor(Math.random() * 1000000);
            const submissionId = randomNumber.toString();

            const formData = new FormData();
            formData.append('file', submissionFile[0]);
            formData.append('id', submissionId);
            formData.append('fileName', 'submission');
            formData.append('fileExtension', 'zip');
            formData.append('userId', id!);
            formData.append('challengeId', challengeId!);
            const response = await createSubmission(axios, formData);
            clearAllInputs();
            window.location.href = `/view-challenge?id=${id}&challengeId=${challengeId}`;
        } catch (error) {
            console.error('Failed to create challenge:', error);
        }
    };

    const onSubmissionFileChange = (e : any) => {
        setsubmissionFile(prev => ({...prev, ...e.target.files}));
        setSubmissionFontColor('green');

    }

    const viewLeaderboard = () => {
        window.location.href = `/leaderboard?id=${id}&challengeId=${challengeId}`;
    }

    useEffect(() => {
        if (challengeId) {
            getChallenge(axios, challengeId)
                .then(response => {
                    setChallenge(response.data);
                })
                .catch(error => {
                    console.error('Error fetching challenge:', error);
                });
        }
    }, [challengeId]);

    return (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', minHeight: '100vh' }}>
        <nav>
            <ul style={{ display: 'flex', justifyContent: 'center', listStyle: 'none', padding: 0 }}>
            <li style={{ marginRight: '10px' }}><button onClick={userHome} >User</button></li>
            <li style={{ marginRight: '10px' }}><button onClick={challengeHome} >Challenges</button></li>
            <li><button onClick={submissionHome} >Submissions</button></li>
            </ul>
        </nav>

        {challenge && (
                <div>
                    {challenge.authorid === parseInt(id!) && (
                    <div>
                        <button onClick={editThisChallenges} style={{ marginRight: '30px' }}>Edit</button>
                        <button onClick={deleteThisChallenges} >Delete</button>
                    </div>
                )}
                    <h2>{challenge.title}</h2>
                    <button onClick = {viewLeaderboard}>View Leaderboard</button>
                    <p>Difficulty: {challenge.difficulty}</p>
                    <p>Author: {challenge.authorid}</p>
                    <p>Challenge ID: {challengeId}</p>
                    <ReactMarkdown>{atob(challenge.readmefile)}</ReactMarkdown>
                    <button onClick={downloadTemplate}>Download Template File</button>
                    <br />
                    <div style={{display:'flex', alignItems: 'center'}}>
                        <ZipFileIcon style={{marginRight:10}}/>
                        <InputWrapper  label="Upload Submission .zip file: "><input onChange={onSubmissionFileChange} id="submissionFileInput" type="file" name="submissionFile" accept=".zip" style={{color: submissionFontColor, fontFamily:'Poppins', marginLeft:20}}/></InputWrapper>
                    </div>
                    <div style={{ padding: '10px', borderRadius: '8px', display: 'flex', alignItems: 'center', background: '#eff7ff'}}>
                        <div style={{color: '#808080'}}>
                            <p style={{ display: 'flex', alignItems: 'center'}}> 
                                <InfoIcon sx={{marginRight: '0.5rem'}}/>
                                Zip your main.py file and upload
                            </p>
                        </div>
                    </div>
                    <button onClick={addSubmission} style={{ marginBottom: '10px', marginTop: '10px' }}>Add Submission</button>
                </div>
            )}

        <div>
            <button onClick={getAllChallenges} style={{ marginRight: '30px' }} >All Challenges</button>
            <button onClick={getMyChallenges} style={{ marginRight: '30px' }}>My Challenges</button>
            <button onClick={createChallengePage} >Create Challenges</button>
        </div>
        </div>
    );
}

export default ViewChallengePage;
