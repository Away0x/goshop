import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { COACH_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '作业辅导';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'coach',
    path: '/coach',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: COACH_AUTH,
    },
    children: [
      {
        default: true,
        path: '/list',
        name: 'coach-list',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'coach_list' */ '@/pages/Coach')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: COACH_AUTH,
        },
      },
    ],
  },
];

export default routes;
