import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { EXAM_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '入学考试';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'exam',
    path: '/exam',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: EXAM_AUTH,
    },
    children: [
      {
        default: true,
        path: '/list',
        name: 'exam-list',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'exam-list' */ '@/pages/Exam')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: EXAM_AUTH,
        },
      },
    ],
  },
];

export default routes;
