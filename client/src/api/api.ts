import axios from "axios"
import type { IBooking, IEvent, IHub, IMentor, IRequest, IUser } from "../types"

const runtimeApiBaseUrl = `${window.location.protocol}//${window.location.hostname}:3000`
const resolvedApiBaseUrl = import.meta.env.VITE_API_BASE_URL ?? runtimeApiBaseUrl

const api = axios.create({
    baseURL: resolvedApiBaseUrl,
})

const authApiClient = axios.create({
        baseURL: resolvedApiBaseUrl,
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
        telegram?: string
        phone?: string
    }
}

type SessionUserResponse = {
    message: string
    user: {
        id: number
        fullname: string
        email: string
        role?: string
        telegram?: string
        phone?: string
    }
}

type LogoutResponse = {
    message: string
}

export interface IEvensApi {
    getEvents: () => Promise<IEvent[]>
    searchEvents: (query: string) => Promise<IEvent[]>
    saveEvent: (event: any) => Promise<string>
}

export interface IUsersApi {
    getUsers: () => Promise<IUser[]>
    getUserByEmail: (email: string) => Promise<IUser>
    saveUser: (user: any) => Promise<string>
}

type UserApiResponse = {
    user_id: number
    fullname: string
    user_role: string
    email: string
    telegram: string
    phone: string
}

const normalizeUser = (user: UserApiResponse): IUser => ({
    id: user.user_id,
    fullname: user.fullname,
    role: user.user_role,
    email: user.email,
    telegram: user.telegram,
    phone: user.phone,
})

export interface IHubsApi {
    getHubs: () => Promise<IHub[]>
    getHubById: (id: number) => Promise<IHub>
    saveHub: (hub: any) => Promise<string>
}

export interface IRequestsApi {
    getRequests: () => Promise<IRequest[]>
    getRequestsByUserId: (userId: number) => Promise<IRequest[]>
    getRequestsByMentorId: (mentorId: number) => Promise<IRequest[]>
    saveRequest: (request: any) => Promise<string>
}

export interface IBookingsApi {
    getBookingsByUserId: (userId: number) => Promise<IBooking[]>
    saveBooking: (booking: any) => Promise<any>
}

type HubApiResponse = {
    hub_id: number
    hub_name: string
    address: string
    status: string
    city: string
    description: string
    schedule: string
    occupancy: number
}

const normalizeHub = (hub: HubApiResponse): IHub => ({
    id: hub.hub_id,
    name: hub.hub_name,
    address: hub.address,
    status: hub.status,
    city: hub.city,
    desription: hub.description,
    schedule: hub.schedule,
    occupancy: hub.occupancy,
})

type EventApiResponse = {
    event_id: number
    name: string
    description: string
    start_time: string
    end_time: string
    hub_id: number
    mentor_id?: number
}

type RequestApiResponse = {
    request_id: number
    request_message: string
    user_id: number
    mentor_id?: number
}

type BookingApiResponse = {
    booking_id: number
    booking_date: string
    booking_zone: string
    booking_slots: number
    user_id: number
}

type MentorApiResponse = {
    mentor_id: number
    hub_id?: number
    user_id: number
    fullname: string
    user_role: string
    email: string
    telegram: string
    phone: string
}

const normalizeEvent = (event: EventApiResponse): IEvent => ({
    id: event.event_id,
    name: event.name,
    description: event.description,
    start: event.start_time,
    end: event.end_time,
    hubId: event.hub_id,
    mentorId: event.mentor_id,
})

const normalizeMentor = (mentor: MentorApiResponse): IMentor => ({
    id: mentor.user_id,
    mentorId: mentor.mentor_id,
    hubId: mentor.hub_id,
    fullname: mentor.fullname,
    role: mentor.user_role,
    email: mentor.email,
    telegram: mentor.telegram,
    phone: mentor.phone,
})

const normalizeRequest = (request: RequestApiResponse): IRequest => ({
    id: request.request_id,
    message: request.request_message,
    userId: request.user_id,
    mentorId: request.mentor_id,
})

const normalizeBooking = (booking: BookingApiResponse): IBooking => ({
    id: booking.booking_id,
    bookingDate: booking.booking_date,
    bookingZone: booking.booking_zone,
    bookingSlots: booking.booking_slots,
    userId: booking.user_id,
})

export const usersApi = {
    getUsers: async () => {
        const response = await api.get("/users/")
        return response.data
    },

    getUserByEmail: async (email: string) => {
        const response = await api.get<UserApiResponse>("/users/email", {
            params: { email },
        })

        return normalizeUser(response.data)
    },

    saveUser: async (user: any) => {
        const response = await api.post("/users/", user)
        return response.data
    }
}

export const eventsApi = {
    getEvents: async () => {
        const response = await api.get<EventApiResponse[]>("/events/")
        return response.data.map(normalizeEvent)
    },

    searchEvents: async (query: string) => {
        const response = await api.get<EventApiResponse[]>("/events/search", {
            params: { q: query },
        })
        return response.data.map(normalizeEvent)
    },

    saveEvent: async (event: any) => {
        const response = await authApiClient.post("/events/", event)
        return response.data
    }
}

export const hubsApi = {
    getHubs: async () => {
        const response = await api.get<HubApiResponse[]>("/hubs/")
        return response.data.map(normalizeHub)
    },

    getHubById: async (id: number) => {
        const response = await api.get<HubApiResponse>(`/hubs/${id}`)
        return normalizeHub(response.data)
    },

    searchHubs: async (query: string) => {
        const response = await api.get<HubApiResponse[]>("/hubs/search", {
            params: { q: query },
        })
        return response.data.map(normalizeHub)
    },

    saveHub: async (hub: any) => {
        const response = await api.post("/hubs/", hub)
        return response.data
    }
}

export const requestsApi = {
    getRequests: async () => {
        const response = await api.get<RequestApiResponse[]>("/requests")
        return response.data.map(normalizeRequest)
    },

    getRequestsByUserId: async (userId: number) => {
        const response = await api.get<RequestApiResponse[]>(`/requests/user/${userId}`)
        return response.data.map(normalizeRequest)
    },

    getRequestsByMentorId: async (mentorId: number) => {
        const response = await api.get<RequestApiResponse[]>(`/requests/mentor/${mentorId}`)
        return response.data.map(normalizeRequest)
    },

    saveRequest: async (request: any) => {
        const response = await api.post("/requests", request)
        return response.data
    }
}

export const bookingsApi = {
    getBookingsByUserId: async (userId: number) => {
        const response = await api.get<BookingApiResponse[]>(`/bookings/user/${userId}`)
        return response.data.map(normalizeBooking)
    },

    saveBooking: async (booking: any) => {
        const response = await api.post("/bookings/", booking)
        return response.data
    }
}

export const mentorsApi = {
    getMentors: async () => {
        const response = await api.get<MentorApiResponse[]>("/mentors/")
        return response.data.map(normalizeMentor)
    },

    getMentorByUserId: async (userId: number) => {
        const response = await api.get<MentorApiResponse>(`/mentors/user/${userId}`)
        return normalizeMentor(response.data)
    },

    saveMentor: async (mentor: any) => {
        const response = await api.post("/mentors/", mentor)
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

    getCurrentUser: async (): Promise<SessionUserResponse> => {
        const response = await authApiClient.get("/auth/me")
        return response.data
    },

    logout: async (): Promise<LogoutResponse> => {
        const response = await authApiClient.post("/auth/logout")
        return response.data
    },
}