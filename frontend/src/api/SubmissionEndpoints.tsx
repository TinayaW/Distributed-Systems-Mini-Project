import { AxiosInstance } from "axios";
import { BFF_URLS } from "../links/backend";

export const createSubmission = async (axios: AxiosInstance, submission : any) => {
    const url = `${BFF_URLS.submissionService}/upload`
    const method = "POST";
    const headers = {
        'Content-Type': 'application/json',
    };
    return axios.request({
        url,
        method,
        headers,
        data: submission,
    });
}

export const getSubmissionByUserId = async (axios: AxiosInstance, userId : string) => {
    const url = `${BFF_URLS.submissionService}/user/${userId}`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const getSubmissionByChallengeId = async (axios: AxiosInstance, challengeId : string) => {
    const url = `${BFF_URLS.submissionService}/challenge/${challengeId}`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const getSubmission = async (axios: AxiosInstance, submissionId : string) => {
    const url = `${BFF_URLS.submissionService}/${submissionId}`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}