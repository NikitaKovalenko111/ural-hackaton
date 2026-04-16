import { call, put, takeLatest } from "redux-saga/effects";
import { eventsApi } from "../../../api/api";
import type { IEvent } from "../../../types";
import { all } from "axios";

function* getEventsWorker() {
    try {
        const response: IEvent[] = yield call(eventsApi.getEvents)

        yield put({ type: 'events/getEventsSuccess', payload: response })
    } catch (error) {
        console.log(error);
    }
}

function* getEventsWatcher() {
  yield takeLatest('events/getEvents', getEventsWorker)
}

function* saveEventWorker(action: any) {
    try {
        const response: string = yield call(eventsApi.saveEvent, action.payload)
        yield put({ type: 'events/saveEventSuccess', payload: response })
    } catch (error) {
        console.log(error);
    }
}

function* saveEventWatcher() {
  yield takeLatest('events/saveEvent', saveEventWorker)
}

export default function* eventSaga() {
    yield all([
        getEventsWatcher(),
        saveEventWatcher()
    ])
}