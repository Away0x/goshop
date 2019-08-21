import React, { FC } from 'react';
import { createPortal } from 'react-dom';
import { Spin } from 'antd';

import './index.less';

import { APP_ROOT_ID } from '@/constants';

interface Props {
  show: boolean
  body?: boolean
  deplay?: number
}

const PageSpin: FC<Props> = ({ body, deplay, show, children }) => {
  if (body) {
    return createPortal(
      <div className="PageSpin_component">
        <Spin size="large" spinning={show} delay={deplay!}>
          {children}
        </Spin>
      </div>,
      document.getElementById(APP_ROOT_ID)!,
    );
  }

  return (
    <Spin size="large" spinning={show} delay={deplay!}>
      {children}
    </Spin>
  );
};

PageSpin.defaultProps = {
  body: false,
  deplay: 500,
}

export default PageSpin;
