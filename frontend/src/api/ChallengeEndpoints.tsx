import { AxiosInstance } from "axios";
import { BFF_URLS } from "../links/backend";

export const createChallenge = async (axios: AxiosInstance, challenge : any) => {
    const url = `${BFF_URLS.challengeService}/create`
    const method = "POST";
    const headers = {
        'Content-Type': 'application/json',
    };
    return axios.request({
        url,
        method,
        headers,
        data: challenge,
    });
}

export const getChallenges = async (axios: AxiosInstance) => {
    const url = `${BFF_URLS.challengeService}/challenges`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const getUserChallenges = async (axios: AxiosInstance, userId : string) => {
    const url = `${BFF_URLS.challengeService}/challenges/user/${userId}`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const getChallenge = async (axios: AxiosInstance, challengeId : string) => {
    const url = `${BFF_URLS.challengeService}/${challengeId}`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const getChallengesByDifficulty = async (axios: AxiosInstance, difficulty : string) => {
    const url = `${BFF_URLS.challengeService}/difficulty/${difficulty}`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const updateChallenge = async (axios: AxiosInstance, challengeId : string, challenge : any) => {
    const url = `${BFF_URLS.challengeService}/update/${challengeId}`
    const method = "PUT";
    const headers = {
        'Content-Type': 'application/json',
    };
    return axios.request({
        url,
        method,
        headers,
        data: challenge,
    });
}

export const deleteChallenge = async (axios: AxiosInstance, challengeId : string) => {
    const url = `${BFF_URLS.challengeService}/delete/${challengeId}`
    const method = "DELETE";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}
