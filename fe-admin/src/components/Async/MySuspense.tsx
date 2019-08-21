import React, { Suspense, lazy, FC, ReactNode, ComponentType } from 'react';
import PminDelay from 'p-min-delay';

interface Prop {
  load: Promise<{default: ComponentType<any>}> // 异步加载的 promise component
  fallback?: NonNullable<ReactNode> | null     // 加载时显示的 loading
  delay?: number                               // 延迟 delay 秒显示 load 加载到的组件，避免加载太快的闪现
}

const MySuspense: FC<Prop> = ({ fallback, load, delay, ...rest }) => {
  let LazyCom: ComponentType;
  if (delay) {
    LazyCom = lazy(() => PminDelay(load, delay));
  } else {
    LazyCom = lazy(() => load);
  }
  return (
    <Suspense fallback={fallback!}>
      <LazyCom {...rest} />
    </Suspense>
  );
}

MySuspense.defaultProps = {
  fallback: <div className="common_page_placeholder">loading ... ...</div>,
  delay: 0,
};


export default MySuspense;
