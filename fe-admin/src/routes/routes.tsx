import React from 'react';

import { AWRouteConfig } from 'aw-react-router';
import AWSuspense from '@/routes/suspense';
import { RouteMeta } from '@/routes/type';

import classroomRoutes from '@/routes/pages/classroom';
import coachRoutes from '@/routes/pages/coach';
import examRoutes from '@/routes/pages/exam';
import guideRoutes from '@/routes/pages/guide';
import questionsRoutes from '@/routes/pages/questions';
import resourceRoutes from '@/routes/pages/resource';
import statisticsRoutes from '@/routes/pages/statistics';
import textbookRoutes from '@/routes/pages/textbook';
import wizardRoutes from '@/routes/pages/wizard';
import workRoutes from '@/routes/pages/work';
import wrongRoutes from '@/routes/pages/wrong';
import settingsRoutes from '@/routes/pages/settings';

export const routes: AWRouteConfig<RouteMeta>[] = [
  {
    name: 'login',
    path: '/login',
    component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'login' */ '@/pages/Login')} />,
    meta: {
      title: '登录',
      hanName: '登录',
    },
  },
  {
    name: 'home',
    path: '/',
    component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'home' */ '@/pages/Home')} />,
    meta: {
      title: '首页',
      hanName: '首页',
    },
  },

  ...classroomRoutes,
  ...coachRoutes,
  ...examRoutes,
  ...guideRoutes,
  ...questionsRoutes,
  ...resourceRoutes,
  ...statisticsRoutes,
  ...textbookRoutes,
  ...wizardRoutes,
  ...workRoutes,
  ...wrongRoutes,
  ...settingsRoutes,

  {
    name: '404',
    path: '/404',
    component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'notfound' */ '@/pages/Errors/NotFound')} />,
    meta: {
      title: '404',
      hanName: '页面未找到',
    },
  },
  {
    name: 'unauthorized',
    path: '/unauthorized',
    component: (props) => <AWSuspense {...props} load={import(/* webpackChunkName: 'unauthorized' */ '@/pages/Errors/Unauthorized')} />,
    meta: {
      title: 'unauthorized',
      hanName: '无权限访问',
    },
  },
];
