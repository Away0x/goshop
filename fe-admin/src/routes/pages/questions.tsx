import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { QUESTIONS_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '题库';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'questions',
    path: '/questions',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: QUESTIONS_AUTH,
    },
    children: [
      {
        default: true,
        path: '/list',
        name: 'questions-list',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'questions-list' */ '@/pages/Questions')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: QUESTIONS_AUTH,
        },
      },
    ],
  },
];

export default routes;
