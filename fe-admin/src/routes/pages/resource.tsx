import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { RESOURCE_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '备课中心';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'resource',
    path: '/resource',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: RESOURCE_AUTH,
    },
    children: [
      {
        default: true,
        path: '/list',
        name: 'resource-list',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'resource-list' */ '@/pages/Resource')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: RESOURCE_AUTH,
        },
      },
    ],
  },
];

export default routes;
