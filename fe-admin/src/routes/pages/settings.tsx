import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';

const TOP_NAV_NAME = '个人设置';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'settings',
    path: '/settings',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: [],
    },
    children: [
      {
        default: true,
        name: 'settings-default',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'settings' */ '@/pages/Settings')} />,
        meta: {
          hanName: `${TOP_NAV_NAME}`,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: [],
        },
      },
    ],
  },
];

export default routes;
