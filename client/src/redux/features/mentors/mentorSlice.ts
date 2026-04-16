import { createSlice, type PayloadAction } from "@reduxjs/toolkit"
import type { RootState } from "../../store"
import type { IMentor } from "../../../types"


interface MentorState {
  mentors: IMentor[]
}

const initialState: MentorState = {
  mentors: []
}

export const mentorSlice = createSlice({
  name: 'mentors',
  initialState,
  reducers: {
    setMentors: (state, action: PayloadAction<IMentor[]>) => {
      state.mentors = action.payload
    }
  },
})

export const { setMentors } = mentorSlice.actions

export const selectMentors = (state: RootState) => state.mentors.mentors

export default mentorSlice.reducer