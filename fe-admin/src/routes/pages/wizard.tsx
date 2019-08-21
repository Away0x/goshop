import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import BaseLayout from '@/layouts/Base';
import { RouteMeta } from '@/routes/type';
import { WIZARD_AUTH } from '@/auth/page';

const TOP_NAV_NAME = '使用向导';

const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'wizard',
    path: '/wizard',
    component: BaseLayout,
    meta: {
      hanName: TOP_NAV_NAME,
      title: `教学平台-${TOP_NAV_NAME}`,
      topNavName: TOP_NAV_NAME,
      auth: WIZARD_AUTH,
    },
    children: [
      {
        default: true,
        name: 'wizard-wk',
        component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'wizard-wk' */ '@/pages/Wizard')} />,
        meta: {
          hanName: TOP_NAV_NAME,
          title: `教学平台-${TOP_NAV_NAME}`,
          topNavName: TOP_NAV_NAME,
          auth: WIZARD_AUTH,
        },
      },
    ],
  },
];

export default routes;
