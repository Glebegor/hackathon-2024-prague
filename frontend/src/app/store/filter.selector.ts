// filter.selectors.ts
import { createSelector, createFeatureSelector } from '@ngrx/store';

const selectFilterState = createFeatureSelector<string>('filter');

export const selectFilterValue = createSelector(
  selectFilterState,
  state => state
);