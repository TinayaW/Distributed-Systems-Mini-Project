import { AxiosInstance } from "axios";
import { BFF_URLS } from "../links/backend";

export const createUser = async (axios: AxiosInstance, user : any) => {
    const url = `${BFF_URLS.userService}/create`
    const method = "POST";
    const headers = {
        'Content-Type': 'application/json',
    };
    return axios.request({
        url,
        method,
        headers,
        data: user,
    }); 
}

export const getUsers = async (axios: AxiosInstance) => {
    const url = `${BFF_URLS.userService}/users`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const getUser = async (axios: AxiosInstance, userId : string) => {
    const url = `${BFF_URLS.userService}/${userId}`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const getUserByUserName = async (axios: AxiosInstance, username : string) => {
    const url = `${BFF_URLS.userService}/username/${username}`
    const method = "GET";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}

export const updateUser = async (axios: AxiosInstance, userId : string, user : any) => {
    const url = `${BFF_URLS.userService}/update/${userId}`
    const method = "PUT";
    const headers = {
        'Content-Type': 'application/json',
    };
    return axios.request({
        url,
        method,
        headers,
        data: user,
    });
}

export const deleteUser = async (axios: AxiosInstance, userId : string) => {
    const url = `${BFF_URLS.userService}/delete/${userId}`
    const method = "DELETE";
    const headers = {};
    return axios.request({
        url,
        method,
        headers,
    });
}
