// base layout 的一些状态
import { useSetState, UseSetState } from "react-hanger";

import { BaseValue, createStore } from 'aw-react-store';

export interface BreadcrumbsType {
  name: string
  routeName?: string
  handler?: (b: BreadcrumbsType) => void
}

interface State {
  readonly breadcrumbs: Readonly<BreadcrumbsType[]> // 面包屑
}

const initialState: State = {
  breadcrumbs: [],
};

export interface BaseLayoutValue extends BaseValue<State> {
  // actions
  setBreadcrumbsAction: (breadcrumbs: BreadcrumbsType[]) => void

  // getters
}

/** 设置 layout 的面包屑 */
function setBreadcrumbsAction(state: UseSetState<State>) {
  return (breadcrumbs: BreadcrumbsType[]) => {
    state.setState({
      breadcrumbs,
    });
  }
}

function useBaseLayout(): BaseLayoutValue {
  const state = useSetState<State>(initialState);

  return {
    state: state.state,

    // actions
    setBreadcrumbsAction: setBreadcrumbsAction(state),
  };
}

export default createStore(useBaseLayout);
