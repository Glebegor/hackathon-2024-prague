// filter.reducer.ts
import { createReducer, on } from '@ngrx/store';
import { setFilter } from './filter.actions';

export const initialState: string = '';

const _filterReducer = createReducer(
  initialState,
  on(setFilter, (state, { filter }) => filter)
);

export function filterReducer(state, action) {
  return _filterReducer(state, action);
}