import { createSlice, type PayloadAction } from "@reduxjs/toolkit"
import type { RootState } from "../../store"
import type { IHub } from "../../../types"

interface HubState {
  hubs: IHub[]
}

const initialState: HubState = {
  hubs: []
}

export const hubSlice = createSlice({
  name: 'hubs',
  initialState,
  reducers: {
    setHubs: (state, action: PayloadAction<IHub[]>) => {
      state.hubs = action.payload
    }
  },
})

export const { setHubs } = hubSlice.actions

export const selectHubs = (state: RootState) => state.hubs.hubs

export default hubSlice.reducer