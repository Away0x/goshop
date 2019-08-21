import { useCallback, useMemo, useReducer } from 'react';
import PminDelay from 'p-min-delay';

import { CommonResponse } from '@/services/type';

interface PromiseErrorType {
  hasError: boolean
  errorMessage?: string
}

interface UsePromiseOptions<T> {
  delay?: number  // 当 promise 执行时间小于 delay 时，会等待 delay 时间过后再改变状态
  checkError?: (result: T) => PromiseErrorType // 用于检测 promise 是否有错误
}

interface UsePromiseReturn<T> {
  initial: boolean         // 初始状态，promise 还未执行
  run: () => Promise<void>          // 触发 promise
  pending: boolean         // promise 执行中
  hasError: boolean        // 发生错误了
  errorMessage: string     // 错误信息
  result: T                // promise 的返回值
}

enum ActionType {
  START = 'START',         // 开始执行 promise
  FINISH = 'FINISH',       // promise 执行结束
  HAS_ERROR = 'HAS_ERROR', // promise 出现错误
}

interface PromiseState<Result = any> {
  initial: boolean
  pending: boolean
  hasError: boolean
  errorMessage: string
  result: Result
}

interface ActionPayload<Result = any> {
  hasError: boolean,
  errorMessage: string
  result: Result
}

interface ActionData<Result = any> {
  type: ActionType
  payload?: ActionPayload<Result>
}

function promiseReducer(state: PromiseState, action: ActionData): PromiseState {
  switch (action.type) {
    case ActionType.START:
      return {
        ...state,
        initial: false,
        pending: true,
        errorMessage: '',
        hasError: false,
      };
    case ActionType.FINISH:
      return {
        ...state,
        initial: false,
        pending: false,
        ...action.payload,
      };
    case ActionType.HAS_ERROR:
      return {
        ...state,
        initial: false,
        pending: false,
        ...action.payload,
      };
    default:
      return state;
  }
}

const startActionCreator = (): ActionData => ({type: ActionType.START});
const finishActionCreator = (result: any): ActionData => {
  return { type: ActionType.FINISH, payload: {
    hasError: false,
    errorMessage: '',
    result,
  } };
}
const hasErrorActionCreator = (errorMessage = '', result: any = null,): ActionData => {
  return {
    type: ActionType.HAS_ERROR,
    payload: { hasError: true, errorMessage, result },
  };
}

export function usePromise<Result>(
  func: () => Promise<Result>,
  initialState: Result,
  {
    delay,
    checkError,
  }: UsePromiseOptions<Result> = {},
): UsePromiseReturn<Result> {
  const [state, dispatch] = useReducer(promiseReducer, {
    initial: true,
    pending: true,
    hasError: false,
    errorMessage: '',
    result: initialState,
  });

  const run = useCallback(
    async () => {
      dispatch(startActionCreator());

      try {
        let resp: Result;
        if (delay) {
          resp = await PminDelay(func(), delay);
        } else {
          resp = await func();
        }

        if (checkError) {
          const e = checkError(resp);
          if (e.hasError) {
            dispatch(hasErrorActionCreator(e.errorMessage, resp));
            return;
          }
        }

        dispatch(finishActionCreator(resp));
      } catch (err) {
        dispatch(hasErrorActionCreator(String(err)));
        console.warn('usePromise error: ', err);
      }
    },
    [delay, checkError],
  );

  return useMemo(
    () => ({
      ...state,
      run,
    }),
    [state, run],
  );
}

export function useRequest<Result> (
  func: () => Promise<CommonResponse<Result>>,
  initialState: Result,
  {
    delay,
    checkError,
  }: UsePromiseOptions<CommonResponse<Result>> = {},
) {
  const check = useMemo(() => {
    if (!checkError) {
      return (result: CommonResponse<Result>) => {
        if (!result.status) {
          return { hasError: true, errorMessage: result.message || '' };
        }

        return { hasError: false, errorMessage: '' };
      };
    }
    return checkError;
  }, []);

  const pr = usePromise(func,
    { status: true, message: '', data: initialState },
    { delay, checkError: check },
  );

  return {
    ...pr,
    result: pr.result.data,
  }
}