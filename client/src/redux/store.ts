import { configureStore } from '@reduxjs/toolkit'
import createSagaMiddleware from 'redux-saga'
import mySaga from './saga/saga'
import eventReducer from './features/events/eventSlice'
import mentorReducer from './features/mentors/mentorSlice'
import requestReducer from './features/requests/requestSlice'
import userReducer from './features/users/userSlice'
import hubReducer from './features/hubs/hubSlice'

const sagaMiddleware = createSagaMiddleware()

export const store = configureStore({
  reducer: {
    events: eventReducer,
    mentors: mentorReducer,
    requests: requestReducer,
    users: userReducer,
    hubs: hubReducer
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(sagaMiddleware)
})

sagaMiddleware.run(mySaga)

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch