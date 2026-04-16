import axios from "axios"
import type { IEvent, IHub, IRequest, IUser } from "../types"

const api = axios.create({
  baseURL: "http://localhost:3000/api",
})

const authApiClient = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL ?? "http://localhost:3000",
    withCredentials: true,
})

type RequestMagicLinkResponse = {
    message: string
}

type VerifyMagicLinkResponse = {
    message: string
    user: {
        id: number
        fullname: string
        email: string
        role?: string
    }
}

export interface IEvensApi {
    getEvents: () => Promise<IEvent[]>
    saveEvent: (event: any) => Promise<string>
}

export interface IUsersApi {
    getUsers: () => Promise<IUser[]>
    saveUser: (user: any) => Promise<string>
}

export interface IHubsApi {
    getHubs: () => Promise<IHub[]>
    saveHub: (hub: any) => Promise<string>
}

export interface IRequestsApi {
    getRequests: () => Promise<IRequest[]>
    saveRequest: (request: any) => Promise<string>
}

export const usersApi = {
    getUsers: async () => {
        const response = await api.get("/users")
        return response.data
    },

    saveUser: async (user: any) => {
        const response = await api.post("/users", user)
        return response.data
    }
}

export const eventsApi = {
    getEvents: async () => {
        const response = await api.get("/events")
        return response.data
    },

    saveEvent: async (event: any) => {
        const response = await api.post("/events", event)
        return response.data
    }
}

export const hubsApi = {
    getHubs: async () => {
        const response = await api.get("/hubs")
        return response.data
    },

    saveHub: async (hub: any) => {
        const response = await api.post("/hubs", hub)
        return response.data
    }
}

export const requestsApi = {
    getRequests: async () => {
        const response = await api.get("/requests")
        return response.data
    },

    saveRequest: async (request: any) => {
        const response = await api.post("/requests", request)
        return response.data
    }
}

export const mentorsApi = {
    getMentors: async () => {
        const response = await api.get("/mentors")
        return response.data
    },

    saveMentor: async (mentor: any) => {
        const response = await api.post("/mentors", mentor)
        return response.data
    },

    deleteMentor: async (id: number) => {
        const response = await api.delete(`/mentors/${id}`)
        return response.data
    }
}

export const authApi = {
    requestMagicLink: async (email: string): Promise<RequestMagicLinkResponse> => {
        const response = await authApiClient.post("/auth/request", { email })
        return response.data
    },

    verifyMagicLink: async (token: string): Promise<VerifyMagicLinkResponse> => {
        const response = await authApiClient.get("/auth/verify", {
            params: { token },
        })
        return response.data
    },
}