export interface IMentor extends IUser {
    id: number
    mentorId?: number
    hubId?: number
}

export interface INotification {

}

export interface IUser {
    id: number
    fullname: string
    email: string
    telegram?: string
    phone?: string
    role?: string
}

export interface IRequest {
    id: number
    message: string
    userId: number
    mentorId?: number
}

export interface IBooking {
    id: number
    bookingDate: string
    bookingZone: string
    bookingSlots: number
    userId: number
}

export interface IEvent {
    id: number
    name: string
    description: string
    start: string
    end: string
    hubId: number
    mentorId?: number
}

export interface IHub {
    id: number
    name: string
    address: string
    status: string
    city: string
    description: string
    schedule: string
    occupancy: number
}