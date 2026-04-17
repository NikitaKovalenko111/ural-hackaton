import { all } from 'redux-saga/effects'
import eventSaga from './sagas/eventSaga'

function* mainSaga() {
  yield all([
    eventSaga()
  ])
}

export default mainSaga