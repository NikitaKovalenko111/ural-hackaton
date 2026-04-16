export interface IMentor extends IUser {
    id: number
}

export interface INotification {

}

export interface IUser {
    id: number
    fullname: string
    email: string
    telegram: string
    phone: string
}

export interface IRequest {
    id: number
    message: string
    userId: number
}

export interface IEvent {
    id: number
    name: string
    description: string
    start: Date
    end: Date
    hubId: number
}

export interface IHub {
    id: number
    name: string
    address: string
    status: string
    city: string
    desription: string
    schedule: string
    occupancy: number
}