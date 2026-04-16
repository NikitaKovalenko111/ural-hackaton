import { createSlice, type PayloadAction } from "@reduxjs/toolkit"
import type { RootState } from "../../store"
import type { IUser } from "../../../types"


interface UserState {
  user: null | IUser
}

const initialState: UserState = {
  user: null
}

export const userSlice = createSlice({
  name: 'users',
  initialState,
  reducers: {
    setUser: (state, action: PayloadAction<IUser | null>) => {
      state.user = action.payload
    }
  },
})

export const { setUser } = userSlice.actions

export const selectUser = (state: RootState) => state.users.user

export default userSlice.reducer