import React, { memo, FC, useCallback } from 'react';
import {
  Breadcrumb,
} from 'antd';

import { awRouter } from '@/routes';
import { withRouter, RouteComponentProps } from 'aw-react-router';
import BaseLayoutStore, { BreadcrumbsType } from '@/store/base-layout';

const Breadcrumbs: FC<RouteComponentProps> = ({ history }) => {
  const { state } = BaseLayoutStore.useStore();

  const defaultHandler = useCallback((b: BreadcrumbsType) => {
    if (!b.handler && b.routeName) {
      const route = awRouter.find(b.routeName);
      history.push(route ? route.fullPath : '/404');
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

export default memo(withRouter(Breadcrumbs));
