
/** 判断是否为 promise */
export function isPromise(val: any) {
  return val && typeof val.then === 'function';
}

/** 判断是否未 undefined */
export function isUndefined(val: any): boolean {
  return val === void 0;
}