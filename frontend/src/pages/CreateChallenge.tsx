import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';
import { createChallenge } from '../api/ChallengeEndpoints';
import MdFileIcon from '@mui/icons-material/FileCopy';
import ZipFileIcon from '@mui/icons-material/FolderZip';
import InfoIcon from '@mui/icons-material/Info';
import { FormControl, FormControlLabel, FormLabel, Radio, RadioGroup, Snackbar, useFormControl } from "@mui/material";
import Button from "@mui/material/Button";
import { InputWrapper } from "../components/InputWrapper";

interface Challenge {
    id: string;
    title: string;
    templateFile: Array<number>;
    readmefile: Array<number>;
    difficulty: string;
    testfasesfile: Array<number>;
    authorid: number;
}

const CreateChallengePage: React.FC = () => {
    const [challengeid, setchallengeId] = useState('');
    const [challengeTitle, setchallengeTitle] = useState<string>("");
    const [challengeDifficulty, setchallengeDifficulty] = useState<string>("MEDIUM");
    const [testCaseFile, settestCaseFile] = useState({} as FileList);
    const [readmeFile, setReadmeFile] = useState({} as FileList);
    const [templateFile, settemplateFile] = useState({} as FileList);
    const [readmeFontColor, setReadmeFontColor] = useState<string>('red');
    const [templateFontColor, setTemplateFontColor] = useState<string>('red');
    const [testFontColor, setTestFontColor] = useState<string>('red');

    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    const id = searchParams.get('id');

    
    const generateId = () => {
        const randomNumber = Math.floor(Math.random() * 1000000);
        setchallengeId(randomNumber.toString()); 
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

    const getAllChallenges = () => {
        window.location.href = `/challenge-home?id=${id}`;
    };

    const getMyChallenges = () => {
        window.location.href = `/user-challenges?id=${id}`;
    };

    const createChallengePage = () => {
        window.location.href = `/create-challenge?id=${id}`;
    };

    const clearAllInputs = () => {
        setchallengeTitle("");
        setchallengeDifficulty("MEDIUM");
        settestCaseFile({} as FileList);
        setReadmeFile({} as FileList);
        settemplateFile({} as FileList);
    }

    const handleSubmit = async () => {
        try {
            const formData = new FormData();
            formData.append('testcase', testCaseFile[0]);
            formData.append('template', templateFile[0]);
            formData.append('readme', readmeFile[0]);
            formData.append('id', challengeid);
            formData.append('title', challengeTitle);
            formData.append('difficulty', challengeDifficulty);
            formData.append('authorid', id!); 
            const response = await createChallenge(axios, formData);
            clearAllInputs();
            window.location.href = `/user-challenges?id=${id}`;
        } catch (error) {
            console.error('Failed to create challenge:', error);
        }
    };
    

    const onTestCaseFileChange = (e : any) => {
        settestCaseFile(prev => ({...prev, ...e.target.files}));
        setTestFontColor('green');
    }

    const onReadmeFileChange = (e : any) => {
        setReadmeFile(prev => ({...prev, ...e.target.files}));
        setReadmeFontColor('green');
    }

    const onTemplateFileChange = (e : any) => {
        settemplateFile(prev => ({...prev, ...e.target.files}));
        setTemplateFontColor('green');

    }

    return (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', minHeight: '100vh' }}>
        <nav>
            <ul style={{ display: 'flex', justifyContent: 'center', listStyle: 'none', padding: 0 }}>
            <li style={{ marginRight: '10px' }}><button onClick={userHome} >User</button></li>
            <li style={{ marginRight: '10px' }}><button onClick={challengeHome} >Challenges</button></li>
            <li><button onClick={submissionHome} >Submissions</button></li>
            </ul>
        </nav>

        <h1>Create Challenge</h1>

        <InputWrapper label="Challenge Title: " >
            <input type="text" value={challengeTitle} onChange={(e) => setchallengeTitle(e.target.value)} style={{fontFamily:'Poppins'}}/>
        </InputWrapper>

        <button onClick={generateId} style={{ marginBottom: '10px' }}>Generate ID</button>
        <p>Auto-generated ID: {challengeid}</p>

        <FormControl sx={{marginY: '1rem'}} >
                <FormLabel id="demo-controlled-radio-buttons-group">D i f f i c u l t y</FormLabel>
                <RadioGroup
                    aria-labelledby="demo-controlled-radio-buttons-group"
                    name="controlled-radio-buttons-group"
                    value={challengeDifficulty}
                    onChange={(e) => setchallengeDifficulty(e.target.value)}
                >
                    <FormControlLabel value="easy" control={<Radio />} label="EASY" />
                    <FormControlLabel value="medium" control={<Radio />} label="MEDIUM" />
                    <FormControlLabel value="hard" control={<Radio />} label="HARD" />
                </RadioGroup>
            </FormControl>
            <br />

            <div style={{display:'flex', alignItems: 'center'}}>
                <MdFileIcon style={{marginRight:10}}/>
                <InputWrapper  label="Upload readme .md file: "><input onChange={onReadmeFileChange} id="readmeInput" type="file" name="readmeFile" accept=".md" style={{color: readmeFontColor, fontFamily:'Poppins', marginLeft:25}}/></InputWrapper>
            </div>
            <div style={{ padding: '10px', borderRadius: '8px', display: 'flex', alignItems: 'center', background: '#eff7ff'}}>
                <div style={{color: '#808080'}}>
                    <p style={{ display: 'flex', alignItems: 'center'}}> 
                        <InfoIcon sx={{marginRight: '0.5rem'}}/>
                        You can include the following details in the .md file:
                    </p>
                    <ul>
                        <li>Challenge description</li>
                        <li>Input format</li>
                        <li>Constraints</li>
                        <li>Code examples</li>
                    </ul>
                </div>
            </div>

            <div style={{display:'flex', alignItems: 'center'}}>
                <ZipFileIcon style={{marginRight:10}}/>
                <InputWrapper  label="Upload template .zip file: "><input onChange={onTemplateFileChange} id="testFileInput" type="file" name="templateFile" accept=".zip" style={{color: templateFontColor, fontFamily:'Poppins', marginLeft:20}}/></InputWrapper>
            </div>
            <div style={{ padding: '10px', borderRadius: '8px', display: 'flex', alignItems: 'center', background: '#eff7ff'}}>
                <div style={{color: '#808080'}}>
                    <p style={{ display: 'flex', alignItems: 'center'}}> 
                        <InfoIcon sx={{marginRight: '0.5rem'}}/>
                        You need to include template to start
                    </p>
                    <ul>
                        <li>main.py(file)</li>
                        <li>Test files for user</li> 
                    </ul>
                </div>
            </div>

            <div style={{display:'flex', alignItems: 'center'}}>
                <ZipFileIcon style={{marginRight:10}}/>
                <InputWrapper  label="Upload test case .zip file: "><input onChange={onTestCaseFileChange} id="testFileInput" type="file" name="submissionFile" accept=".zip" style={{color: testFontColor, fontFamily:'Poppins', marginLeft:20}}/></InputWrapper>
            </div>
            <div style={{ padding: '10px', borderRadius: '8px', display: 'flex', alignItems: 'center', background: '#eff7ff'}}>
                <div style={{color: '#808080'}}>
                    <p style={{ display: 'flex', alignItems: 'center'}}> 
                        <InfoIcon sx={{marginRight: '0.5rem'}}/>
                        You need to include the testcases for final evaluation
                    </p>
                </div>
            </div>

            <Button sx={{margin: '1rem'}}variant="contained" onClick={()=>handleSubmit()}>Submit</Button>


        <div>
            <button onClick={getAllChallenges} style={{ marginRight: '30px' }} >All Challenges</button>
            <button onClick={getMyChallenges} style={{ marginRight: '30px' }}>My Challenges</button>
            <button onClick={createChallengePage} >Create Challenges</button>
        </div>
        </div>
    );
}

export default CreateChallengePage;
