import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { STATISTICS_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '统计分析';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'statistics',
    path: '/statistics',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: STATISTICS_AUTH,
    },
    children: [
      {
        default: true,
        name: 'resource-view',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'resource-view' */ '@/pages/Statistics')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: STATISTICS_AUTH,
        },
      },
    ],
  },
];

export default routes;
