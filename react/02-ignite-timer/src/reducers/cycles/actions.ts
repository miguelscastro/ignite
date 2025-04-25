import { Cycle } from './reducer'

interface addNewCycleActionProps {
  type: ActionTypes.ADD_NEW_CYCLE
  payload: {
    newCycle: Cycle
  }
}
interface markCurrentCycleAsFinishedActionProps {
  type: ActionTypes.MARK_CURRENT_CYCLE_AS_FINISHED
}
interface interruptCurrentCycleActionProps {
  type: ActionTypes.INTERRUPT_CURRENT_CYCLE
}

export type ActionTypesProps =
  | addNewCycleActionProps
  | markCurrentCycleAsFinishedActionProps
  | interruptCurrentCycleActionProps

export enum ActionTypes {
  ADD_NEW_CYCLE = 'ADD_NEW_CYCLE',
  INTERRUPT_CURRENT_CYCLE = 'INTERRUPT_CURRENT_CYCLE',
  MARK_CURRENT_CYCLE_AS_FINISHED = 'MARK_CURRENT_CYCLE_AS_FINISHED',
}
export function addNewCycleAction(newCycle: Cycle): addNewCycleActionProps {
  return {
    type: ActionTypes.ADD_NEW_CYCLE,
    payload: {
      newCycle,
    },
  }
}
export function markCurrentCycleAsFinishedAction(): markCurrentCycleAsFinishedActionProps {
  return {
    type: ActionTypes.MARK_CURRENT_CYCLE_AS_FINISHED,
  }
}
export function interruptCurrentCycleAction(): interruptCurrentCycleActionProps {
  return {
    type: ActionTypes.INTERRUPT_CURRENT_CYCLE,
  }
}
