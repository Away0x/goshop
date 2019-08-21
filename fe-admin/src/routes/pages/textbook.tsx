import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { TEXTBOOK_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '课本管理';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'textbook',
    path: '/textbook',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: TEXTBOOK_AUTH,
    },
    children: [
      {
        default: true,
        name: 'textbook-chapter',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'textbook-chapter' */ '@/pages/Textbook')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: TEXTBOOK_AUTH,
        },
      },
    ],
  },
];

export default routes;
