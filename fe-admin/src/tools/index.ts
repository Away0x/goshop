import { isUndefined } from '@/tools/instance';

/**
 * safeGet 安全的获取对象值
 * obj = {a: {b: {c: 123}}}
 *
 * safeGet(obj, 'a.b.c') // 123
 * safeGet(obj, 'a.b.d', '我是默认值') // 我是默认值
 *
 * @param   {Object}  目标对象
 * @param   {String}  路径字符串 "a.b.c"
 * @param   {Any}     默认值
 * @return  {Any}     获取到的值
 */
export function safeGet<T>(targetObj: object, keyString: string, defaultValue: T) {
  const path = String(keyString).split('.')
  const result = path.reduce((preObj, curPath) => {
    return (preObj !== null && !isUndefined(preObj))
      ? (preObj as any)[curPath]
      : preObj
  }, targetObj)

  return isUndefined(result)
    ? (isUndefined(defaultValue) ? result : defaultValue)
    : result
}
