import { createSlice, type PayloadAction } from "@reduxjs/toolkit"
import type { RootState } from "../../store"
import type { IEvent } from "../../../types"


interface EventState {
  events: IEvent[]
}

const initialState: EventState = {
  events: []
}

export const eventSlice = createSlice({
  name: 'events',
  initialState,
  reducers: {
    setEvents: (state, action: PayloadAction<IEvent[]>) => {
      state.events = action.payload
    }
  },
})
  
export const { setEvents } = eventSlice.actions

export const selectEvents = (state: RootState) => state.events

export default eventSlice.reducer