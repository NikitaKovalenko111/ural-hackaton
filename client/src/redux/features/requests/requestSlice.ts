import { createSlice, type PayloadAction } from "@reduxjs/toolkit"
import type { RootState } from "../../store"
import type { IRequest } from "../../../types"


interface RequestState {
  requests: IRequest[]
}

const initialState: RequestState = {
  requests: []
}

export const requestSlice = createSlice({
  name: 'requests',
  initialState,
  reducers: {
    setRequests: (state, action: PayloadAction<IRequest[]>) => {
      state.requests = action.payload
    }
  },
})

export const { setRequests } = requestSlice.actions

export const selectRequests = (state: RootState) => state.requests.requests

export default requestSlice.reducer