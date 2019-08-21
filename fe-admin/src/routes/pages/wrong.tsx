import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { WRONG_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '错题集';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'wrong',
    path: '/wrong',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: WRONG_AUTH,
    },
    children: [
      {
        default: true,
        path: '/list',
        name: 'wrong-list',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'wrong-list' */ '@/pages/Wrong')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: WRONG_AUTH,
        },
      },
    ],
  },
];

export default routes;
