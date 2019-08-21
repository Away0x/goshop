import React from 'react';

import AWSuspense from '@/routes/suspense';
import { AWRouteConfig } from 'aw-react-router';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { WORK_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '网络作业';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'work',
    path: '/work',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: WORK_AUTH,
    },
    children: [
      {
        default: true,
        path: '/list',
        name: 'work-list',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'work-list' */ '@/pages/Work/List')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: WORK_AUTH,
        },
      },
    ],
  },
];

export default routes;
