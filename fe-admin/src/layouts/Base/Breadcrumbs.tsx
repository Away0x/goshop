import React, { memo, FC, useCallback } from 'react';
import {
  Breadcrumb,
} from 'antd';

import AWRouter from 'aw-react-router';
import BaseLayoutStore, { BreadcrumbsType } from '@/store/base-layout';

const Breadcrumbs: FC = () => {
  const { state } = BaseLayoutStore.useStore();

  const defaultHandler = useCallback((b: BreadcrumbsType) => {
    if (!b.handler && b.routeName) {
      AWRouter.instance().pushByName(b.routeName);
    }
  }, []);

  return (
    <Breadcrumb separator=">">
      {
        state.breadcrumbs.map((b) => {
          const handler = b.handler || defaultHandler;

          return (
            <Breadcrumb.Item key={b.name} onClick={() => handler(b)}>
              {b.name}
            </Breadcrumb.Item>
          );
        })
      }
    </Breadcrumb>
  );
}

export default memo(Breadcrumbs);
