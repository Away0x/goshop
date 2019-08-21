import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { GUIDE_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '导学本';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'guide',
    path: '/guide',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: GUIDE_AUTH,
    },
    children: [
      {
        default: true,
        path: '/list',
        name: 'guide-list',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'guide-list' */ '@/pages/Guide')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: GUIDE_AUTH,
        },
      },
    ],
  },
];

export default routes;
