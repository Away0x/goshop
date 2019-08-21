import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { CLASSROOM_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '智慧课堂';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'classroom',
    path: '/classroom',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: CLASSROOM_AUTH,
    },
    children: [
      {
        default: true,
        path: '/list',
        name: 'classroom-list',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'classroom-list' */ '@/pages/ClassRoom')} />,
        meta: {
          hanName: `${TOP_NAV_NAME}列表`,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: CLASSROOM_AUTH,
        },
      },
    ],
  },
];

export default routes;
